package main

import (
	"net/http"
	"os/exec"
)

func JsHandler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("./node_modules/browserify/bin/cmd.js", "-t", "hbsfy", "./js/app.js").Output()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
