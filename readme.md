# Element searcher

## How to use

1. Build generator ```go install github.com/koly6868/searcher```

2. Create config (cfg.json)

3. Run ```searcher```


## Description

This library allows to generate code for search data structure. Also there is an ability to store data
in desired oreder. You must pass functions for sorting and searching while initializing searcher.

Upper-bound acessments:

- O(k*n) memory for storing data;
- O(k*log(n)) time for retriving data;

n - count elements
k - count search parametrs.


## Example

Suppouse we have struct in package example:

```
type Person struct {
	ID   int
	Name string
	Age  int
}
```

To generate searcher the struct do following steps:

1. Instal search generator:
```
go install github.com/koly6868/searcher
```

2. Create ```cfg.json``` :
```
{
    "searchers" : [
        {
            "Name":      "NameSearcher",
            "KeyType":   "string",
            "Key":       "Name"
        },
        {
            "Name":      "AgeSearcher",
            "KeyType":   "int",
            "Key":       "Age"
        }
    ],
    "ModelName": "example.Person",
    "PackageName" : "searcher"
}
```

- PackageName - name of package;
- ModelName - name of struct;
- searchers.Name - name of searcher;
- searchers.Key - name of field in struct to search 
- searchers.KeyType - type of field in struct to search 


3. If you are in desired place just run:
```
searcher
```

If you want to specify destination dir, config path or teplate path please call
``search --help``` for details.

4. Create main file and define following functions:

```
func sortFn(data []*example.Person) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
}

func searchFN(data []*example.Person, e *example.Person) bool {
	index := sort.Search(
		len(data),
		func(index int) bool {
			return e.ID <= data[index].ID
		})

	return (index < len(data)) && (data[index].ID == e.ID)
}
```

- sortFn - function to sort data;
- searchFN - function to search data;

The must be corresponding. 

5. Create main function, initialize some test data and searcher, make some requests to searcher:

```
func main() {
	data := []example.Person{
		{ID: 1, Name: "One", Age: 32},
		{ID: 2, Name: "cos", Age: 22},
		{ID: 3, Name: "Lo", Age: 12},
		{ID: 4, Name: "Si", Age: 41},
		{ID: 5, Name: "Ni", Age: 100},
	}
	s := searcher.NewSearher(
		data,
		[]searcher.SearhModule{
			&searcher.AgeSearcher{},
			&searcher.NameSearcher{},
		},
		sortFn,
		searchFN,
	)

	res, err := s.Find(&searcher.Query{
		Count: 1,
		Name:  "Lo",
		Age:   12,
	})

	//prints res : []example.Person{example.Person{ID:3, Name:"Lo", Age:12}}; err : %!s(<nil>)
	fmt.Printf("res : %#v; err : %s\n", res, err)

	res, err = s.Find(&searcher.Query{
		Count: 1,
		Name:  "L",
		Age:   12,
	})
	// prints res : []example.Person{}; err : %!s(<nil>)
	fmt.Printf("res : %#v; err : %s\n", res, err)

	res, err = s.Find(&searcher.Query{
		Count: 1,
		Name:  "Lo",
		Age:   122,
	})
	// prints res : []example.Person{}; err : %!s(<nil>)
	fmt.Printf("res : %#v; err : %s\n", res, err)
}
```
