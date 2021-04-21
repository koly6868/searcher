package searcher

import (
	"bytes"
	"fmt"
	"path"
	"text/template"
)

var moduleTemplate = `
type {{.Name}} struct {
	data        map[{{.KeyType}}][]*{{.ModelName}}
	sortSliceFn func([]*{{.ModelName}})
	searchFn func([]*example.TestModel, *example.TestModel) int
}

func (sm *{{.Name}}) init(data []{{.ModelName}},
	 sortSliceFn func([]*{{.ModelName}}),
	 searchFn func([]*example.TestModel, *example.TestModel) int) {
	
	sm.data = map[{{.KeyType}}][]*{{.ModelName}}{}
	sm.sortSliceFn = sortSliceFn
	sm.searchFn = searchFn

	for i, e := range data {
		sm.data[e.{{.Key}}] = append(sm.data[e.{{.Key}}], &data[i])
	}

	for k := range sm.data {
		sortSliceFn(sm.data[k])
	}
}

func (sm *{{.Name}}) find(q *Query) searchResult {
	return newSimpleResult(
		sm.data[q.{{.Key}}],
		sm.searchFn)
}`

const serachModulesPath = "generated_search_modules.go"

func GenerateSearchers(gd *GenData, basePath string) error {
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
	err := FormatAndWrite(path.Join(basePath, serachModulesPath), codeBuff.Bytes())

	return err
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
