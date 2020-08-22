package searchindex

import (
	"github.com/blevesearch/bleve"
	"fmt"
	"rohandvivedi.com/src/data"
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

func InsertProjectInSearchIndex(projectName string) {
	proj_db := data.GetProjectByName(projectName)
	if(proj_db == nil) {
		return
	}

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
			readmeContent, err := getGithubFile("RohanVDvivedi", proj_db_hyperlink.Name.String, "README.md")
			if(err == nil) {
				p.ReadmeFiles[proj_db_hyperlink.Name.String] = readmeContent
			}
		}
	}

	fmt.Println(p)

	//projectSearchIndex.Delete(p.name)
	projectSearchIndex.Index(p.Name, p)
}

func GetProjectSearchQueryResults(queryString string) []string {
	getSimpleSearchQueryResults(projectSearchIndex, queryString)
	return []string{}
}