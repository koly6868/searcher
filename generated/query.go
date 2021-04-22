package searcher_templates

type Query struct {
	Count int
	 {{ range .Searchers }}
	{{ .Key }} {{ .KeyType }}
	 {{ end }}
}
