package searcher

type GenData struct {
	Searchers []struct {
		Name    string `json:"Name"`
		KeyType string `json:"KeyType"`
		Key     string `json:"Key"`
	} `json:"searchers"`
	ModelName string `json:"ModelName"`
}

const pkg = "searcher"
const gencCodeHeader = `
package searcher


`
