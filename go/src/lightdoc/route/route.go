package route

import (
	"lightdoc/config"
	"lightdoc/controller"
	"net/http"
)

func Init() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")                  //允许访问所有域
		writer.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,PATCH") //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")      //header的类型
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		switch request.Method {
		case http.MethodGet:
			http.FileServer(http.Dir(config.Root)).ServeHTTP(writer, request)
		case http.MethodPatch:
			controller.IndexHandler(writer, request)
		}
	})
	go http.ListenAndServe(":8000", nil)
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(config.Dist)))
	return http.ListenAndServe(":8080", mux)
}
