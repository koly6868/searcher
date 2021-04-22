package main

import (
	"io/ioutil"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

func main() {
	templatesDir := "searcher_templates"
	err := fillTemplates(templatesDir, "generated")
	onError(err)
}

func fillTemplates(templatesDir string, targetDir string) error {
	log.Info("generating starts")

	fileInfos, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return err
	}
	for _, info := range fileInfos {
		if !EndsWith(info.Name(), ".go") {
			continue
		}
		log.Infof("%s is processing", info.Name())
		filePath := path.Join(templatesDir, info.Name())
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}

		dataStr := string(data)
		dataStr = preprocessCodeTemplate(dataStr)
		ioutil.WriteFile(path.Join(targetDir, info.Name()), []byte(dataStr), os.ModePerm)
		log.Infof("%s has been processed", info.Name())
	}

	return nil
}

func onError(err error) {
	if err != nil {
		panic("top err " + err.Error())
	}
}
