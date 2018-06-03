package common

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	Debug bool = true
)

func RootPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path = strings.Replace(path, "\\", "/", -1)
	return path
}
func Token() string {
	crutime := time.Now().Unix()
	t := md5.New()
	io.WriteString(t, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", t.Sum(nil))
	return token
}
func Tomd5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return fmt.Sprintf("%x", md.Sum(nil))
}
func Show(a ...interface{}) {
	if Debug {
		fmt.Println(a...)
	}

}
