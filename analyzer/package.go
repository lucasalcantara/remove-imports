package analyzer

import (
	"file"
	"os"
	"path"
	"strings"
)

const (
	srcPrefix = "src/"
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

	if _, found := packages[importPath]; found {
	} else if _, err := os.Stat(goroot + importPath); err == nil {

		packages[importPath] = importPath
	} else if _, err := os.Stat(gopath + importPath); err == nil {

		packages[importPath] = findPackageNameInDirectory(gopath + importPath)
	}

	return packages[importPath]
}

func findPackageNameInDirectory(p string) string {
	packageName := ""
	files := file.Files(p)

	for _, file := range files {
		fullPath := path.Join(p, file.Name())
		isGoFile := strings.Contains(path.Base(fullPath), ".go")
		isGoFileTest := strings.Contains(path.Base(fullPath), "_test.go")

		if isGoFile && !isGoFileTest {
			packageName = findPackageNameInFile(fullPath)
			break
		}
	}

	return packageName
}

func findPackageNameInFile(p string) string {
	packageName := ""
	comment := &ScopeComments{}
	f, scanner := file.Scanner(p)
	defer f.Close()

	for scanner.Scan() {
		text := scanner.Text()

		if !comment.InCommentLine(text) && text != "" {
			values := strings.Split(text, " ")
			packageName = values[1]

			break
		}
	}

	return packageName
}
