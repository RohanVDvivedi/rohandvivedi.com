package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
)

func FindProject(w http.ResponseWriter, r *http.Request) {
	project_name := r.URL.Query().Get("name");
	//project_type := r.URL.Query().Get("type");

	project := data.GetProjectByName(project_name);

	if(project != nil) {
		json, _ := json.Marshal(*project);
		w.Write(json);
	}
}