package searcher_templates

//# {{ range .Searchers }}
type _TemplateSearcherName struct {
	data        map[_TemplateSearcherKeyType][]*_TemplateModelName
	sortSliceFn SliceSortFN
	searchFn    SliceElementSearchFN
}

func (sm *_TemplateSearcherName) init(data []_TemplateModelName,
	sortSliceFn SliceSortFN,
	searchFn SliceElementSearchFN) {

	sm.data = map[_TemplateSearcherKeyType][]*_TemplateModelName{}
	sm.sortSliceFn = sortSliceFn
	sm.searchFn = searchFn

	for i, e := range data {
		sm.data[e._TemplateSearcherKey] = append(sm.data[e._TemplateSearcherKey], &data[i])
	}

	for k := range sm.data {
		sortSliceFn(sm.data[k])
	}
}

func (sm *_TemplateSearcherName) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q._TemplateSearcherKey],
		sm.searchFn)
}

//# {{ end }}
