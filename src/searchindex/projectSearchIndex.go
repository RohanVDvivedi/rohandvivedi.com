package searchindex

import (
	"strings"
	//"fmt"
)

import (
	"github.com/blevesearch/bleve"
	"rohandvivedi.com/src/data"
	"rohandvivedi.com/src/githubsync"
)

var projectSearchIndex bleve.Index;
var indexOpen = false;

type projectSearchIndexObject struct {
	ProjectName string
	GithubRepositoryName string
	Description string
	Categories []string
	ProgrammingLanguages []string
	LibrariesBeingUsed []string
	SkillSetsAcquired []string
	ReadmeFiles map[string]string
}

func InitProjectSearchIndex(projectSearchIndexPath string) {
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
	r := proj_db.GetProjectGithubRepositoryLink();
	o := proj_db.GetProjectOwner();

	p := projectSearchIndexObject{};

	p.ProjectName = proj_db.Name.String
	if(r!=nil) {
		p.GithubRepositoryName = r.Name.String
	}
	p.Description = proj_db.Descr.String

	p.Categories = []string{};
	proj_db_categories := proj_db.GetProjectCategories();
	for _, proj_db_category := range proj_db_categories {
		if(proj_db_category.Category.Valid) {
			p.Categories = append(p.Categories, proj_db_category.Category.String)
		}
	}

	if(proj_db.ProgrLangs.Valid) {
		p.ProgrammingLanguages = strings.Split(proj_db.ProgrLangs.String, ",")
	}

	if(proj_db.LibsUsed.Valid) {
		p.LibrariesBeingUsed = strings.Split(proj_db.LibsUsed.String, ",")
	}

	if(proj_db.SkillSets.Valid) {
		p.SkillSetsAcquired = strings.Split(proj_db.SkillSets.String, ",")
	}

	p.ReadmeFiles = map[string]string{}

	if(o != nil) {
		projectOwnerGithubSocials := o.FindSocialsOfType("github")
		if(len(projectOwnerGithubSocials) > 0) {

			proj_db_hyperlinks := proj_db.GetProjectHyperlinks();
			for _, proj_db_hyperlink := range proj_db_hyperlinks {
				if(proj_db_hyperlink.Name.Valid && proj_db_hyperlink.LinkType.Valid && proj_db_hyperlink.LinkType.String == "GITHUB"){
					readmeContent, err := githubsync.GetGithubFile(projectOwnerGithubSocials[0].Username.String, proj_db_hyperlink.Name.String, "README.md")
					if(err == nil) {
						p.ReadmeFiles[proj_db_hyperlink.Name.String] = readmeContent
					}
				}
			}

		}
	}

	//fmt.Println(p)

	projectSearchIndex.Delete(p.ProjectName)
	projectSearchIndex.Index(p.ProjectName, p)
}

func GetProjectSearchQueryResults(queryString string) []string {
	return getSimpleSearchQueryResults(projectSearchIndex, queryString)
}