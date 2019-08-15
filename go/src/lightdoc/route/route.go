package route

import (
	"lightdoc/config"
	"lightdoc/controller"
	"net/http"
	"os"
)

func Init() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问所有域
		writer.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,PATCH,FILE,LOGO") //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")               //header的类型
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		switch request.Method {

		case http.MethodGet:
			f, err := os.Stat(config.Root + "/" + request.URL.Path)
			if err == nil && !f.IsDir() {
				http.FileServer(http.Dir(config.Root)).ServeHTTP(writer, request)
				break
			}

			F := http.FileServer(http.Dir(config.Dist))
			F.ServeHTTP(writer, request)
		case "PATCH":		//解决ss不支持自定义method的问题
			controller.TreeHandler(writer, request)
		case "LOGO":
			controller.LogoHandler(writer, request)
		case "SEARCH":
			controller.SearchHandler(writer, request)
		}
	})

	http.ListenAndServe(":8050", nil)
	return nil
}
