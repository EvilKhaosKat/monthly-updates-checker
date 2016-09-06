package main

import (
	"testing"
	"fmt"
)

func TestGetMagicNumberFromValue(t *testing.T) {
	number, found := getMagicNumberFromValue("test test\nstest - 75 000")

	if !found {
		t.Error("Number must be found!")
	}

	CorrectValue := 75000
	if number != CorrectValue {
		t.Error(fmt.Sprintf("Number must be '%v', but equals '%v'", CorrectValue, number))
	}
}

func TestRemoveNonDigits(t *testing.T) {
	rawString := "q 75 000"
	handledString := removeNonDigits(rawString)

	CorrectValue := "75000"
	if handledString != CorrectValue {
		t.Fatal(fmt.Sprintf("Handled strings must be '%v', but equals '%v'",
			CorrectValue,
			handledString))
	}
}
