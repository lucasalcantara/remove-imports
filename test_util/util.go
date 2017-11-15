package test_util

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(code, path string) {
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			log.Panic(err)
		}

		defer file.Close()
	}

	err = ioutil.WriteFile(path, []byte(code), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		log.Panic(err)
	}

	return string(b)
}

var (
	FileAllImportsUsing = `
package main 

import (
    "log"
    "os"
    "path/filepath"
)

import "fmt"

func main() {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
    fmt.Println(dir)
}
`

	FileRemoveImport = `
package main 

import (
    "log"
    "os"
    "path/filepath"
)

import "fmt"

func main() {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
}
`

	FileRemoveImportExpected = `
package main 

import (
    "log"
    "os"
    "path/filepath"
)


func main() {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
}
`

	FileRemoveImportWithComments = `
	/* 
	import "fmt"
	import (
	    "teste"
	    "os"
	    "path/filepath"
	)
	*/
package main 

import (
    "log"
    "os"
    "path/filepath"
)

import "fmt"

func main() {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
}
`

	FileRemoveImportWithCommentsExpected = `
	/* 
	import "fmt"
	import (
	    "teste"
	    "os"
	    "path/filepath"
	)
	*/
package main 

import (
    "log"
    "os"
    "path/filepath"
)


func main() {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
}
`
)
