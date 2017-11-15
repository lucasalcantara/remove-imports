package file

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func Files(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func Override(output, fileName string) {
	err := ioutil.WriteFile(fileName, []byte(output), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func Scanner(path string) (*os.File, *bufio.Scanner) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return file, scanner
}
