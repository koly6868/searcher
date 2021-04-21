package searcher

import (
	"bytes"
	"path"
	"text/template"
)

const serachResulPath = "generated_search_result.go"
const serachResulTemplate = `
type searchResult interface {
	contains(*{{.ModelName}}) bool
	len() int
	next() (*{{.ModelName}}, error)
}

// simpleblcok block start
type simpleResult struct {
	data    []*{{.ModelName}}
	pointer int
	searchFn func([]*example.TestModel, *example.TestModel) int
}

func newSimpleResult(data []*{{.ModelName}}, searchFn func([]*example.TestModel, *example.TestModel) int) *simpleResult {
	if data == nil {
		data = []*{{.ModelName}}{}
	}
	return &simpleResult{
		data: data,
		searchFn : searchFn,
	}
}

func (sr *simpleResult) contains(element *{{.ModelName}}) bool {
	index := sr.searchFn(sr.data, element)

	return (index < len(sr.data)) && (sr.data[index].ID == element.ID)
}

func (sr *simpleResult) len() int {
	return len(sr.data)
}

func (sr *simpleResult) next() (*{{.ModelName}}, error) {
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

func (ir *intesectionResult) next() (*{{.ModelName}}, error) {
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

func (ir *intesectionResult) contains(element *{{.ModelName}}) bool {
	for _, e := range ir.results {
		if !e.contains(element) {
			return false
		}
	}
	return true
}
`

func GenElementSerachResult(gd *GenData, basePath string) error {
	b := &bytes.Buffer{}
	b.WriteString(gencCodeHeader)

	t := template.Must(template.New("serachResultTemplate").Parse(serachResulTemplate))
	err := t.Execute(b, gd)
	if err != nil {
		return err
	}
	return FormatAndWrite(path.Join(basePath, serachResulPath), b.Bytes())
}
