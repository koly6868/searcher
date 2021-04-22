package searcher_templates

const queryAnyValue = -1

type SliceSortFN func([]*_TemplateModelName)
type SliceElementSearchFN func([]*_TemplateModelName, *_TemplateModelName) bool
type SearhModule interface {
	init(
		[]_TemplateModelName,
		SliceSortFN,
		SliceElementSearchFN,
	)
	find(*Query) searchResult
}

// Searcher ...
type Searcher interface {
	Find(q *Query) ([]_TemplateModelName, error)
}
