package searcher

import (
	"sort"
	"testing"
)

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

func defaultSortFn(data []*TestModel) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})
}
