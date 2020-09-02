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
		projdb = &data.Project{
			Name: data.NewNullString(projGithub.Name),
			Descr: data.NewNullString(projGithub.Description),
			ProjectOwner: data.NewNullInt64(1),
		}
		data.InsertProject(projdb)
	} else {
		// update project
		projdb.Name = data.NewNullString(projGithub.Name);
		projdb.Descr = data.NewNullString(projGithub.Description);
		projdb.ProjectOwner = data.NewNullInt64(1);
		data.UpdateProject(projdb)
		/*proj_hyperlinkdb := projdb.GetProjectGithubRepositoryLink())
		if(proj_hyperlinkdb == nil) {
			// insert github link
		} else {
			// update github link
		}*/
	}

	fmt.Println(projdb)
}

