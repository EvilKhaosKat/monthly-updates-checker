package main

import (
	_ "github.com/tealeg/xlsx"
	"io/ioutil"
	"strings"
	"os"
	"log"
)

const Dir = "files"

type result struct {
	year    int
	month   int

	value   int

	updated bool
	delta   int
}

func main() {
	resultsChan := make(chan result)

	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if isSuitableFile(f) {
			go handleFile(f.Name(), resultsChan)
		}
	}

	//var results []result
}
func isSuitableFile(file os.FileInfo) bool {
	return !file.IsDir() && strings.HasSuffix(file.Name(), ".xls")
}

func handleFile(filename string, resultsChan chan result) {
	log.Printf("handle file %s", filename)
}
