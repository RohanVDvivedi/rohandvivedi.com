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
	GithubRepositoryLink *data.ProjectHyperlink
	Hyperlinks []data.ProjectHyperlink
	Categories []data.ProjectCategory
}

func findProject(w http.ResponseWriter, r *http.Request) {
	get_all, exists_get_all := r.URL.Query()["get_all"];

	names, exists_names := r.URL.Query()["names"];
	var names_list []string = nil;
	if(exists_names) {
		names_list = strings.Split(names[0], ",")
	}

	categories, exists_categories := r.URL.Query()["categories"];
	var categories_list []string = nil;
	if(exists_categories) {
		categories_list = strings.Split(categories[0], ",")
	}

	projects_db := []data.Project{};

	if(exists_names) {
		projects_db = data.GetProjectsByNames(names_list)
	} else if(exists_get_all) {
		if(get_all[0] == "true"){
			projects_db = data.GetAllProjects()
		}
	} else if(exists_categories) {
		projects_db = data.GetProjectsForCategoryNames(categories_list)
	}

	requested_github_link, exists_get_github_repo_link := r.URL.Query()["get_github_repo_link"];
	var withGithubRepoLink bool = exists_get_github_repo_link && (requested_github_link[0] == "true")

	requested_hyperlinks, exists_get_hyperlinks := r.URL.Query()["get_hyperlinks"];
	var withHyperlinks bool = exists_get_hyperlinks && (requested_hyperlinks[0] == "true")
	
	requested_categories, exists_get_categories := r.URL.Query()["get_categories"];
	var withCategories bool = exists_get_categories && (requested_categories[0] == "true")

	if(!withHyperlinks && !withCategories && !withGithubRepoLink) {
		json, _ := json.Marshal(projects_db);
		w.Write(json);
		return
	}

	var projects []Project = []Project{};
	for _, proj_db := range projects_db {
		p := Project{};
		p.Project = proj_db
		if withGithubRepoLink {
			p.GithubRepositoryLink = proj_db.GetProjectGithubRepositoryLink();
		}
		if withHyperlinks {
			p.Hyperlinks = proj_db.GetProjectHyperlinks();
		}
		if withCategories {
			p.Categories = proj_db.GetProjectCategories();
		}
		projects = append(projects, p)
	}

	json, _ := json.Marshal(projects);
	w.Write(json);
}