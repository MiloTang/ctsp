package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	Conf               = make(map[string]string, 2000)
	ConfModTime int64  = 0
	ConfigPath  string = ""
	ConfigFile  string = ""
)

func InitConfig(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}
		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			continue
		}
		temp := strings.Split(s, "->")
		if len(temp) > 1 {
			Conf[strings.TrimSpace(temp[0])] = strings.TrimSpace(temp[1])
		}
	}
	ConfModTime = GetFileModTime(path)
}
func ClearConf(modTime int64) {
	Conf = make(map[string]string, 2000)
	ConfModTime = modTime
}
func GetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Println("stat fileinfo error")
		return time.Now().Unix()
	}
	return fi.ModTime().Unix()
}
func UpdateConfig(path string) {
	modTime := GetFileModTime(path)
	if modTime != ConfModTime {
		ClearConf(modTime)
		InitConfig(path)
		fmt.Println("重新加载被修改的文件")
	}
}
