package main

type SearcherConfig struct {
	Searchers []SeacrherModuleConfig `json:"searchers"`
	ModelName string                 `json:"ModelName"`
}

type SeacrherModuleConfig struct {
	Name    string `json:"Name"`
	KeyType string `json:"KeyType"`
	Key     string `json:"Key"`
}
