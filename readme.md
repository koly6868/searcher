# Creative searcher

## How to use

1. change ```searcher.json```
2. run ```go generate```


## searcher.json description

- ModelName: name of struct with package (e.g. ```example.TestModell```)
- searchers: list of searchers with ```Name``` - searcher name, ```Key``` - public field of struct and
 ```KeyType``` - type of field


## Example

1. Define sort function 

```
func defaultSortFn(data []*TestModel) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
}
```

2. Define  searcher, serach modules and test

```
func TestSearcher(t *testing.T) {
	targetInd := 3
	data := []TestModel{
		{ID: 1, Name: "ass"},
		{ID: 2, Name: "asss"},
		{ID: 3, Name: "assss"},
		{ID: 4, Name: "awss"},
		{ID: 5, Name: "assw"},
	}
	s := NewCreativeSearh(data, []searhModule{
		NewNameSearchModule(defaultSortFn),
	})

	res, err := s.Find(&Query{
		Count: 1,
		Name:  "awss",
	})
	if err != nil {
		t.Error(err)
	}
	if len(res) != 1 || res[0] != data[targetInd] {
		t.Errorf("got: %#v;\nwant: %#v", res, data[targetInd])
	}
}
```