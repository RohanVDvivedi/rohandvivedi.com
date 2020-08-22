package searchindex

import (
	"os"
	"github.com/blevesearch/bleve"
	"fmt"
)

var searchIndexesRoot = "./db/"

func createOrOpenSearchIndex(indexFilename string) (bleve.Index, error) {
	searchProjectIndex, err := bleve.Open(searchIndexesRoot + indexFilename)
	if(err == bleve.ErrorIndexPathDoesNotExist) {
		mapping := bleve.NewIndexMapping()
		searchProjectIndex, err = bleve.New(searchIndexesRoot + indexFilename, mapping)
	}
	return searchProjectIndex, err
}

func deleteSearchIndex(indexFilename string) error {
	return os.RemoveAll(searchIndexesRoot + indexFilename)
}

func getSimpleSearchQueryResults(searchIndex bleve.Index, searchPhrase string) {
	query := bleve.NewMatchQuery(searchPhrase)
	search := bleve.NewSearchRequest(query)
	searchResults, _ := searchIndex.Search(search)

	fmt.Println(searchResults);
}