package searcher

import (
	"bytes"
	"os"
	"path"
	"text/template"

	"golang.org/x/tools/imports"
)

const creativeSearcherPath = "generated_creative_search.go"
const creativeSearcherTemplate = `
// CreativeSearcher ...
type CreativeSearcher interface {
	Find(q *Query) ([]TestModel, error)
}

// CreativeSearh ...
type CreativeSearh struct {
	modules []searhModule
}

// NewCreativeSearh creates search system.
// data should no be changed after initialization.
func NewCreativeSearh(data []TestModel) CreativeSearcher {
	modules := []searhModule{
	}

	for _, module := range modules {
		module.init(data)
	}

	return &CreativeSearh{
		modules: modules,
	}
}

// Find ...
func (cs *CreativeSearh) Find(q *Query) ([]TestModel, error) {
	teasers := []TestModel{}

	if cs.modules == nil || len(cs.modules) == 0 {
		return teasers, &CreativeSearhInitializationError{msg: "no modules"}
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
		creative, err := results[0].next()
		if err != nil {
			// early tteration stop
			if _, ok := err.(*StopIterationError); ok {
				break
			}

			log.Errorf("creative search error. err : %s", err.Error())

		}
		appropiateCreative := true
		for j := 1; j < len(results); j++ {
			if !results[j].contains(creative) {
				appropiateCreative = false
				break
			}
		}

		if appropiateCreative {
			teasers = append(teasers, *creative)
		}
	}

	return teasers, nil
}

type searhModule interface {
	init([]TestModel)
	find(*Query) searchResult
}
`

func GenCreativeSearcher(gd *GenData, basePath string) error {
	f, err := os.Create(path.Join(basePath, creativeSearcherPath))
	if err != nil {
		return err
	}

	data, err := genCreativeSearcher(gd)
	if err != nil {
		return err
	}

	_, err = f.WriteString(gencCodeHeader)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	f.Close()
	imports.Process(path.Join(basePath, creativeSearcherPath), nil, &imports.Options{
		FormatOnly: false,
	})
	return err
}

func genCreativeSearcher(gd *GenData) ([]byte, error) {
	t := template.Must(template.New("creative_search").Parse(creativeSearcherTemplate))
	b := &bytes.Buffer{}
	err := t.Execute(b, gd)

	return b.Bytes(), err
}
