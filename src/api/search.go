package api

import (
	"net/http"
	"rohandvivedi.com/src/searchindex"
	"encoding/json"
)

// api handlers in this file
var BuildSearchIndex = http.HandlerFunc(buildSearchIndex)

func buildSearchIndex(w http.ResponseWriter, r *http.Request) {

	searchindex.InitProjectSearchIndex()
	
	searchindex.InsertProjectInSearchIndex("DummyProject")

	searchResult := searchindex.GetProjectSearchQueryResults("source")

	json, _ := json.Marshal(searchResult)
	w.Write(json)
}