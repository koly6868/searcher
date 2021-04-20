package searcher

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"

	"golang.org/x/tools/imports"
)

const queryAnyValue = -1
const queryTemplate = `
type Query struct {
	Count int
	{{.Fields}}
}`
const queryPath = "generated_query.go"

func genQuery(gd *GenData) ([]byte, error) {
	genFields := bytes.Buffer{}
	for _, e := range gd.Searchers {
		genFields.WriteString(fmt.Sprintf("%s %s\n", e.Key, e.KeyType))
	}

	t := template.Must(template.New("q").Parse(queryTemplate))
	b := &bytes.Buffer{}
	err := t.Execute(b, map[string]string{"Fields": genFields.String()})

	return b.Bytes(), err
}

func GenQuery(gd *GenData, basePath string) error {
	f, err := os.Create(path.Join(basePath, queryPath))
	if err != nil {
		return err
	}

	data, err := genQuery(gd)
	if err != nil {
		return err
	}

	_, err = f.WriteString(gencCodeHeader)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	f.Close()
	imports.Process(path.Join(basePath, queryPath), nil, nil)
	return err
}
