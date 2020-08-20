package api

import (
	"net/http"
	//"rohandvivedi.com/src/data"
	"rohandvivedi.com/src/searchindex"
)

// api handlers in this file
var BuildSearchIndex = http.HandlerFunc(buildSearchIndex)

func buildSearchIndex(w http.ResponseWriter, r *http.Request) {
	err := searchindex.GetAndStoreGithubFile("RohanVDvivedi", "rohandvivedi.com", "README.md")
	w.Write([]byte("Cool\n"))
	if(err != nil) {
		w.Write([]byte(err.Error()))
	}
}