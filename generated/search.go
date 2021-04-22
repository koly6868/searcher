package searcher_templates

import (
	"sort"

	"github.com/labstack/gommon/log"
)

// searher ...
type searher struct {
	modules []SearhModule
}

// NewSearher creates search system.
// data should no be changed after initialization.
func NewSearher(data []_TemplateModelName,
	modules []SearhModule,
	sortSliceFn func([]*_TemplateModelName),
	searchFn SliceElementSearchFN) Searcher {
	for _, module := range modules {
		module.init(data, sortSliceFn, searchFn)
	}

	return &searher{
		modules: modules,
	}
}

// Find ...
func (cs *searher) Find(q *Query) ([]_TemplateModelName, error) {
	teasers := []_TemplateModelName{}

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
