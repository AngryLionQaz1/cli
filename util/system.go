package util

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

/**执行基本Linux命令*/
func ExecSys(cd string, ps ...string) {
	cmd := exec.Command(cd, ps...)
	buf, err := cmd.CombinedOutput()
	CheckErr(err)
	//fmt.Printf("%s\n",buf)
	PrintlnRight(string(buf))
}

/**执行需要输入 y 命令的 Linux 命令*/
func SysExec2(cd string, s ...string) {
	cmd := exec.Command(cd, s...)
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in
	in.WriteString("y")
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		PrintlnErr(err)
		return
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		//fmt.Println("Error:The command is err,", err)
		PrintlnErr(err)
		return
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {

			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				//fmt.Printf("Error :%s\n", err)
				PrintlnErr(err)
			}
			return
		}
		//fmt.Printf("%s\n", string(output))
		PrintlnRight(string(output))
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		//fmt.Println("wait:", err.Error())
		PrintlnErr(err)
		return
	}
}

/**执行Linux 命令*/
func SysExec(cd string, s ...string) {
	cmd := exec.Command(cd, s...)
	//in:=bytes.NewBuffer(nil)
	//cmd.Stdin=in
	//in.WriteString("y")
	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		PrintlnErr(err)
		return
	}
	//执行命令
	if err := cmd.Start(); err != nil {
		//fmt.Println("Error:The command is err,", err)
		PrintlnErr(err)
		return
	}
	//使用带缓冲的读取器
	outputBuf := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		output, _, err := outputBuf.ReadLine()
		if err != nil {

			// 判断是否到文件的结尾了否则出错
			if err.Error() != "EOF" {
				//fmt.Printf("Error :%s\n", err)
				PrintlnErr(err)
			}
			return
		}
		//fmt.Printf("%s\n", string(output))
		PrintlnRight(string(output))
	}
	//wait 方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		//fmt.Println("wait:", err.Error())
		PrintlnErr(err)
		return
	}
}

/**获取Linux 系统*/
func Sys() string {
	cd := `os=""
if [[ -f /etc/redhat-release ]]; then
    os="centos"
elif cat /etc/issue | grep -Eqi "debian"; then
    os="debian"
elif cat /etc/issue | grep -Eqi "ubuntu"; then
    os="ubuntu"
elif cat /etc/issue | grep -Eqi "deepin"; then
    os="deepin"
elif cat /etc/issue | grep -Eqi "centos|red hat|redhat"; then
    os="centos"
elif cat /proc/version | grep -Eqi "debian"; then
    os="debian"
elif cat /proc/version | grep -Eqi "ubuntu"; then
    os="ubuntu"
elif cat /proc/version | grep -Eqi "centos|red hat|redhat"; then
    os="centos"    
fi
echo $os`
	cmd := exec.Command("/bin/bash", "-c", cd)
	buf, err := cmd.CombinedOutput()
	CheckErr(err)
	return strings.Replace(string(buf), "\n", "", -1)
}
