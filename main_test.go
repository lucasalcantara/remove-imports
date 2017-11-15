package main

import (
	"log"
	"os"
	util "test_util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var filesPath = "test_util/files/"

func setup() {
	initParallelism()

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

func TestWithoutRemoveImports(t *testing.T) {
	execute("without_remove_imports.go", util.FileAllImportsUsing, util.FileAllImportsUsing, t)
}

func TestRemovingImports(t *testing.T) {
	execute("remove_import_necessary.go", util.FileRemoveImport, util.FileRemoveImportExpected, t)
}

func TestRemovingImportsWithCommentsInCode(t *testing.T) {
	execute("with_comment.go", util.FileRemoveImportWithComments, util.FileRemoveImportWithCommentsExpected, t)
}

func execute(fileName, codeToFile, ExpectedCode string, t *testing.T) {
	path := filesPath + fileName
	util.CreateFile(codeToFile, path)

	analyzeDirectory(filesPath)

	wg.Wait()

	result := util.ReadFile(path)
	assert.Equal(t, result, ExpectedCode)
}
