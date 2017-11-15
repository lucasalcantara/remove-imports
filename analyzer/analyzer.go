package analyzer

import "strings"

type AnalyzerText interface {
	Filter(text string) bool
	Analyze(text string) string
}

type Analyzer struct {
	analyzeLines bool
	analyzers    []AnalyzerText
}

func InitAnalyzer() *Analyzer {
	return &Analyzer{
		analyzeLines: false,
		analyzers:    newAnalyzes(),
	}
}

func newAnalyzes() []AnalyzerText {
	return []AnalyzerText{
		AnalyzerImportLine{},
		&AnalyzerImportBlock{},
	}
}

func (a *Analyzer) AnalyzeText(text string) string {
	if a.isImportScope(text) || a.analyzeLines {
		a.analyzeLines = true

		analyzerText := a.getAnalyzerText(text)
		if analyzerText != nil {
			return analyzerText.Analyze(text)
		}
	}

	return ""
}

func (a *Analyzer) isImportScope(text string) bool {
	text = strings.Trim(text, " ")
	values := strings.Split(text, " ")

	return values[0] == "import"
}

func (a *Analyzer) getAnalyzerText(text string) (analyzer AnalyzerText) {
	for _, a := range a.analyzers {
		if a.Filter(text) {
			analyzer = a
			break
		}
	}

	return
}
