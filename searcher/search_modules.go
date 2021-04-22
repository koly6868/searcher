package searcher_templates

type Name struct {
	data        map[String][]*TestModel
	sortSliceFn SliceSortFN
	searchFn    SliceElementSearchFN
}

func (sm *Name) init(data []TestModel,
	sortSliceFn SliceSortFN,
	searchFn SliceElementSearchFN) {

	sm.data = map[String][]*TestModel{}
	sm.sortSliceFn = sortSliceFn
	sm.searchFn = searchFn

	for i, e := range data {
		sm.data[e.Name] = append(sm.data[e.Name], &data[i])
	}

	for k := range sm.data {
		sortSliceFn(sm.data[k])
	}
}

func (sm *Name) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q.Name],
		sm.searchFn)
}

type Age struct {
	data        map[Int][]*TestModel
	sortSliceFn SliceSortFN
	searchFn    SliceElementSearchFN
}

func (sm *Age) init(data []TestModel,
	sortSliceFn SliceSortFN,
	searchFn SliceElementSearchFN) {

	sm.data = map[Int][]*TestModel{}
	sm.sortSliceFn = sortSliceFn
	sm.searchFn = searchFn

	for i, e := range data {
		sm.data[e.Age] = append(sm.data[e.Age], &data[i])
	}

	for k := range sm.data {
		sortSliceFn(sm.data[k])
	}
}

func (sm *Age) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q.Age],
		sm.searchFn)
}
