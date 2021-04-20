package searcher

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"
)

var moduleTemplate = `
type {{.Name}} map[{{.KeyType}}][]*{{.ModelName}}

func (sm {{.Name}}) init(data []{{.ModelName}}, sortSliceFn func([]*{{.ModelName}})) {
	for i, e := range data {
		sm[e.{{.Key}}] = append(sm[e.{{.Key}}], &data[i])
	}

	for k := range sm {
		sortSliceFn(sm[k])
	}
}

func (sm {{.Name}}) find(q *Query) searchResult {
	return newSimpleResult(
		sm[q.{{.Key}}])
}`

const genCodePath = "generated_search_modules.go"

func GenerateCrativeSearchers(gd *GenData, basePath string) error {
	f, err := os.Create(path.Join(basePath, genCodePath))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(gencCodeHeader)
	if err != nil {
		return err
	}

	for _, e := range gd.Searchers {
		s, err := genModule(&SearchModuleGenData{
			Name:      e.Name,
			KeyType:   e.KeyType,
			Key:       e.Key,
			ModelName: gd.ModelName,
		})
		if err != nil {
			return err
		}
		_, err = f.Write(s)
		if err != nil {
			return err
		}
	}

	return nil
}

type SearchModuleGenData struct {
	Name      string
	KeyType   string
	Key       string
	ModelName string
}

func (smgd *SearchModuleGenData) normalize() error {
	errKey := ""
	switch {
	case smgd.KeyType == "":
		errKey = "KeyType"
	case smgd.Key == "":
		errKey = "Key"
	case smgd.ModelName == "":
		errKey = "ModelName"
	}
	if errKey != "" {
		return fmt.Errorf("%s is required", errKey)
	}

	if smgd.Name == "" {
		smgd.Name = smgd.Key
	}
	smgd.Name = smgd.Name + "SearchModule"
	return nil
}

func genModule(gd *SearchModuleGenData) ([]byte, error) {
	err := gd.normalize()
	if err != nil {
		return []byte{}, err
	}

	t := template.Must(template.New("sg").Parse(moduleTemplate))
	b := &bytes.Buffer{}
	err = t.Execute(b, gd)

	return b.Bytes(), err
}
