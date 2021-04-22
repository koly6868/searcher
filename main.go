package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git"
	log "github.com/sirupsen/logrus"
	"golang.org/x/tools/imports"
)

func main() {
	loadTemplate := flag.Bool("l", true, "hz l")
	templatesDir := flag.String("tmpls", "templates", "hz tpl")
	genPackageDir := flag.String("dst", "searcher", "hz")
	configPath := flag.String("cfg", "cfg.json", "hz cfg")

	err := os.MkdirAll(*genPackageDir, os.ModePerm)
	onError(err)

	if *loadTemplate {
		err := GitClone("github.com/koly6868/searcher", *templatesDir)
		onError(err)
	}
	defer os.RemoveAll(*templatesDir)

	cfg, err := readConfig(*configPath)
	onError(err)

	err = fillTemplates(*templatesDir, *genPackageDir, cfg)
	onError(err)
}

func fillTemplates(templatesDir string, targetDir string, cfg interface{}) error {
	log.Info("generating starts")

	fileInfos, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return err
	}
	for _, info := range fileInfos {
		if !strings.HasSuffix(info.Name(), ".go") {
			continue
		}
		log.Infof("%s is processing", info.Name())
		filePath := path.Join(templatesDir, info.Name())
		saveFilePath := path.Join(targetDir, info.Name())
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}

		dataStr := string(data)
		dataStr = preprocessCodeTemplate(dataStr)
		dataStr = fillCodeTemplate(dataStr, cfg)
		data, err = imports.Process(saveFilePath, []byte(dataStr), nil)
		if err != nil {
			return err
		}
		ioutil.WriteFile(saveFilePath, data, os.ModePerm)
		log.Infof("%s has been processed", info.Name())
	}

	return nil
}

func onError(err error) {
	if err != nil {
		panic("top err " + err.Error())
	}
}

func readConfig(path string) (interface{}, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &SearcherConfig{}
	err = json.Unmarshal(data, cfg)
	return cfg, err
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
