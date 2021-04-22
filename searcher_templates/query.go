package searcher_templates

type Query struct {
	Count int
	//# {{ range .Searchers }}
	_TemplateKey _TemplateKeyType
	//# {{ end }}
}
