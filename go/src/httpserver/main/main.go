package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)
func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	bytes,_ := ReadAll("/home/zh/testforgo/a/1/1.1/test1/无标题文档")
	w.Write(bytes)
	var s []string
	s, _ = GetAllFile("/home/zh/testforgo", s)
	b, _ := json.Marshal(s)
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)

}

/**
	获取文件名称
 */
func GetAllFile(pathname string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + "/" + fi.Name()
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + "/" + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

/**
	读取文件内容
 */
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}