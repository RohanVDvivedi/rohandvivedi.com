package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
	"strings"
)

// api handlers in this file
var FindProject = http.HandlerFunc(findProject)

type Project struct {
	data.Project
	Hyperlinks []data.ProjectHyperlink
	Categories []data.ProjectCategory
}

func findProject(w http.ResponseWriter, r *http.Request) {
	var projects []Project = []Project{};

	name, exists_name := r.URL.Query()["name"];

	get_all, exists_get_all := r.URL.Query()["get_all"];

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
		project_p := findProjectByName(name[0], withHyperlinks, withCategories)
		if(project_p != nil) {
			projects = append(projects, *project_p)
		}
	} else if (exists_get_all) {
		if(get_all[0] == "true"){
			projects = getAllProjects()
		}
	} else if (exists_query) {
		projects = findProjectsForSearchStringInCategories(query[0]);
	} else if (exists_categories) {
		projects = findProjectsByCategories(categories_list)
	}

	json, _ := json.Marshal(projects);
	w.Write(json);
}

func findProjectByName(name string, withHyperlinks bool, withCategories bool) *Project {
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

func getAllProjects() []Project {
	projects := []Project{}
	projects_db := data.GetAllProjects();
	for _, project_db := range projects_db {
		projects = append(projects, Project{project_db,nil,nil}) 
	}
	return projects;
}

func findProjectsForSearchStringInCategories(queryString string) []Project {
	projects := []Project{}
	projects_db := data.SearchProjectsByQueryString(queryString);
	for _, project_db := range projects_db {
		projects = append(projects, Project{project_db,nil,nil}) 
	}
	return projects;
}

func findProjectsByCategories(categories []string) []Project {
	projects := []Project{}
	projects_db := data.GetProjectsForCategoryNames(categories)
	for _, project_db := range projects_db {
		projects = append(projects, Project{project_db,nil,nil}) 
	}
	return projects;
}