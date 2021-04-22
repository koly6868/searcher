package searcher_templates

type _TemplateModelName struct {
	 {{ range  .Searchers }}
	{{ .Key }} {{ .KeyType }}
	 {{ end }}
}

type _TemplateKeyType string
