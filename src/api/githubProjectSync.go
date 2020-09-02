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
		if(err != nil) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No such Github repository exists for the owner on Github"))
		} else {
			projdb := data.GetProjectByName(projectName[0])
			if(projdb == nil) {
				// insert
				fmt.Println("inserting project");
				fmt.Println(projGithub);
			} else {
				// update
				fmt.Println("updating project");
				fmt.Println(projdb);
				fmt.Println("using");
				fmt.Println(projGithub);
				fmt.Println("repo link = ")
				fmt.Println(projdb.GetProjectGithubRepositoryLink())
			}
		}
	}
}

