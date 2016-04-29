package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

func JSONHandler(w http.ResponseWriter, r *http.Request) {
	filename := "all.json"
	if strings.HasSuffix(r.URL.Path, "packages/") {
		if r.FormValue("q") != "" {
			filename = "search.json"
		} else if r.FormValue("installed_only") == "true" {
			filename = "installed.json"
		}
	} else {
		filename = "individual.json"
	}

	path := path.Join("json", filename)
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Can't read %s", filename))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(contents)
}
