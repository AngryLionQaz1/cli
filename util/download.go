package util

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var queue, redo, finish chan int
var cor, size, length, timeout = 10, 2 * 1024 * 1024, 0, 0
var hash, dst = "md5", ""
var verify, cache bool

/**多线程下载文件*/
func Download(url, path string) {

	if dst == "" {
		_, dst = filepath.Split(url)
		dst = path + "/" + dst
		fmt.Println(dst)
	}

	startTime := time.Now()

	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(request)
	response.Body.Close()
	num := response.Header.Get("Content-Length")
	length, _ = strconv.Atoi(num)
	log.Println("Conetnt-Length", length)
	ranges := response.Header.Get("Accept-Ranges")
	log.Println("Ranges:", ranges)

	if size <= 0 {
		size = int(math.Ceil(float64(length) / float64(cor)))
	}
	fragment := int(math.Ceil(float64(length) / float64(size)))
	queue = make(chan int, cor)
	redo = make(chan int, int(math.Floor(float64(cor)/2)))
	go func() {
		for i := 0; i < fragment; i++ {
			queue <- i
		}
		for {
			j := <-redo
			queue <- j
		}
	}()
	finish = make(chan int, cor)
	for j := 0; j < cor; j++ {
		go do(request, fragment, j)
	}
	for k := 0; k < fragment; k++ {
		_ = <-finish
		//log.Printf("[%s][%d]Finished\n", "-", i)
	}
	log.Println("Start to combine files...")

	file, err := os.Create(dst)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	var offset int64 = 0
	for x := 0; x < fragment; x++ {
		filename := fmt.Sprintf("%s_%d", dst, x)
		buf, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Println(err)
			continue
		}
		file.WriteAt(buf, offset)
		offset += int64(len(buf))
		os.Remove(filename)
	}
	log.Println("Written to ", dst)
	//hash
	if hash == "sha1" {
		h := sha1.New()
		io.Copy(h, file)
		r := h.Sum(nil)
		log.Printf("sha1 of file: %x\n", r)
	} else if hash == "md5" {
		h := md5.New()
		io.Copy(h, file)
		r := h.Sum(nil)
		log.Printf("sha1 of file: %x\n", r)
	}

	finishTime := time.Now()
	duration := finishTime.Sub(startTime).Seconds()
	log.Printf("Time:%f Speed:%f Kb/s\n", duration, float64(length)/duration/1024)

}
func do(request *http.Request, fragment, no int) {
	var req http.Request
	err := deepCopy(&req, request)
	if err != nil {
		log.Println("ERROR|prepare request:", err)
		log.Panic(err)
		return
	}
	for {
		cStartTime := time.Now()

		i := <-queue
		//log.Printf("[%d][%d]Start download\n",no, i)
		start := i * size
		var end int
		if i < fragment-1 {
			end = start + size - 1
		} else {
			end = length - 1
		}

		filename := fmt.Sprintf("%s_%d", dst, i)
		if cache {
			filesize := int64(end - start + 1)
			file, err := os.Stat(filename)
			if err == nil && file.Size() == filesize {
				log.Printf("[%d][%d]Hint cached %s, size:%d\n", no, i, filename, filesize)
				finish <- i
				continue
			}
		}

		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
		log.Printf("[%d][%d]Start download:%d-%d\n", no, i, start, end)
		cli := http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		}
		resp, err := cli.Do(&req)
		if err != nil {
			log.Printf("[%d][%d]ERROR|do request:%s\n", no, i, err.Error())
			redo <- i
			continue
		}

		//log.Printf("[%d]Content-Length:%s\n", i, resp.Header.Get("Content-Length"))
		log.Printf("[%d][%d]Content-Range:%s\n", no, i, resp.Header.Get("Content-Range"))

		file, err := os.Create(filename)
		if err != nil {
			log.Printf("[%d][%d]ERROR|create file %s:%s\n", no, i, filename, err.Error())
			file.Close()
			resp.Body.Close()
			redo <- i
			continue
		}
		log.Printf("[%d][%d]Writing to file %s\n", no, i, filename)
		n, err := io.Copy(file, resp.Body)
		if err != nil {
			log.Printf("[%d][%d]ERROR|write to file %s:%s\n", no, i, filename, err.Error())
			file.Close()
			resp.Body.Close()
			redo <- i
			continue
		}
		cEndTime := time.Now()
		duration := cEndTime.Sub(cStartTime).Seconds()
		log.Printf("[%d][%d]Download successfully:%f Kb/s\n", no, i, float64(n)/duration/1024)

		file.Close()
		resp.Body.Close()

		finish <- i
	}
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
