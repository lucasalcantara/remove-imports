package analyzer

import (
	"file"
	"os"
	"strings"
)

const (
	srcPrefix = "/src/"
)

var (
	gopath   string
	packages map[string]string
	goroot   = os.Getenv("GOROOT") + srcPrefix
)

func init() {
	packages = make(map[string]string, 0)
}

func SetGopath(gopath string) {
	gopath = gopath + srcPrefix
}

func findPackageName(importPath string) string {
	var packageName string

	if packageName, found := packages[importPath]; found {
	} else if _, err := os.Stat(goroot + importPath); err == nil {

		packageName = importPath
		packages[importPath] = packageName
	} else if _, err := os.Stat(gopath + importPath); err == nil {

		packageName = findPackageNameInFile(gopath + importPath)
		packages[importPath] = packageName
	}

	return packageName
}

func findPackageNameInFile(path string) string {
	packageName := ""
	comment := new(ScopeComments)
	f, scanner := file.Scanner(path)
	defer f.Close()

	for scanner.Scan() {
		text := scanner.Text()
		if !comment.InCommentLine(text) {
			values := strings.Split(text, " ")
			packageName = values[1]

			break
		}
	}

	return packageName
}
