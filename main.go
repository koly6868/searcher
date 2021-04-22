package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/tools/imports"
)

func main() {
	templatesDir := flag.String("tmpls", "templates", "hz tpl")
	genPackageDir := flag.String("dst", "", "hz")
	configPath := flag.String("cfg", "cfg.json", "hz cfg")
	err := fillTemplates(templatesDir, *genPackageDir, cfg)
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

func readConfig(path string) interface{} {

}
