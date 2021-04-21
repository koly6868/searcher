package searcher

import (
	"bytes"
	"fmt"
	"path"
	"text/template"
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
	codeBuff := &bytes.Buffer{}
	data, err := genQuery(gd)
	if err != nil {
		return err
	}

	codeBuff.WriteString(gencCodeHeader)
	codeBuff.Write(data)

	err = formatAndWrite(path.Join(basePath, queryPath), codeBuff.Bytes())

	return err
}
