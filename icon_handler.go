package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func IconHandler(w http.ResponseWriter, r *http.Request) {
	path := path.Join("files", "icon.png")
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Can't read %s", path))
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(contents)
}
