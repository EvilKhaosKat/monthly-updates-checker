package main

import (
	"testing"
	"fmt"
)

func TestGetDate(t *testing.T) {
	year, month := getDate("11.12.xls")

	CorrectYearValue := 11
	if year != CorrectYearValue {
		t.Error(fmt.Sprintf("Year must be '%v', but equals '%v'", CorrectYearValue, year))
	}

	CorrectMonthValue := 12
	if month != CorrectMonthValue {
		t.Error(fmt.Sprintf("Month must be '%v', but equals '%v'", CorrectMonthValue, month))
	}
}