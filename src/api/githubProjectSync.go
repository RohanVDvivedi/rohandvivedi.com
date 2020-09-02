package api

import (
	"fmt"
	"net/http"
	"rohandvivedi.com/src/githubsync"
	"rohandvivedi.com/src/data"
)

// api handlers in this file
var SyncProjectFromGithubRepository = http.HandlerFunc(githubRepositorySyncUp)

func githubRepositorySyncUp(w http.ResponseWriter, r *http.Request) {
	projectName, exists_name := r.URL.Query()["name"];
	if !exists_name {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You must provide name of the project"))
		return
	}

	// find appropriate project on gihub of the owner
	projGithub, err := githubsync.GetGithubProject("RohanVDvivedi", projectName[0]);
	if(err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No such Github repository exists for the owner on Github"))
		return
	}

	// update or insert project details of owner, from the github api call
	projdb := data.GetProjectByName(projectName[0])
	if(projdb == nil) {
		// insert project and project_hyperlink of the github repository
	} else {
		// update project
		proj_hyperlinkdb := projdb.GetProjectGithubRepositoryLink())
		if(proj_hyperlinkdb == nil) {
			// insert github link
		} else {
			// update github link
		}
	}
}

