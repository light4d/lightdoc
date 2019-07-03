package controller

import (
	"encoding/json"
	"lightdoc/config"
	"lightdoc/model"
	"net/http"
	"strconv"
)

func TreeHandler(w http.ResponseWriter, r *http.Request) {

	tree := model.InitTree(config.Root)
	bs, err := json.Marshal(tree)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(bs)))
	w.Write(bs)

}
