package web

import (
	"log"
	"net/http"
	"strconv"
)

func getIntValue(r *http.Request, key string) int {
	v := r.Form.Get(key)
	if v == "" {
		v = r.PostForm.Get(key)
	}

	res, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalln(err)
		return -1
	}
	return res
}
