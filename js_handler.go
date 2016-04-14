package main

import (
	"net/http"
	"os"
	"os/exec"
)

var (
	js []byte
)

func getJS() []byte {
	out, err := exec.Command("./node_modules/browserify/bin/cmd.js", "-t", "hbsfy", "./js/app.js").Output()
	if err != nil {
		panic(err)
	}
	return out
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("CACHE_JS") != "1" || len(js) == 0 {
		js = getJS()
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
