package searcher_templates

import (
	"os"

	"golang.org/x/tools/imports"
)

const queryAnyValue = -1

type SliceSortFN func([]*TestModel)
type SliceElementSearchFN func([]*TestModel, *TestModel) bool
type SearhModule interface {
	init(
		[]TestModel,
		SliceSortFN,
		SliceElementSearchFN,
	)
	find(*Query) searchResult
}

// Searcher ...
type Searcher interface {
	Find(q *Query) ([]TestModel, error)
}

type GenData struct {
	Searchers []struct {
		Name    string `json:"Name"`
		KeyType string `json:"KeyType"`
		Key     string `json:"Key"`
	} `json:"searchers"`
	ModelName string `json:"ModelName"`
}

func FormatAndWrite(path string, code []byte) error {
	code, err := imports.Process(path, code, &imports.Options{
		FormatOnly: false,
	})
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(code)

	return err
}

const pkg = "searcher"
const gencCodeHeader = `
package searcher_templates


`
