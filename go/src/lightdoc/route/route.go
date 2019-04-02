package route

import (
	"fmt"
	"io/ioutil"
	"lightdoc/config"
	"lightdoc/controller"
	"net/http"
	"strconv"
)

func Init() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("content-type", "application/json")             //返回数据格式是json
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		switch request.Method {
		case http.MethodGet:
			controller.FileHandler(writer, request)
		case http.MethodPatch:
			controller.IndexHandler(writer, request)
		}
	})
	go http.ListenAndServe(":8000", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := config.Dist + "/" + request.URL.Path
		bs, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}

		writer.Header().Set("Content-Length", strconv.Itoa(len(bs)))
		writer.Write(bs)
	})
	return http.ListenAndServe(":8080", mux)
}
