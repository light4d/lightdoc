package route

import (
	"lightdoc/controller"
	"net/http"
)

func Init() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			controller.FileHandler(writer, request)
		case http.MethodPatch:
			controller.IndexHandler(writer, request)
		}
	})
	http.ListenAndServe(":8000", nil)
}
