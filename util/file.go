package util

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

// fileName:文件名字(带全路径)
// content: 写入的内容
func Tracefile(fileName string, content string) {
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	CheckErr(err)
	fd.WriteString("\n\n")
	fd.WriteString(content)
	defer fd.Close()
}

/**检查文件是否存在*/
func CheckFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**创建文件*/
func CreateFile(fpath string) {
	_, err := os.Create(fpath)
	CheckErr(err)
}

/**删除文件或文件夹*/
func RemoveFile(path string) {
	file, err := os.Stat(path)
	CheckErr(err)
	if file.IsDir() {
		os.RemoveAll(path)
	} else {
		os.Remove(path)
	}
}

/**创建目录*/
func Dir(path string) {
	//创建多级目录和设置权限
	os.MkdirAll(path, 0777)
}

/**复制文件*/
func Capy(oPath, nPath string) {
	//打开原始文件
	originalFile, err := os.Open(oPath)
	CheckErr(err)
	defer originalFile.Close()
	//创建新的文件
	newFile, err := os.Create(nPath)
	CheckErr(err)
	defer newFile.Close()
	//从源中复制字节到目标文件
	_, err = io.Copy(newFile, originalFile)
	CheckErr(err)
	//将文件内容flush到硬盘中
	err = newFile.Sync()
	CheckErr(err)
}

/**解压*/
func UnTarGz(srcFilePath string, destDirPath string) string {
	fmt.Println("UnTarGzing " + srcFilePath + "...")
	pathName := ""
	// Create destination directory
	os.Mkdir(destDirPath, os.ModePerm)

	fr, err := os.Open(srcFilePath)
	CheckErr(err)
	defer fr.Close()

	// Gzip reader
	gr, err := gzip.NewReader(fr)

	// Tar reader
	tr := tar.NewReader(gr)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		//handleError(err)
		if pathName == "" {
			pathName = getPathName(hdr.Name)
		}
		fmt.Println("UnTarGzing file..." + hdr.Name)
		// Check if it is diretory or file
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create diretory before create file
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			// Write data to file
			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
			CheckErr(err)
			_, err = io.Copy(fw, tr)
			CheckErr(err)
		}
	}
	fmt.Println("Well done!")
	return pathName
}

/**获取文件夹名称*/
func getPathName(path string) string {
	strs := strings.Split(path, "/")
	return strs[0]
}

/**加压*/
func TarGz(srcDirPath string, destFilePath string) {
	fw, err := os.Create(destFilePath)
	CheckErr(err)
	defer fw.Close()

	// Gzip writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// Tar writer
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Check if it's a file or a directory
	f, err := os.Open(srcDirPath)
	CheckErr(err)
	fi, err := f.Stat()
	CheckErr(err)
	if fi.IsDir() {
		// handle source directory
		fmt.Println("Cerating tar.gz from directory...")
		tarGzDir(srcDirPath, path.Base(srcDirPath), tw)
	} else {
		// handle file directly
		fmt.Println("Cerating tar.gz from " + fi.Name() + "...")
		tarGzFile(srcDirPath, fi.Name(), tw, fi)
	}
	fmt.Println("Well done!")
}

// Deal with directories
// if find files, handle them with tarGzFile
// Every recurrence append the base path to the recPath
// recPath is the path inside of tar.gz
func tarGzDir(srcDirPath string, recPath string, tw *tar.Writer) {
	// Open source diretory
	dir, err := os.Open(srcDirPath)
	CheckErr(err)
	defer dir.Close()

	// Get file info slice
	fis, err := dir.Readdir(0)
	CheckErr(err)
	for _, fi := range fis {
		// Append path
		curPath := srcDirPath + "/" + fi.Name()
		// Check it is directory or file
		if fi.IsDir() {
			// Directory
			// (Directory won't add unitl all subfiles are added)
			fmt.Printf("Adding path...%s\\n", curPath)
			tarGzDir(curPath, recPath+"/"+fi.Name(), tw)
		} else {
			// File
			fmt.Printf("Adding file...%s\\n", curPath)
		}

		tarGzFile(curPath, recPath+"/"+fi.Name(), tw, fi)
	}
}

// Deal with files
func tarGzFile(srcFile string, recPath string, tw *tar.Writer, fi os.FileInfo) {
	if fi.IsDir() {
		// Create tar header
		hdr := new(tar.Header)
		// if last character of header name is '/' it also can be directory
		// but if you don't set Typeflag, error will occur when you untargz
		hdr.Name = recPath + "/"
		hdr.Typeflag = tar.TypeDir
		hdr.Size = 0
		//hdr.Mode = 0755 | c_ISDIR
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()

		// Write hander
		err := tw.WriteHeader(hdr)
		CheckErr(err)
	} else {
		// File reader
		fr, err := os.Open(srcFile)
		CheckErr(err)
		defer fr.Close()

		// Create tar header
		hdr := new(tar.Header)
		hdr.Name = recPath
		hdr.Size = fi.Size()
		hdr.Mode = int64(fi.Mode())
		hdr.ModTime = fi.ModTime()

		// Write hander
		err = tw.WriteHeader(hdr)
		CheckErr(err)

		// Write file data
		_, err = io.Copy(tw, fr)
		CheckErr(err)
	}
}
