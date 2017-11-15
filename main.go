package main

import (
	"analyzer"
	"bufio"
	"file"
	"flag"
	"fmt"
	"log"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	regex *regexp.Regexp
	ch    chan string
	wg    sync.WaitGroup
)

const parallelism = 4

func init() {
	ch = make(chan string)
	regex = regexp.MustCompile("\"(.*?)\"")
}

func main() {

	initParallelism()

	var path = flag.String("path", "", "path to remove the unused imports")
	var gopath = flag.String("gopath", "", "path of the go project")

	flag.Parse()
	analyzer.SetGopath(*gopath)

	start := time.Now()

	if *path != "" && *gopath != "" {
		log.Printf("Running to path %s", *path)
		analyzeDirectory(*path)
	} else {
		log.Printf("Path or gopath was not provided.")
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Time: %s", elapsed)
}

func initParallelism() {
	for i := 0; i < parallelism; i++ {
		go parallelRemoveImports()
	}
}

func analyzeDirectory(p string) {
	files := file.Files(p)
	wg.Add(1)

	for _, file := range files {
		fullPath := path.Join(p, file.Name())

		if file.IsDir() {
			analyzeDirectory(fullPath)
		} else if isGoFile := strings.Contains(path.Base(fullPath), ".go"); isGoFile {
			wg.Add(1)
			ch <- fullPath
		}
	}

	wg.Done()
}

func parallelRemoveImports() {
	for {
		select {
		case p := <-ch:
			removeImportsNotUsed(p)
			wg.Done()
		}
	}
}

func removeImportsNotUsed(path string) {
	imports := importsNotUsed(path)

	if len(imports) > 0 {
		f, scanner := file.Scanner(path)
		defer f.Close()

		output := newFileOutput(scanner, imports)
		file.Override(output, f.Name())
	}
}

func importsNotUsed(path string) map[string]bool {
	imports := getImports(path)

	if len(imports) > 0 {
		f, scanner := file.Scanner(path)
		defer f.Close()

		for scanner.Scan() {
			importsFound := importsFoundInLine(scanner.Text(), imports)

			for _, importFound := range importsFound {
				delete(imports, importFound)
			}

			if len(imports) == 0 {
				break
			}
		}
	}

	return imports
}

func getImports(path string) map[string]bool {
	f, scanner := file.Scanner(path)
	analyzes := analyzer.InitAnalyzer()
	imports := make(map[string]bool, 0)
	defer f.Close()

	for scanner.Scan() {
		text := scanner.Text()

		passedImportDeclaration := strings.Contains(text, "func") || strings.Contains(text, "const") ||
			strings.Contains(text, "type") || strings.Contains(text, "var")

		if passedImportDeclaration {
			break
		}

		importName := analyzes.AnalyzeText(text)
		if importName != "" {
			imports[importName] = true
		}
	}

	return imports
}

func importsFoundInLine(line string, imports map[string]bool) []string {
	importsFound := make([]string, 0)

	for importName, _ := range imports {
		libInUse := importName + "."
		values := strings.Split(line, libInUse)

		if len(values) > 1 {
			importsFound = append(importsFound, importName)
		}
	}

	return importsFound
}

func newFileOutput(scanner *bufio.Scanner, imports map[string]bool) string {
	commentScope := &analyzer.ScopeComments{}
	output := ""

	for scanner.Scan() {
		line := scanner.Text()
		lineWithoutSpaces := strings.Replace(line, " ", "", -1)

		if lineWithoutSpaces == "" {
			output += "\n"
		} else if commentScope.InCommentLine(line) {
			output += line + "\n"
		} else {
			output += removeImportUnusedOnTheLine(line, imports)
		}
	}

	return output
}

func removeImportUnusedOnTheLine(line string, imports map[string]bool) string {
	result := ""

	for lib, _ := range imports {
		outputAux := removeLibOnTheLine(line, lib)

		if outputAux == line {
			result = line + "\n"
		} else {
			result = ""
			break
		}
	}

	return result
}

func removeLibOnTheLine(line, lib string) string {
	aux := line

	formatImport := fmt.Sprintf("import \"%s\"", lib)
	formatImport2 := fmt.Sprintf("\"%s\"", lib)

	aux = strings.Replace(aux, formatImport, "", -1)
	aux = strings.Replace(aux, formatImport2, "", -1)

	return aux
}
