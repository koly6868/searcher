
package searcher



type NameSearchModule map[string][]*TestModel

func (sm NameSearchModule) init(data []TestModel, sortSliceFn func([]*TestModel)) {
	for i, e := range data {
		sm[e.Name] = append(sm[e.Name], &data[i])
	}

	for k := range sm {
		sortSliceFn(sm[k])
	}
}

func (sm NameSearchModule) find(q *Query) searchResult {
	return newSimpleResult(
		sm[q.Name])
}