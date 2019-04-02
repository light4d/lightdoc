package controller

import (
	"io/ioutil"
	"lightdoc/config"
	"net/http"
	"strconv"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	path := config.Root + "/" + r.URL.Path
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(bs)))
	w.Write(bs)
}
