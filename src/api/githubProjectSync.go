package api

import (
	"fmt"
	"net/http"
	"rohandvivedi.com/src/githubsync"
)

// api handlers in this file
var SyncProjectFromGithubRepository = http.HandlerFunc(githubRepositorySyncUp)

func githubRepositorySyncUp(w http.ResponseWriter, r *http.Request) {
	projectName, exists_name := r.URL.Query()["name"];
	if exists_name {
		fmt.Println(githubsync.GetGithubProject("RohanVDvivedi", projectName[0]));
	}
}

