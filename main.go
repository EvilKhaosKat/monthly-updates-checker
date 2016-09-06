package main

import (
	"io/ioutil"
	"strings"
	"os"
	"log"
	"sync"
	"github.com/extrame/xls"
	"fmt"
	"strconv"
	"regexp"
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
	workBook, err := xls.Open(fmt.Sprintf("%v/%v", Dir, filename), "UTF-8")
	if err != nil {
		log.Fatal(err)
	}

	number := getMagicNumber(workBook, filename)
	log.Println(number)

	wg.Done()
}
func getMagicNumber(workBook *xls.WorkBook, filename string) int {
	rows := workBook.GetSheet(0).Rows
	row14 := rows[14] //for some strange reason sometimes magic number in 14th row
	row15 := rows[15] //sometimes in 15th

	var magicNumber int
	var found bool

	value14A := getValueFromRow(row14, workBook)
	magicNumber, found = getMagicNumberFromValue(value14A)
	if found {
		return magicNumber
	}

	value15A := getValueFromRow(row15, workBook)
	magicNumber, found = getMagicNumberFromValue(value15A)
	if found {
		return magicNumber
	}

	log.Println("value14A:", value14A)
	log.Println("value15A:", value15A)
	panic("Magic number cannot be found for " + filename)
}

func getValueFromRow(row *xls.Row, workBook *xls.WorkBook) string {
	column := row.Cols[0]

	values := column.String(workBook)
	return values[len(values) - 1]
}

func getMagicNumberFromValue(value string) (int, bool) {
	if strings.HasSuffix(value, "000") {
		lastValue := getMagicNumberRawValue(value)
		lastValue = removeNonDigits(lastValue)

		magicNumber, err := strconv.Atoi(lastValue)
		if err == nil {
			return magicNumber, true
		}
	}

	return 0, false
}

func getMagicNumberRawValue(value string) string {
	splittedValues := strings.Split(value, " - ")
	return splittedValues[len(splittedValues) - 1]
}

func removeNonDigits(value string) string {
	regex, err := regexp.Compile("[^\\d]+")
	if err != nil {
		panic(err)
	}

	return regex.ReplaceAllString(value, "")
}