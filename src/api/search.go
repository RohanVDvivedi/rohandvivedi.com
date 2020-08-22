package api

import (
	"net/http"
	"rohandvivedi.com/src/searchindex"
	"encoding/json"
)

// api handlers in this file
var ProjectsSearch = http.HandlerFunc(projectsSearch)

func projectsSearch(w http.ResponseWriter, r *http.Request) {

	query, exists_query := r.URL.Query()["query"];
	searchResult := []string{}

	if(exists_query) {
		searchResult = searchindex.GetProjectSearchQueryResults(query[0])
	}

	json, _ := json.Marshal(searchResult)
	w.Write(json)
}