package searcher

{{ range $key, $value := . }}
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
    }
{{ end }}