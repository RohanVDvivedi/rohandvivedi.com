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

	var project_p *Project = nil;
	//var projects []Project;

	name, exists_name := r.URL.Query()["name"];
	if (exists_name) {
		requested_hyperlinks, exists_get_hyperlinks := r.URL.Query()["get_hyperlinks"];
		var withHyperlinks bool = exists_get_hyperlinks && (requested_hyperlinks[0] == "true")
		project_p = FindProjectByName(name[0], withHyperlinks)
	}

	if(project_p != nil) {
		json, _ := json.Marshal(*project_p);
		w.Write(json);
	} else {
		w.Write([]byte("{}"));
	}
}

func FindProjectByName(name string, withHyperlinks bool) *Project {
	proj_db := data.GetProjectByName(name)
	if(proj_db != nil) {
		p := &Project{};
		p.Project = *proj_db

		// if hyperlinks are requested, add hyperlinks to the project dto
		if withHyperlinks {
			p.Hyperlinks = proj_db.GetHyperlinks();
		}

		return p;
	}
	return nil
}

func FindProjectsForSearchStringInCategories(queryString string, categories []string) []Project {
	return []Project{};
}

func FindProjectsByCategories(categories []string) []Project {
	return []Project{};
}

//