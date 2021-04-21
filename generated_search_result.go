package searcher

import (
	"sort"

	"github.com/koly6868/searcher/example"
)

type searchResult interface {
	contains(*example.TestModel) bool
	len() int
	next() (*example.TestModel, error)
}

type simpleResult struct {
	data     []*example.TestModel
	pointer  int
	searchFn func([]*example.TestModel, *example.TestModel) int
}

func newSimpleResult(data []*example.TestModel, searchFn func([]*example.TestModel, *example.TestModel) int) *simpleResult {
	if data == nil {
		data = []*example.TestModel{}
	}
	return &simpleResult{
		data:     data,
		searchFn: searchFn,
	}
}

func (sr *simpleResult) contains(element *example.TestModel) bool {
	index := sr.searchFn(sr.data, element)

	return (index < len(sr.data)) && (sr.data[index].ID == element.ID)
}

func (sr *simpleResult) len() int {
	return len(sr.data)
}

func (sr *simpleResult) next() (*example.TestModel, error) {
	if sr.pointer == len(sr.data) {
		return nil, &StopIterationError{msg: "empty"}
	}
	res := sr.data[sr.pointer]
	sr.pointer++
	return res, nil
}

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

func (ir *intesectionResult) next() (*example.TestModel, error) {
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

func (ir *intesectionResult) contains(element *example.TestModel) bool {
	for _, e := range ir.results {
		if !e.contains(element) {
			return false
		}
	}
	return true
}
