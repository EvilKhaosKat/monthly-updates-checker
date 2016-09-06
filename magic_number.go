package main

import (
	"github.com/extrame/xls"
	"log"
	"strings"
	"strconv"
	"regexp"
	"fmt"
)

func getMagicNumber(filename string) int {
	workBook, err := xls.Open(fmt.Sprintf("%v/%v", Dir, filename), "UTF-8")
	if err != nil {
		log.Fatal(err)
	}

	return getMagicNumberByXls(workBook, filename)
}

func getMagicNumberByXls(workBook *xls.WorkBook, filename string) int {
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