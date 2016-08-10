package main

import (
	_ "github.com/tealeg/xlsx"
	"io/ioutil"
	"fmt"
)

const Dir = "files"

func main() {
	files, _ := ioutil.ReadDir(Dir)
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
