package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
)

type Project struct {
	data.Project
	Hyperlinks []data.ProjectHyperlink
	Categories []data.ProjectCategory
}

func FindProject(w http.ResponseWriter, r *http.Request) {
	project_name := r.URL.Query().Get("name");
	//project_type := r.URL.Query().Get("type");

	project := data.GetProjectByName(project_name);

	if(project != nil) {
		json, _ := json.Marshal(*project);
		w.Write(json);
	}
}

func FindProjectByName(projectName string) *Project {
	return nil;
}

func FindProjectsForSearchStringInCategories(queryString string, categories []string) []Project {
	return []Project{};
}

func FindProjectsByCategories(categories []string) []Project {
	return []Project{};
}