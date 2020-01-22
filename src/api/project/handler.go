package project

import (
	"net/http"
	"encoding/json"
)



func Handler(w http.ResponseWriter, r *http.Request) {
	projectNameToQuery := r.URL.Query().Get("projectName");

	if(projectNameToQuery == "") {
		w.Write([]byte("{}"));
		return
	}

	project := GetProjectByName(projectNameToQuery);
	if(project == nil) {
		w.Write([]byte("{}"));
		return
	}

	json, err := json.Marshal(*project);
	if(err == nil) {
		w.Write(json);
	} else {
		w.Write([]byte("{}"));
	}
}