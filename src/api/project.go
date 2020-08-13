package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
	"strings"
)

type Project struct {
	data.Project
	Hyperlinks []data.ProjectHyperlink
	Categories []data.ProjectCategory
}

func FindProject(w http.ResponseWriter, r *http.Request) {
	var projects []Project = []Project{};

	name, exists_name := r.URL.Query()["name"];

	categories, exists_categories := r.URL.Query()["categories"];
	var categories_list []string = nil;
	if(exists_categories) {
		categories_list = strings.Split(categories[0], ",")
	}

	query, exists_query := r.URL.Query()["query"];

	if (exists_name) {
		requested_hyperlinks, exists_get_hyperlinks := r.URL.Query()["get_hyperlinks"];
		var withHyperlinks bool = exists_get_hyperlinks && (requested_hyperlinks[0] == "true")
		requested_categories, exists_get_categories := r.URL.Query()["get_categories"];
		var withCategories bool = exists_get_categories && (requested_categories[0] == "true")
		project_p := FindProjectByName(name[0], withHyperlinks, withCategories)
		if(project_p != nil) {
			projects = append(projects, *project_p)
		}
	} else if (exists_query) {
		projects = FindProjectsForSearchStringInCategories(query[0], categories_list);
	} else if (exists_categories) {
		projects = FindProjectsByCategories(categories_list)
	}

	json, _ := json.Marshal(projects);
	w.Write(json);
}

func FindProjectByName(name string, withHyperlinks bool, withCategories bool) *Project {
	proj_db := data.GetProjectByName(name)
	if(proj_db != nil) {
		p := &Project{};
		p.Project = *proj_db

		// if hyperlinks are requested, add hyperlinks to the project dto
		if withHyperlinks {
			p.Hyperlinks = proj_db.GetProjectHyperlinks();
		}

		// if hyperlinks are requested, add hyperlinks to the project dto
		if withCategories {
			p.Categories = proj_db.GetProjectCategories();
		}

		return p;
	}
	return nil
}

func FindProjectsForSearchStringInCategories(queryString string, categories []string) []Project {
	// if categories is nil, search in all projects
	return []Project{};
}

func FindProjectsByCategories(categories []string) []Project {
	projects := []Project{};
	pcs := []data.ProjectCategory{};
	for _, category := range categories {
		pc := data.GetProjectCategoryByName(category);
		if(pc != nil) {
			pcs = append(pcs, *pc);
		}
	}
	for _, pc := range pcs {
		projects_db := pc.GetProjects();
		for _, project_db := range projects_db {
			projects = append(projects, Project{project_db,nil,nil})
		} 
	}
	return projects;
}

//