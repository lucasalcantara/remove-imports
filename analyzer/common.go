package analyzer

import "strings"

func importName(threshold int, text string) (importName string) {
	values := strings.Split(formatText(text), " ")
	hasAlias := len(values) > threshold

	if hasAlias {
		importName = values[len(values)-2]
	} else {
		values = strings.Split(values[len(values)-1], "/")
		importName = values[len(values)-1]
	}

	if importName == "_" {
		importName = ""
	}

	return
}

func formatText(text string) string {
	formattedText := strings.Trim(text, " ")
	formattedText = strings.Replace(formattedText, "\t", "", -1)
	formattedText = strings.Replace(formattedText, "\"", "", -1)

	return formattedText
}
