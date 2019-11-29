package searcher_procs

import (
	"fmt"
	"github.com/blevesearch/bleve"
	_ "github.com/blevesearch/bleve/analysis/lang/cjk"
	"os"
	"testing"
)

//⊕ [Bleve with Japanese](https://gist.github.com/mosuka/46f5f22cab389e559fb1)
//⊕ [bleve/analyzer_cjk_test.go at master · blevesearch/bleve](https://github.com/blevesearch/bleve/blob/master/analysis/lang/cjk/analyzer_cjk_test.go)
func TestBleve_ja(test *testing.T) {
	indexMapping := bleve.NewIndexMapping()
	docMapping := bleve.NewDocumentMapping()

	fieldMapping := bleve.NewTextFieldMapping()
	fieldMapping.Analyzer = "cjk"

	docMapping.AddFieldMappingsAt("Body", fieldMapping)

	indexMapping.AddDocumentMapping("Example", docMapping)

	// create index
	var index bleve.Index
	var err error
	var indexPath="example_ja.bleve"

	_, err = os.Stat(indexPath)
	if err == nil {
		index, err = bleve.Open(indexPath)
	} else {
		index, err = bleve.New(indexPath, indexMapping)
	}
	if err != nil {
		fmt.Println("open index error:", err)
		return
	}

	// index data
	data := struct {
		Body string
	}{
		Body: "東京都港区六本木",
	}

	index.Index("000", data)  // expect it to be tokeinzed "東京","都","港","区" and "六本木" from "東京都港区六本木"

	// search index
	query := bleve.NewMatchQuery("京都") // search for "京都"
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println("a error:", err)
		return
	}

	// show results
	fmt.Println(searchResults) // expect "No matches", but "京都" will match
}