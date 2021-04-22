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
	findRangeBlocksExp    *regexp.Regexp
	extractRangeHeaderExp *regexp.Regexp
	// TODO delete
	findBodyRangeVars *regexp.Regexp
	findWhiteSpace    *regexp.Regexp
)

func init() {
	findRangeBlocksExp = regexp.MustCompile(`(?s)//#\s*{{\s*range\s*.*?}}.*//#\s*{{\s*end\s*}}`)
	extractRangeHeaderExp = regexp.MustCompile(`{{\s*range\s*.*?}}`)
	findBodyRangeVars = regexp.MustCompile(templateVarHeaer + `[[:alnum:]]*`)
	findWhiteSpace = regexp.MustCompile(`\s`)

}

func fillCodeTemplate(tmpl string, cfg interface{}) string {
	tpl := template.Must(template.New("ds").Parse(tmpl))
	b := &bytes.Buffer{}
	tpl.Execute(b, cfg)

	return b.String()
}

func preprocessCodeTemplate(tmpl string) string {
	tmpl = preprocessRangeStatements(tmpl)
	tmpl = strings.ReplaceAll(tmpl, "//#", "")
	tmpl = prerpocessTemplateVars(tmpl, "")
	return tmpl
}

func preprocessRangeStatements(tmpl string) string {
	inds := findRangeBlocks(tmpl)
	if len(inds) == 0 {
		return tmpl
	}
	log.Infof("found %d range blocks", len(inds))

	res := bytes.Buffer{}
	last := 0
	for _, rangeBlockInds := range inds {
		res.WriteString(tmpl[last:rangeBlockInds[0]])

		rangeBlock := tmpl[rangeBlockInds[0]:rangeBlockInds[1]]
		skipHeader := extractRangeIterVariable(rangeBlock)
		skipHeader = normalizeRangeIterVariable(skipHeader)
		rangeBlock = prerpocessTemplateVars(rangeBlock, skipHeader)

		res.WriteString(rangeBlock)
		last = rangeBlockInds[1]
	}
	res.WriteString(tmpl[last:])
	return res.String()
}

func findRangeBlocks(s string) [][]int {
	return findRangeBlocksExp.FindAllIndex([]byte(s), -1)
}

func prerpocessTemplateVars(s, header string) string {
	vars := findBodyRangeVars.FindAll([]byte(s), -1)
	sort.Slice(vars, func(i, j int) bool {
		return len(vars[j]) < len(vars[i])
	})
	log.Infof("found %d variabels processing with header '%s'", len(vars), header)
	for _, e := range vars {
		v := string(e)

		// skip template prefix
		nv := v[len(templateVarHeaer):]
		if len(header) > 0 && strings.HasPrefix(nv, header) {
			nv = nv[len(header):]
			nv = fmt.Sprintf("{{ .%s }}", nv)
		} else {
			nv = fmt.Sprintf("{{ $.%s }}", nv)
		}

		s = strings.ReplaceAll(s, v, nv)
	}
	return s
}

func extractRangeIterVariable(s string) string {
	header := string(extractRangeHeaderExp.Find([]byte(s)))
	header = header[2 : len(header)-2]
	header = strings.Trim(header, ` \t\n\r`)
	header = string(findWhiteSpace.ReplaceAll([]byte(header), []byte(" ")))
	parts := strings.Split(header, " ")
	header = parts[len(parts)-1]
	if (len(header) > 0) && (header[0] == '.') {
		header = header[1:]
	}

	return header
}

func normalizeRangeIterVariable(s string) string {
	n := len(s)
	if n > 0 && s[n-1:] == "s" {
		s = s[:n-1]
	}

	return s
}
