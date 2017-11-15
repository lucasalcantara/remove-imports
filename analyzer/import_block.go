package analyzer

import "strings"

type AnalyzerImportBlock struct {
	blockOpened bool
	ScopeComments
}

func (s *AnalyzerImportBlock) Filter(text string) bool {
	s.blockOpened = (strings.Contains(text, "import") && strings.Contains(text, "(")) || s.blockOpened
	return s.blockOpened
}

func (s *AnalyzerImportBlock) Analyze(text string) string {
	result := ""

	if text == ")" {
		s.blockOpened = false
		result = ""
	} else if !strings.Contains(text, "import") {
		result = importName(1, text)
	}

	return result
}
