package parse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pighead4u/option-parser/src/model"
	"log"
	"os"
	"strings"
)

func ParseDataFromDaLian(path string) {
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var code = ""
	var date = ""
	var canEnter = false
	var count = 0
	for scanner.Scan() {
		var tmp = scanner.Text()
		if len(tmp) == 0 {
			fmt.Println("==========================================")
			continue
		}
		if len(tmp) != 0 && strings.HasPrefix(tmp, "合约代码：") {
			code, date = getCodeAndDateFromDaLian(tmp)
			println(code)
			println(date)
			continue
		}

		if len(tmp) != 0 && strings.HasPrefix(tmp, "名次") {
			canEnter = true
			count++
			continue
		}

		if len(tmp) != 0 && strings.HasPrefix(tmp, "总计") {
			canEnter = false
			continue
		}

		if canEnter {
			var content model.Content
			switch count {
			case 1:
				content = getOptionPOFromDaLian(tmp, NORMAL)
			case 2:
				content = getOptionPOFromDaLian(tmp, BUY)
			case 3:
				content = getOptionPOFromDaLian(tmp, SELL)
			}
			content.ContractCode = code
			content.TransactionDate = date
			content.Insert()

			b, err := json.Marshal(content)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(b))
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getCodeAndDateFromDaLian(content string) (string, string) {
	var lenOfCode = len("合约代码：")
	var index = strings.Index(content, "	")
	var code = content[lenOfCode:index]
	index = strings.Index(content, "Date：")
	var lenOfDate = len("Date：")
	var date = strings.ReplaceAll(content[index+lenOfDate:index+lenOfDate+10], "-", "")

	return code, date
}

func getOptionPOFromDaLian(content, transactionType string) model.Content {
	var type0 = new(model.Content)
	var data = strings.Split(content, "\t")
	fmt.Println(content)
	fmt.Println(len(data))
	fmt.Println("0:" + data[0])
	fmt.Println("1:" + data[1])
	fmt.Println("2:" + data[2])
	fmt.Println("3:" + data[3])
	fmt.Println("4:" + data[4])
	fmt.Println("5:" + data[5])
	fmt.Println("6:" + data[6])
	//fmt.Println(data[7])
	type0.Ranking = data[0]
	type0.Company = data[2]
	type0.Volumn = data[3]
	type0.Change = data[5]
	type0.TransactionType = transactionType

	return *type0
}
