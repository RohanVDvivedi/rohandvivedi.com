package searchindex

import (
	"github.com/blevesearch/bleve"
	"rohandvivedi.com/src/data"
	"rohandvivedi.com/src/githubsync"
)

var projectSearchIndexPath = "project_search.bleve"
var projectSearchIndex bleve.Index;
var indexOpen = false;

type projectSearchIndexObject struct {
	Name string
	Description string
	Categories []string
	ReadmeFiles map[string]string
}

func InitProjectSearchIndex() {
	if(!indexOpen) {
		idx, err := createOrOpenSearchIndex(projectSearchIndexPath);
		if(err == nil) {
			projectSearchIndex = idx
			indexOpen = true
		}
	}
}

func InsertProjectInSearchIndex_ByName(projectName string) {
	proj_db := data.GetProjectByName(projectName)
	if(proj_db == nil) {
		return
	}

	InsertProjectInSearchIndex(proj_db)
}

func InsertAllProjectsInSearchIndex() {
	proj_dbs := data.GetAllProjects()
	for _, proj_db := range proj_dbs {
		InsertProjectInSearchIndex(&proj_db)
	}
}

func InsertProjectInSearchIndex(proj_db *data.Project) {
	p := projectSearchIndexObject{};

	p.Name = proj_db.Name.String
	p.Description = proj_db.Descr.String

	p.Categories = []string{};

	proj_db_categories := proj_db.GetProjectCategories();
	for _, proj_db_category := range proj_db_categories {
		if(proj_db_category.Category.Valid) {
			p.Categories = append(p.Categories, proj_db_category.Category.String)
		}
	}

	p.ReadmeFiles = map[string]string{}

	proj_db_hyperlinks := proj_db.GetProjectHyperlinks();
	for _, proj_db_hyperlink := range proj_db_hyperlinks {
		if(proj_db_hyperlink.Name.Valid && proj_db_hyperlink.LinkType.Valid && proj_db_hyperlink.LinkType.String == "GITHUB"){
			readmeContent, err := githubsync.GetGithubFile("RohanVDvivedi", proj_db_hyperlink.Name.String, "README.md")
			if(err == nil) {
				p.ReadmeFiles[proj_db_hyperlink.Name.String] = readmeContent
			}
		}
	}

	projectSearchIndex.Delete(p.Name)
	projectSearchIndex.Index(p.Name, p)
}

func GetProjectSearchQueryResults(queryString string) []string {
	return getSimpleSearchQueryResults(projectSearchIndex, queryString)
}