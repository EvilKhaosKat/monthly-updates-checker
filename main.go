package main

import (
	"io/ioutil"
	"strings"
	"os"
	"log"
	"sync"
	"strconv"
	"fmt"
	"sort"
)

const Dir = "files"

type Result struct {
	year    int
	month   int

	value   int

	updated bool
	delta   int
}

func (r Result) String() string {
	return fmt.Sprintf("(year:%v, month:%v, value:%v, updated:%v, delta:%v)",
		r.year,
		r.month,
		r.value,
		r.updated,
		r.delta,
	)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	files := getSuitableFiles()
	filesCount := len(files)

	resultsChan := make(chan *Result, filesCount)

	wg := &sync.WaitGroup{}
	wg.Add(filesCount)

	for _, f := range files {
		go handleFile(f.Name(), resultsChan, wg)
	}

	wg.Wait()

	analyzeResults(resultsChan)
}

func analyzeResults(resultsChan chan *Result) {
	results := getResultsSlice(resultsChan)
	sort.Sort(ByDate(results))

	fmt.Println(results)
}

func getResultsSlice(resultsChan chan *Result) []*Result {
	resultsLen := len(resultsChan)

	results := make([]*Result, resultsLen)
	for i := 0; i < resultsLen; i++ {
		results[i] = <-resultsChan
	}

	return results
}

func getSuitableFiles() []os.FileInfo {
	var suitableFiles []os.FileInfo

	files := getFiles()
	for _, f := range files {
		if isSuitableFile(f) {
			suitableFiles = append(suitableFiles, f)
		}
	}

	return suitableFiles
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

func handleFile(filename string, resultsChan chan *Result, wg *sync.WaitGroup) {
	number := getMagicNumber(filename)
	year, month := getDate(filename)

	resultsChan <- &Result{year:year, month:month, value:number}

	wg.Done()
}

func getDate(filename string) (int, int) {
	splittedValues := strings.Split(filename, ".")

	year, err := strconv.Atoi(splittedValues[0])
	if err != nil {
		panic(fmt.Sprintf("Error during getDate year for %v:%v", filename, err))
	}

	month, err := strconv.Atoi(splittedValues[1])
	if err != nil {
		panic(fmt.Sprintf("Error during getDate month for %v:%v", filename, err))
	}

	return year, month
}
