package main

import (
	"strings"
	"testing"
)

func TestExtractRangeVar(t *testing.T) {
	want := "Searchers"
	s := `//# {{ range  .Searchers }}
		_TemplateSearcherKey _TemplateSearcherKeyType
		//# {{ end }}`

	got := extractRangeIterVariable(s)

	if got != want {
		t.Errorf("want: %s;\n got: %s", want, got)
	}
}

func TestPreprocessRangeStatements(t *testing.T) {
	s := `package searcher_templates

	type Query struct {
		Count int
		//# {{ range $key, $value := .Searchers }}
		_TemplateSearcherKey _TemplateSearcherKeyType
		//# {{ end }}
	}
	`
	want := `package searcher_templates

	type Query struct {
		Count int
		//# {{ range $key, $value := .Searchers }}
		{{ .Key }} {{ .KeyType }}
		//# {{ end }}
	}`

	got := preprocessRangeStatements(s)
	if strings.TrimSpace(want) != strings.TrimSpace(got) {
		t.Errorf("want: %s;\n got: %s", want, got)
	}

}

//func TestPreprocessCodeTemplate(t *testing.T) {
//	s := `package searcher_templates
//
//	//# {{ range .Searchers }}
//	type _TemplateSearcherName struct {
//		data        map[_TemplateSearcherKeyType][]*_TemplateModelName
//		sortSliceFn SliceSortFN
//		searchFn    SliceElementSearchFN
//	}
//
//	func (sm *_TemplateSearcherName) init(data []_TemplateModelName,
//		sortSliceFn SliceSortFN,
//		searchFn SliceElementSearchFN) {
//
//		sm.data = map[_TemplateSearcherKeyType][]*_TemplateModelName{}
//		sm.sortSliceFn = sortSliceFn
//		sm.searchFn = searchFn
//
//		for i, e := range data {
//			sm.data[e._TemplateSearcherKey] = append(sm.data[e._TemplateSearcherKey], &data[i])
//		}
//
//		for k := range sm.data {
//			sortSliceFn(sm.data[k])
//		}
//	}
//
//	func (sm *_TemplateSearcherName) find(q *Query) searchResult {
//		return newSimpleResult(
//			sm.data[q._TemplateSearcherKey],
//			sm.searchFn)
//	}
//
//	//# {{ end }}
//	`
//
//	want := `package searcher_templates
//
//	{{ range .Searchers }}
//   type {{ .Name }} struct {
//	   data        map[{{ .KeyType }}][]*{{ $.ModelName }}
//	   sortSliceFn SliceSortFN
//	   searchFn    SliceElementSearchFN
//   }
//
//   func (sm *{{ .Name }}) init(data []{{ $.ModelName }},
//	   sortSliceFn SliceSortFN,
//	   searchFn SliceElementSearchFN) {
//
//	   sm.data = map[{{ .KeyType }}][]*{{ $.ModelName }}{}
//	   sm.sortSliceFn = sortSliceFn
//	   sm.searchFn = searchFn
//
//	   for i, e := range data {
//		   sm.data[e.{{ .Key }}] = append(sm.data[e.{{ .Key }}], &data[i])
//	   }
//
//	   for k := range sm.data {
//		   sortSliceFn(sm.data[k])
//	   }
//   }
//
//   func (sm *{{ .Name }}) find(q *Query) searchResult {
//	   return newSimpleResult(
//		   sm.data[q.{{ .Key }}],
//		   sm.searchFn)
//   }
//
//	{{ end }}`
//
//	got := preprocessCodeTemplate(s)
//	if strings.TrimSpace(want) != strings.TrimSpace(got) {
//		t.Errorf("want: %s;\n got: %s", want, got)
//	}
//}

//func TestPrerpocessRangeBodyTemplate(t *testing.T) {
//	s := `//# {{ range .Searchers }}
//		_TemplateSearcherKey _TemplateSearcherKeyType
//		//# {{ end }}`
//
//	want := `//# {{ range .Searchers }}
// {{ $.SearcherKey }} {{ $.SearcherKeyType }}
// //# {{ end }}`
//	got := prerpocessTemplateVars(s, "")
//	if strings.TrimSpace(want) != strings.TrimSpace(got) {
//		t.Errorf("want: %s\n got: %s", want, got)
//	}
//}
