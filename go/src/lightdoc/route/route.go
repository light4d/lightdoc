package route

import (
	"lightdoc/config"
	"lightdoc/controller"
	"net/http"
	"os"
)

func Init() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")                  //允许访问所有域
		writer.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,PATCH") //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")      //header的类型
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		switch request.Method {
		case http.MethodGet:
			f, err := os.Stat(config.Root + "/" + request.URL.Path)
			if err == nil && !f.IsDir() {
				http.FileServer(http.Dir(config.Root)).ServeHTTP(writer, request)
			}

		case http.MethodConnect:
			writer.Write([]byte(config.Pubip))
		case http.MethodPatch:
			controller.IndexHandler(writer, request)
		case http.MethodOptions:
			controller.PictureHandler(writer, request)
		}
	})

	go http.ListenAndServe(":8050", nil)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(config.Dist)))
	return http.ListenAndServe(":8051", mux)
}
