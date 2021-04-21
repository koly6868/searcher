package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"

	"github.com/go-git/go-git"
	"github.com/koly6868/searcher"
)

const tempDir = "temp"

func main() {

	configFPath := flag.String("cfg", "cmd/example.json", "gen --cfg example.json")
	dstFPath := flag.String("d", "cmd/example.json", "gen --d .")
	flag.Parse()

	err := GitClone("github.com/koly6868/searcher", tempDir)
	closeProgramIfErr(err)

	f, err := os.Open(*configFPath)
	closeProgramIfErr(err)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	closeProgramIfErr(err)

	gd := &searcher.GenData{}
	err = json.Unmarshal(data, gd)
	closeProgramIfErr(err)

	err = searcher.GenQuery(gd, basePath)
	closeProgramIfErr(err)

	err = searcher.GenerateSearchers(gd, basePath)
	closeProgramIfErr(err)

	err = searcher.GenSearcher(gd, basePath)
	closeProgramIfErr(err)

	err = searcher.GenElementSerachResult(gd, basePath)
	closeProgramIfErr(err)

}

func GitClone(repoPath, directory string) error {
	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL: "https://" + repoPath,
	})
	if err != nil {
		return err
	}
	return nil
}

func closeProgramIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
