package searcher_templates

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
