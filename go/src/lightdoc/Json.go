package main

//导包
import (
	"net/http"  //http协议
	"encoding/json"  //json解析
	"fmt"  //读写
)

func main()  {
	//返回hello world
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"hello world")
	})
	//返回json
	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		word,_:=json.Marshal(map[string]string{"word":"hello world"})
		fmt.Fprint(writer,string(word))
	})
	http.ListenAndServe(":8888",nil)
}