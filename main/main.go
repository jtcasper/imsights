package main

import (
	"fmt"
	_ "github.com/jtcasper/imsights/types"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func visit(path string, f os.FileInfo, e error) (err error) {
	if f.IsDir() {
		return
	}

	file, err := os.Open(path)

	if err != nil {
		return
	}

	defer file.Close()

	body, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(body))

	return
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Root argument required. Usage: main <root>")
	}
	root := os.Args[1]
	filepath.Walk(root, visit)
}
