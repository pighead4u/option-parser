package main

import (
	"fmt"
	"github.com/pighead4u/option-parser/src/model"
	"github.com/pighead4u/option-parser/src/parse"
	"os"
	"path/filepath"
)

const DALIAN = "dalian"
const SHANGHAI = "shanghai"
const ZHENGZHOU = "zhengzhou"
const ZHONGJIANSUO = "zhongjiansuo"

func main() {
	getFilelist("./zhongjinsuo", ZHONGJIANSUO)
	getFilelist("./zhengzhou", ZHENGZHOU)
	getFilelist("./shanghai", SHANGHAI)
	getFilelist("./dalian", DALIAN)
	model.CloseDB()
}

func getFilelist(file, city string) {
	err := filepath.Walk(file, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		switch city {
		case ZHONGJIANSUO:
			parse.ParseDataFromZhongJinSuo(path)
		case ZHENGZHOU:
			parse.ParseDataFromZhengZhou(path)
		case SHANGHAI:
			parse.ParseDataFromShangHai(path)
		case DALIAN:
			parse.ParseDataFromDaLian(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
