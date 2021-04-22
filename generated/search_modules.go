package searcher_templates

 {{ range .Searchers }}
type {{ .SearcherName }} struct {
	data        map[_TemplateKeyType][]*{{ .ModelName }}
	sortSliceFn SliceSortFN
	searchFn    SliceElementSearchFN
}

func (sm *{{ .SearcherName }}) init(data []{{ .ModelName }},
	sortSliceFn SliceSortFN,
	searchFn SliceElementSearchFN) {

	sm.data = map[_TemplateKeyType][]*{{ .ModelName }}{}
	sm.sortSliceFn = sortSliceFn
	sm.searchFn = searchFn

	for i, e := range data {
		sm.data[e._TemplateKey] = append(sm.data[e._TemplateKey], &data[i])
	}

	for k := range sm.data {
		sortSliceFn(sm.data[k])
	}
}

func (sm *{{ .SearcherName }}) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q._TemplateKey],
		sm.searchFn)
}

 {{ end }}
