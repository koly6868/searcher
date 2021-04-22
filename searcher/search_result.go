package searcher_templates

import (
	"sort"
)

type searchResult interface {
	contains(*TestModel) bool
	len() int
	next() (*TestModel, error)
}

// simpleblcok block start
type simpleResult struct {
	data     []*TestModel
	pointer  int
	searchFn SliceElementSearchFN
}

func newSimpleResult(data []*TestModel, searchFn SliceElementSearchFN) *simpleResult {
	if data == nil {
		data = []*TestModel{}
	}
	return &simpleResult{
		data:     data,
		searchFn: searchFn,
	}
}

func (sr *simpleResult) contains(element *TestModel) bool {
	return sr.searchFn(sr.data, element)
}

func (sr *simpleResult) len() int {
	return len(sr.data)
}

func (sr *simpleResult) next() (*TestModel, error) {
	if sr.pointer == len(sr.data) {
		return nil, &StopIterationError{msg: "empty"}
	}
	res := sr.data[sr.pointer]
	sr.pointer++
	return res, nil
}

// simpleblcok block end

// intesectionResult block start
type intesectionResult struct {
	results  []searchResult
	length   int
	curCount int
}

func newIntesectionResult(results []searchResult) searchResult {
	length := 0
	if len(results) == 0 {
		results = []searchResult{}
	} else {
		// to keep oreder in results
		sortedResults := make([]searchResult, len(results))
		copy(sortedResults, results)

		results = sortedResults
		length = results[0].len()
		for _, e := range results {
			if e.len() < length {
				length = e.len()
			}
		}
		sort.Slice(results, func(i, j int) bool {
			return results[i].len() <= results[j].len()
		})
	}
	return &intesectionResult{
		results: results,
		length:  length,
	}
}

func (ir *intesectionResult) len() int {
	return ir.length
}

func (ir *intesectionResult) next() (*TestModel, error) {
	if ir.length == ir.curCount {
		return nil, &StopIterationError{msg: "empty"}
	}
	for {
		element, err := ir.results[0].next()
		if err != nil {
			return nil, err
		}
		elementInItersection := true
		for j := 1; j < len(ir.results); j++ {
			if !ir.results[j].contains(element) {
				elementInItersection = false
				break
			}
		}
		if elementInItersection {
			return element, nil
		}
	}
}

func (ir *intesectionResult) contains(element *TestModel) bool {
	for _, e := range ir.results {
		if !e.contains(element) {
			return false
		}
	}
	return true
}
