package searcher

type nameSearchModule struct {
	data        map[string][]*TestModel
	sortSliceFn func([]*TestModel)
}

func NewNameSearchModule(sortSliceFn func([]*TestModel)) searhModule {
	return &nameSearchModule{
		sortSliceFn: sortSliceFn,
		data:        map[string][]*TestModel{},
	}
}

func (sm nameSearchModule) init(data []TestModel) {
	for i, e := range data {
		sm.data[e.Name] = append(sm.data[e.Name], &data[i])
	}

	for k := range sm.data {
		sm.sortSliceFn(sm.data[k])
	}
}

func (sm nameSearchModule) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q.Name])
}
