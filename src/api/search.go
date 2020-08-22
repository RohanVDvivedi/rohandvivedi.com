package api

import (
	"net/http"
	//"rohandvivedi.com/src/data"
	"rohandvivedi.com/src/searchindex"
	"fmt"
)

// api handlers in this file
var BuildSearchIndex = http.HandlerFunc(buildSearchIndex)

func buildSearchIndex(w http.ResponseWriter, r *http.Request) {
	responseFilename, err1 := searchindex.GetAndStoreGithubFile("RohanVDvivedi", "rohandvivedi.com", "README.md")
	if(err1 != nil) {
		fmt.Println(err1)
	}
	fmt.Println(responseFilename)

	err2 := searchindex.CreateProjectSearchIndex("./db/projectSearchIndex.bleve")
	if(err2 != nil) {
		fmt.Println(err2)
	}

	err3 := searchindex.AddSearchIndexProjectDocument("rohandvivedi.com", responseFilename)
	if(err3 != nil) {
		fmt.Println(err3)
	}

	searchindex.GetSearchResults("backend")

	w.Write([]byte("Cool\n"))
}