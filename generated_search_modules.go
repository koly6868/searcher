package searcher

import "github.com/koly6868/searcher/example"

type NameSearchModule struct {
	data        map[string][]*example.TestModel
	sortSliceFn func([]*example.TestModel)
	searchFn    func([]*example.TestModel, *example.TestModel) int
}

func (sm *NameSearchModule) init(data []example.TestModel,
	sortSliceFn func([]*example.TestModel),
	searchFn func([]*example.TestModel, *example.TestModel) int) {

	sm.data = map[string][]*example.TestModel{}
	sm.sortSliceFn = sortSliceFn
	sm.searchFn = searchFn

	for i, e := range data {
		sm.data[e.Name] = append(sm.data[e.Name], &data[i])
	}

	for k := range sm.data {
		sortSliceFn(sm.data[k])
	}
}

func (sm *NameSearchModule) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q.Name],
		sm.searchFn)
}

type AgeSearchModule struct {
	data        map[int][]*example.TestModel
	sortSliceFn func([]*example.TestModel)
	searchFn    func([]*example.TestModel, *example.TestModel) int
}

func (sm *AgeSearchModule) init(data []example.TestModel,
	sortSliceFn func([]*example.TestModel),
	searchFn func([]*example.TestModel, *example.TestModel) int) {

	sm.data = map[int][]*example.TestModel{}
	sm.sortSliceFn = sortSliceFn
	sm.searchFn = searchFn

	for i, e := range data {
		sm.data[e.Age] = append(sm.data[e.Age], &data[i])
	}

	for k := range sm.data {
		sortSliceFn(sm.data[k])
	}
}

func (sm *AgeSearchModule) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q.Age],
		sm.searchFn)
}
