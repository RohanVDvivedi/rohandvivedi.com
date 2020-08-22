package searchindex

import (
	"github.com/blevesearch/bleve"
	"fmt"
)

var searchProjectIndex bleve.Index;

func CreateProjectSearchIndex(indexFilename string) error {
	mapping := bleve.NewIndexMapping()
	var err error = nil
	searchProjectIndex, err = bleve.New("./db/projectSearchIndex.bleve", mapping)
	return err
}

func OpenProjectSearchIndex(indexFilename string) error {
	var err error = nil
	searchProjectIndex, err = bleve.Open(indexFilename)
	return err
}

func AddSearchIndexProjectDocument(projectName string, documentString string) {
	searchProjectIndex.Index(projectName, documentString);
}

func GetSearchResults(searchPhrase string) {
	query := bleve.NewMatchQuery("text")

	search := bleve.NewSearchRequest(query)

	searchResults, _ := searchProjectIndex.Search(search)

	fmt.Println(searchResults);
}