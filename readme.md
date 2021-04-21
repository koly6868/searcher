# Element searcher

## How to use

1. change ```searcher.json```
2. run ```go generate```


## searcher.json description

- ModelName: name of struct with package (e.g. ```example.TestModell```)
- searchers: list of searchers with ```Name``` - searcher name, ```Key``` - public field of struct and
 ```KeyType``` - type of field


## Example

1. change file cmd/searcher.json in following way
```
{
    "searchers" : [
        {
            "Name":      "Name",
            "KeyType":   "string",
            "Key":       "Name"
        },
        {
            "Name":      "Age",
            "KeyType":   "int",
            "Key":       "Age"
        }
    ],
    "ModelName": "example.TestModel"
}
```

2. run ```go generate```

3.  create project and import packages

```
	"github.com/koly6868/searcher"
	"github.com/koly6868/searcher/example"
```

4. define sort and coressponding serach functions for the model
```
func defaultSortFn(data []*example.TestModel) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
}

func defaultSerachFn(data []*example.TestModel, e *example.TestModel) int {
	return sort.Search(
		len(data),
		func(index int) bool {
			return e.ID <= data[index].ID
		})
}
```

5. create sample data and searcher. Call Find method of searcher on diffirent queries
```
func main() {
	data := []example.TestModel{
		{ID: 1, Name: "ass", Age: 32},
		{ID: 2, Name: "asss", Age: 3},
		{ID: 3, Name: "assss", Age: 23},
		{ID: 4, Name: "awss", Age: 12},
		{ID: 5, Name: "assw", Age: 45},
	}
	s := searcher.NewSearher(data, []searcher.SearhModule{
		&searcher.NameSearchModule{},
		&searcher.AgeSearchModule{},
	}, defaultSortFn, defaultSerachFn)

	// prints res :[]example.TestModel{example.TestModel{ID:4, Name:"awss", Age:12}}; err: %!s(<nil>)
	res, err := s.Find(&searcher.Query{
		Count: 1,
		Name:  "awss",
		Age:   12,
	})
	fmt.Printf("res :%#v; err: %s \n", res, err)

	// prints res :[]example.TestModel{}; err: %!s(<nil>)
	res, err = s.Find(&searcher.Query{
		Count: 1,
		Name:  "s",
		Age:   12,
	})
	fmt.Printf("res :%#v; err: %s \n", res, err)

	// res :[]example.TestModel{}; err: %!s(<nil>)
	res, err = s.Find(&searcher.Query{
		Count: 0,
		Name:  "s",
		Age:   12,
	})
	fmt.Printf("res :%#v; err: %s \n", res, err)
}
```