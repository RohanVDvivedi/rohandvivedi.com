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
	if exists_name {
		projGithub, err := githubsync.GetGithubProject("RohanVDvivedi", projectName[0]);
		if(err) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No such Github repository exists for the owner"))
		} else {
			projdb := data.GetProjectByName(projectName[0])
			if(projdb == nil) {
				// insert
			} else {
				// update
			}
		}
	}
}

