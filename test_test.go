package main

import (
	"fmt"
	"os"
	"regexp"
	"testing"
	"text/template"
)

func TestFillTemplate(t *testing.T) {
	s := `package searcher_templates

	type Query struct {
		Count int
		 {{ range .Searchers }}
		{{ .Key }} {{ .KeyType }}
		 {{ end }}
	}`
	tpl := template.Must(template.New("ds").Parse(s))
	tpl.Execute(os.Stdout, SearcherConfig{
		Searchers: []SeacrherModuleConfig{
			{
				KeyType: "string",
				Key:     "Name",
			},
		},
	})

	t.Error("err")
}

func TestPreprocessCodeTemplate(t *testing.T) {
	s := `package searcher_templates

	type Query struct {
		Count int
		//# {{ range $key, $value := .searchers }}
		_TemplateKey _TemplateKeyType
		//# {{ end }}
	}
	`
	res := preprocessCodeTemplate(s)

	fmt.Println(res)
	t.Error("err")
}

func TestPreprocessRangeStatements(t *testing.T) {
	s := `package searcher_templates

	type Query struct {
		Count int
		//# {{ range $key, $value := .Searchers }}
		_TemplateKey _TemplateKeyType
		//# {{ end }}
	}
	`
	res := preprocessRangeStatements(s)

	fmt.Println(res)
	t.Error("err")
}

func TestPrerpocessRangeBodyTemplate(t *testing.T) {
	s := `//# {{ range .Searchers }}
		_TemplateKey _TemplateKeyType
		//# {{ end }}`

	res := prerpocessRangeBodyTemplate(s)

	fmt.Println(res)
	t.Error("err")
}
func TestExtractRangeKeyValue(t *testing.T) {
	s := `//# {{ range $key, $value := .Searchers }}
		_TemplateKey _TemplateKeyType
		//# {{ end }}`

	res := extractRangeKeyValue(s)
	for _, e := range res {
		fmt.Println(s[e[0]:e[1]])
	}
	fmt.Printf("%#v\n", res)
	t.Error("err")
}

func TestTemp(t *testing.T) {
	s := `{{ range .Vals}}
		{{.Name}}
		{{ end }}`

	tpl := template.Must(template.New("ds").Parse(s))
	tpl.Execute(os.Stdout, map[string][]map[string]string{
		"Vals": []map[string]string{
			{"Name": "Shine"},
		},
	})

	t.Error("err")
}

func TestFindRangeBlocks(t *testing.T) {
	s := `package searcher_templates

	type Query struct {
		Count int
		//# {{ range $key, $value := .Searchers }}
		_TemplateKey _TemplateKeyType
		//# {{ end }}
	}
	`
	res := findRangeBlocks(s)

	for _, e := range res {
		fmt.Println(s[e[0]:e[1]])
	}
	fmt.Printf("%#v\n", res)
	t.Error("err")
}

func TestRg(t *testing.T) {
	s := `package searcher_templates

	type Query struct {
		Count int
		//# {{ range $key, $value := .Searchers }}
		_TemplateKey _TemplateKeyType
		//# {{ end }}
	}
	`
	r, err := regexp.Compile(`_Template[[:alnum:]]*[\s\n]`)
	if err != nil {
		t.Error(err)
	}
	res := r.FindAllIndex([]byte(s), -1)
	for _, e := range res {
		fmt.Println(s[e[0]:e[1]])
	}
	fmt.Printf("%#v\n", res)
	t.Error("err")
}
