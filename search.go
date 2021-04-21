package searcher

import (
	"bytes"
	"path"
	"text/template"
)

const SearcherPath = "generated_search.go"
const SearcherTemplate = `
// Searcher ...
type Searcher interface {
	Find(q *Query) ([]{{.ModelName}}, error)
}

// Searher ...
type Searher struct {
	modules []SearhModule
}

// NewSearher creates search system.
// data should no be changed after initialization.
func NewSearher(data []{{.ModelName}},
	modules []SearhModule,
	sortSliceFn func([]*{{.ModelName}}),
	searchFn func([]*example.TestModel, *example.TestModel) int) Searcher {
	for _, module := range modules {
		module.init(data, sortSliceFn, searchFn)
	}

	return &Searher{
		modules: modules,
	}
}

// Find ...
func (cs *Searher) Find(q *Query) ([]{{.ModelName}}, error) {
	teasers := []{{.ModelName}}{}

	if cs.modules == nil || len(cs.modules) == 0 {
		return teasers, &SearherInitializationError{msg: "no modules"}
	}

	results := []searchResult{}
	for _, module := range cs.modules {
		results = append(results, module.find(q))
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].len() <= results[j].len()
	})

	n := results[0].len()
	if n > q.Count {
		n = q.Count
	}
	for i := 0; i < n; {
		element, err := results[0].next()
		if err != nil {
			// early tteration stop
			if _, ok := err.(*StopIterationError); ok {
				break
			}

			log.Errorf("element search error. err : %s", err.Error())

		}
		appropiateElement := true
		for j := 1; j < len(results); j++ {
			if !results[j].contains(element) {
				appropiateElement = false
				break
			}
		}

		if appropiateElement {
			teasers = append(teasers, *element)
		}
	}

	return teasers, nil
}

type SearhModule interface {
	init(
		[]example.TestModel,
		func([]*example.TestModel),
		func([]*example.TestModel, *example.TestModel) int,
	)
	find(*Query) searchResult
}
`

func GenSearcher(gd *GenData, basePath string) error {
	codeBuff := &bytes.Buffer{}
	data, err := genSearcher(gd)
	if err != nil {
		return err
	}

	codeBuff.WriteString(gencCodeHeader)
	codeBuff.Write(data)

	err = FormatAndWrite(path.Join(basePath, SearcherPath), codeBuff.Bytes())

	return err
}

func genSearcher(gd *GenData) ([]byte, error) {
	t := template.Must(template.New("element_search").Parse(SearcherTemplate))
	b := &bytes.Buffer{}
	err := t.Execute(b, gd)

	return b.Bytes(), err
}
