package searcher

import (
	"sort"

	"github.com/koly6868/searcher/example"
	"github.com/labstack/gommon/log"
)

type Searcher interface {
	Find(q *Query) ([]example.TestModel, error)
}

type Searher struct {
	modules []SearhModule
}

func NewSearher(data []example.TestModel,
	modules []SearhModule,
	sortSliceFn func([]*example.TestModel),
	searchFn func([]*example.TestModel, *example.TestModel) int) Searcher {
	for _, module := range modules {
		module.init(data, sortSliceFn, searchFn)
	}

	return &Searher{
		modules: modules,
	}
}

func (cs *Searher) Find(q *Query) ([]example.TestModel, error) {
	teasers := []example.TestModel{}

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
