package searcher

import (
	"sort"

	"github.com/labstack/gommon/log"
)

type CreativeSearcher interface {
	Find(q *Query) ([]TestModel, error)
}

type CreativeSearh struct {
	modules []searhModule
}

func NewCreativeSearh(data []TestModel) CreativeSearcher {
	modules := []searhModule{}

	for _, module := range modules {
		module.init(data)
	}

	return &CreativeSearh{
		modules: modules,
	}
}

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
