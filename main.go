package main

import (
	"io/ioutil"
	"strings"
	"os"
	"log"
	"sync"
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
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	files := getFiles()
	filesCount := len(files)

	resultsChan := make(chan result, filesCount)

	wg := &sync.WaitGroup{}

	for _, f := range files {
		if isSuitableFile(f) {
			wg.Add(1)
			/*go*/ handleFile(f.Name(), resultsChan, wg)
		}
	}

	wg.Wait()

	analyzeResults(resultsChan)
}

func analyzeResults(results chan result) {

}

func getFiles() []os.FileInfo {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func isSuitableFile(file os.FileInfo) bool {
	return !file.IsDir() && strings.HasSuffix(file.Name(), ".xls")
}

func handleFile(filename string, resultsChan chan result, wg *sync.WaitGroup) {
	//log.Printf("handle file %s", filename)

	number := getMagicNumber(filename)
	log.Println(number)

	wg.Done()
}
