package api

import (
	"net/http"
	"rohandvivedi.com/src/githubsync"
	"rohandvivedi.com/src/data"
	"strings"
)

// api handlers in this file
var SyncProjectFromGithubRepository = http.HandlerFunc(githubRepositorySyncUp)

func githubRepositorySyncUp(w http.ResponseWriter, r *http.Request) {
	projectName, exists_name := r.URL.Query()["name"];
	if !exists_name || strings.Contains(projectName[0], " ") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You must provide name of the project as in the github repository"))
		return
	}

	// find appropriate project on gihub of the owner
	projGithub, err := githubsync.GetGithubProject("RohanVDvivedi", projectName[0]);
	if(err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No such Github repository exists for the owner on Github"))
		return
	}

	// this is the name that goes in the name of the project in project table
	projectNameDb := strings.Replace(projectName[0], "-", " ", -1)

	// update or insert project details, from the github api call
	projdb := data.GetProjectByName(projectNameDb)
	if(projdb == nil) {
		projdb = &data.Project{
			Name: data.NewNullString(projectNameDb),
			Descr: data.NewNullString(projGithub.Description),
			ProjectOwner: data.NewNullInt64(1),
		}
		data.InsertProject(projdb)
	} else {
		projdb.Name = data.NewNullString(projectNameDb);
		projdb.Descr = data.NewNullString(projGithub.Description);
		projdb.ProjectOwner = data.NewNullInt64(1);
		data.UpdateProject(projdb)
	}

	// update or insert project github repository link, from the github api call
	proj_hyperlinkdb := projdb.GetProjectGithubRepositoryLink()
	if(proj_hyperlinkdb == nil) {
		proj_hyperlinkdb = &data.ProjectHyperlink{
			Name: data.NewNullString(projGithub.Name),
			Href: data.NewNullString(projGithub.HTMLURL),
			Descr: data.NewNullString("Github repository of the project"),
			LinkType: data.NewNullString("GITHUB"),
			ProjectId: projdb.Id,
		}
		data.InsertProjectHyperlink(proj_hyperlinkdb)
	} else {
		proj_hyperlinkdb.Name = data.NewNullString(projGithub.Name);
		proj_hyperlinkdb.Href = data.NewNullString(projGithub.HTMLURL);
		proj_hyperlinkdb.Descr = data.NewNullString("Github repository of the project");
		proj_hyperlinkdb.LinkType = data.NewNullString("GITHUB");
		proj_hyperlinkdb.ProjectId = projdb.Id;
		data.UpdateProjectHyperlink(proj_hyperlinkdb)
	}
}

