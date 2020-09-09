package api

import (
	"net/http"
	"encoding/json"
	"rohandvivedi.com/src/data"
	"strings"
	"rohandvivedi.com/src/searchindex"
)

// api handlers in this file
var ProjectsSearch = http.HandlerFunc(projectsSearch)

func projectsSearch(w http.ResponseWriter, r *http.Request) {
	query, exists_query := r.URL.Query()["query"];
	use_search_index := exists_query && (len(query[0]) > 0);

	categories, exists_categories := r.URL.Query()["categories"];
	var categories_list []string = nil;
	if(exists_categories) {
		categories_list = strings.Split(categories[0], ",")
	}

	projects_db := []data.Project{};

	if(use_search_index && exists_categories) {
		projectNamesResult := searchindex.GetProjectSearchQueryResults(query[0])
		projects_db = data.GetProjectsByNames(projectNamesResult)
	} else if(use_search_index && !exists_categories) {
		projectNamesResult := searchindex.GetProjectSearchQueryResults(query[0])
		projects_db = data.GetProjectsByNames(projectNamesResult)
	} else if (!use_search_index && exists_categories) {
		projects_db = data.GetProjectsForCategoryNames(categories_list)
	}

	requested_hyperlinks, exists_get_hyperlinks := r.URL.Query()["get_hyperlinks"];
	var withHyperlinks bool = exists_get_hyperlinks && (requested_hyperlinks[0] == "true")
	
	requested_categories, exists_get_categories := r.URL.Query()["get_categories"];
	var withCategories bool = exists_get_categories && (requested_categories[0] == "true")

	if(!withHyperlinks && !withCategories) {
		json, _ := json.Marshal(projects_db);
		w.Write(json);
		return
	}

	var projects []Project = []Project{};
	for _, proj_db := range projects_db {
		p := Project{};
		p.Project = proj_db
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