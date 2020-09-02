package api

import (
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

	// update or insert project details, from the github api call
	projdb := data.GetProjectByName(projectName[0])
	if(projdb == nil) {
		projdb = &data.Project{
			Name: data.NewNullString(projGithub.Name),
			Descr: data.NewNullString(projGithub.Description),
			ProjectOwner: data.NewNullInt64(1),
		}
		data.InsertProject(projdb)
	} else {
		projdb.Name = data.NewNullString(projGithub.Name);
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

