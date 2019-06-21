package route

import (
	"lightdoc/config"
	"lightdoc/controller"
	"net/http"
	"os"
)

func Init() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")                      //允许访问所有域
		writer.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,TREE,FILE,LOGO") //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")          //header的类型
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		switch request.Method {
		case "FILE":
			f, err := os.Stat(config.Root + "/" + request.URL.Path)
			if err == nil && !f.IsDir() {
				http.FileServer(http.Dir(config.Root)).ServeHTTP(writer, request)
			}
		case http.MethodGet:
			F := http.FileServer(http.Dir(config.Dist))
			F.ServeHTTP(writer, request)
		case "TREE":
			controller.TreeHandler(writer, request)
		case "LOGO":
			controller.LogoHandler(writer, request)
		}
	})

	http.ListenAndServe(":8050", nil)
	return nil
}
