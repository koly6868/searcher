package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/koly6868/searcher"
)

func main() {
	fp := flag.String("cfg", "cmd/example.json", "gen --cfg example.json")
	basePath := ""
	flag.Parse()
	f, err := os.Open(*fp)
	closeProgramIfErr(err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	closeProgramIfErr(err)

	gd := &searcher.GenData{}
	err = json.Unmarshal(data, gd)
	closeProgramIfErr(err)

	err = searcher.GenQuery(gd, basePath)
	closeProgramIfErr(err)

	err = searcher.GenerateCrativeSearchers(gd, basePath)
	closeProgramIfErr(err)

	err = searcher.GenCreativeSearcher(gd, basePath)
	closeProgramIfErr(err)
}

func closeProgramIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
