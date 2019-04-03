package controller

import (
	"encoding/json"
	"lightdoc/config"
	"lightdoc/model"
	"net/http"
	"strconv"
)

func PictureHandler(w http.ResponseWriter, r *http.Request) {
	pic := model.InitPic(config.Root)
	ps, err := json.Marshal(pic)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Length", strconv.Itoa(len(ps)))
	w.Write(ps)

}
