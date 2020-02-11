package main

import (
	"fmt"
	"github.com/pighead4u/option-parser/src/model"
	"github.com/pighead4u/option-parser/src/parse"
	"os"
	"path/filepath"
)

func main() {
	getFilelist("./shanghai")
	model.CloseDB()
}

func getFilelist(file string) {
	err := filepath.Walk(file, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		parse.ParseDataFromShangHai(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
