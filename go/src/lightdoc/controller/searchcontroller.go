package controller

import (
	"encoding/json"
	"lightdoc/model"
	"net/http"
	"strconv"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("SEARCH")
	S := model.SerchKeyWords(key)
	res, _ := json.Marshal(S)
	w.Header().Set("Content-Length", strconv.Itoa(len(res)))
	w.Write(res)
}
