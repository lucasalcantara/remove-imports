package analyzer

import "strings"

type AnalyzerImportLine struct {
	ScopeComments
}

func (AnalyzerImportLine) Filter(text string) bool {
	return strings.Contains(text, "import") && strings.Contains(text, "\"")
}

func (AnalyzerImportLine) Analyze(text string) string {
	return importName(2, text)
}
