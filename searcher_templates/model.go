package searcher_templates

type _TemplateModelName struct {
	//# {{ range  .Searchers }}
	_TemplateKey _TemplateKeyType
	//# {{ end }}
}

//# {{ range  .Searchers }}
type _TemplateKeyType string

//# {{ end }}
