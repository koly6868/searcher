package _TemplatePackageName

type Query struct {
	Count int
	//# {{ range .Searchers }}
	_TemplateSearcherKey _TemplateSearcherKeyType
	//# {{ end }}
}
