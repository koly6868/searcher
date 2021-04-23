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
	defaultTemplatesDir := "templates"

	loadTemplate := flag.Bool("l", true, "if load templates from github")
	templatesDir := flag.String("tmpls", defaultTemplatesDir, "template direcory")
	genPackageDir := flag.String("dst", "searcher", "gen destenation directory")
	configPath := flag.String("cfg", "cfg.json", "config path")
	help := flag.String("help", "help", "help")
	flag.Parse()

	if *help == "" {
		flag.Usage()
		return
	}

	err := os.MkdirAll(*genPackageDir, os.ModePerm)
	onError(err)

	if *loadTemplate && defaultTemplatesDir == *templatesDir {
		err := GitClone("github.com/koly6868/searcher", *templatesDir)
		defer os.RemoveAll(*templatesDir)
		*templatesDir = *templatesDir + "/_TemplatePackageName"
		onError(err)
	}

	cfg, err := readConfig(*configPath)
	onError(err)

	err = fillTemplates(*templatesDir, *genPackageDir, cfg)
	onError(err)
}

func fillTemplates(templatesDir string, targetDir string, cfg interface{}) error {
	ignoreFiles := []string{"model.go"}
	log.Info("generating starts")

	fileInfos, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return err
	}
	for _, info := range fileInfos {
		if !strings.HasSuffix(info.Name(), ".go") || strInArr(info.Name(), ignoreFiles) {
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
			log.Error(err)
			data = []byte(dataStr)
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

func strInArr(e string, arr []string) bool {
	for _, ce := range arr {
		if e == ce {
			return true
		}
	}
	return false
}
