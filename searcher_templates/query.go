package searcher_templates

type Query struct {
	Count int
	//# {{ range .Searchers }}
	_TemplateSearcherKey _TemplateSearcherKeyType
	//# {{ end }}
}
