package main

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

func TestFillTemplate(t *testing.T) {
	want := `package searcher_templates
        
        	
	type Name struct {
		data        map[string][]*TestModel
		sortSliceFn SliceSortFN
		searchFn    SliceElementSearchFN
	};`
	s := `package searcher_templates

	{{ range .Searchers }}
   type {{ .Name }} struct {
	   data        map[{{ .KeyType }}][]*{{ $.ModelName }}
	   sortSliceFn SliceSortFN
	   searchFn    SliceElementSearchFN
   }{{ end }}`

	tpl := template.Must(template.New("ds").Parse(s))
	gotB := &bytes.Buffer{}
	err := tpl.Execute(gotB, SearcherConfig{
		Searchers: []SeacrherModuleConfig{
			{
				Name:    "Name",
				KeyType: "string",
				Key:     "Name",
			},
		},
		ModelName: "TestModel",
	})
	if err != nil {
		t.Error(err)
	}

	got := gotB.String()
	if got != want {
		t.Errorf("want: %s;\n got: %s", want, got)
	}
}

func TestPreprocessCodeTemplate(t *testing.T) {
	s := `package searcher_templates

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
	
	`
	res := preprocessCodeTemplate(s)

	fmt.Println(res)
	t.Error("err")
}

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
	res := preprocessRangeStatements(s)

	fmt.Println(res)
	t.Error("err")
}

func TestPrerpocessRangeBodyTemplate(t *testing.T) {
	s := `//# {{ range .Searchers }}
		_TemplateSearcherKey _TemplateSearcherKeyType
		//# {{ end }}`

	res := prerpocessTemplateVars(s, "")

	fmt.Println(res)
	t.Error("err")
}
