package main

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

const templateVarHeaer = "_Template"

var (
	findRangeBlocksExp      *regexp.Regexp
	extractRangeKeyValueExp *regexp.Regexp
	// TODO delete
	findRangeVars     *regexp.Regexp
	findBodyRangeVars *regexp.Regexp
)

func init() {
	findRangeBlocksExp = regexp.MustCompile(`(?s)//#\s*{{\s*range\s*.*?}}.*//#\s*{{\s*end\s*}}`)
	extractRangeKeyValueExp = regexp.MustCompile(`{{\s*range\s*.*?}}`)
	findRangeVars = regexp.MustCompile(`\$[[:alnum:]]*[\s,]`)
	findBodyRangeVars = regexp.MustCompile(templateVarHeaer + `[[:alnum:]]*[\s,]`)

}

type SearcherConfig struct {
	Searchers []SeacrherModuleConfig `json:"searchers"`
	ModelName string                 `json:"ModelName"`
}

type SeacrherModuleConfig struct {
	Name    string `json:"Name"`
	KeyType string `json:"KeyType"`
	Key     string `json:"Key"`
}

func fillCodeTemplate(tmpl string, cfg *SearcherConfig) string {
	tpl := template.Must(template.New("ds").Parse(tmpl))
	b := &bytes.Buffer{}
	tpl.Execute(b, cfg)

	return b.String()
}

func preprocessCodeTemplate(tmpl string) string {
	tmpl = preprocessRangeStatements(tmpl)
	tmpl = strings.ReplaceAll(tmpl, "//#", "")
	return tmpl
}

func preprocessRangeStatements(tmpl string) string {
	return prerpocessRangeBodyTemplate(tmpl)
}

func findRangeBlocks(s string) [][]int {
	return findRangeBlocksExp.FindAllIndex([]byte(s), -1)
}

func extractRangeKeyValue(s string) [][]int {
	rangeBlockInd := extractRangeKeyValueExp.FindAllIndex([]byte(s), -1)[0]
	rangeBlock := s[rangeBlockInd[0]:rangeBlockInd[1]]
	varInds := findRangeVars.FindAllIndex([]byte(rangeBlock), -1)

	for _, e := range varInds {
		// shif from start
		e[0] += rangeBlockInd[0]
		// do not include last space or comma
		e[1] += rangeBlockInd[0] - 1
	}
	return varInds
}

func prerpocessRangeBodyTemplate(s string) string {
	vars := findBodyRangeVars.FindAll([]byte(s), -1)
	sort.Slice(vars, func(i, j int) bool {
		return len(vars[j]) < len(vars[i])
	})
	log.Infof("found %d range variabels", len(vars))
	for _, e := range vars {
		v := string(e[:len(e)-1])
		nv := fmt.Sprintf("{{ .%s }}", v[len(templateVarHeaer):])
		s = strings.ReplaceAll(s, v, nv)
	}
	return s
}
