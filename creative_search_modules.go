package searcher

import (
	"bytes"
	"fmt"
	"path"
	"text/template"
	"unicode"
)

var moduleTemplate = `
type {{.NameLower}} struct {
	data        map[{{.KeyType}}][]*{{.ModelName}}
	sortSliceFn func([]*{{.ModelName}})
}

func New{{.Name}}(sortSliceFn func([]*{{.ModelName}})) searhModule {
	return &{{.NameLower}}{
		sortSliceFn: sortSliceFn,
		data:        map[string][]*{{.ModelName}}{},
	}
}

func (sm {{.NameLower}}) init(data []{{.ModelName}}) {
	for i, e := range data {
		sm.data[e.{{.Key}}] = append(sm.data[e.{{.Key}}], &data[i])
	}

	for k := range sm.data {
		sm.sortSliceFn(sm.data[k])
	}
}

func (sm {{.NameLower}}) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q.{{.Key}}])
}`

const serachModulesPath = "generated_search_modules.go"

func GenerateCrativeSearchers(gd *GenData, basePath string) error {
	codeBuff := &bytes.Buffer{}
	codeBuff.WriteString(gencCodeHeader)

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
		codeBuff.Write(s)
	}
	err := formatAndWrite(path.Join(basePath, serachModulesPath), codeBuff.Bytes())

	return err
}

type SearchModuleGenData struct {
	Name      string
	KeyType   string
	Key       string
	ModelName string
}

func (smgd *SearchModuleGenData) NameLower() string {
	a := []rune(smgd.Name)
	a[0] = unicode.ToLower(a[0])
	return string(a)
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
