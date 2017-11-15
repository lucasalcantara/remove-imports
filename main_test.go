package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	util "test_util"
	"testing"
)

var (
	filesPath  = "test_util/files/"
	fileToTest = []struct {
		name           string
		code           string
		expectedResult string
		errorMessage   string
	}{
		{"no_remove_import_necessary.go", util.FileAllImportsUsing, util.FileAllImportsUsing, "the code in file %s is not equals of the expected result."},
		{"remove_import_necessary.go", util.FileRemoveImport, util.FileRemoveImportExpected, "the code in file %s is not equals of the expected result."},
		{"import_necessary_with_comment.go", util.FileRemoveImportWithComments, util.FileRemoveImportWithCommentsExpected, "the code in file %s is not equals of the expected result."},
	}
)

func setup() {
	err := os.MkdirAll(filesPath, 0777)
	if err != nil {
		log.Panic(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	run := m.Run()
	os.RemoveAll(filesPath)

	os.Exit(run)
}

func TestImports(t *testing.T) {
	for _, file := range fileToTest {
		path := filesPath + file.name
		util.CreateFile(file.code, path)
	}

	initParallelism()
	analyzeDirectory(filesPath)

	wg.Wait()

	for _, file := range fileToTest {
		path := filesPath + file.name
		result := util.ReadFile(path)

		assert.Equal(t, result, file.expectedResult, fmt.Sprintf(file.errorMessage, file.name))
	}
}
