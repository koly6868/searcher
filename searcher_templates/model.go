package searcher_templates

type _TemplateModelName struct {
	//# {{ range  .Searchers }}
	_TemplateKey _TemplateKeyType
	//# {{ end }}
}

type _TemplateKeyType string
