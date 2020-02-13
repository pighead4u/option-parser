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
			continue
		}

		if len(tmp) != 0 && strings.HasPrefix(tmp, "名次") {
			canEnter = true
			count++
			continue
		}

		if len(tmp) != 0 && strings.HasPrefix(tmp, "总计") {
			data := strings.Split(tmp, "\t")
			var type0 = new(model.Content)
			type0.Ranking = data[0]
			type0.Company = ""
			type0.ContractCode = code
			type0.TransactionDate = date
			type0.Volumn = data[4]
			type0.Change = data[6]
			switch count {
			case 1:
				type0.TransactionType = NORMAL
			case 2:
				type0.TransactionType = BUY
			case 3:
				type0.TransactionType = SELL
			}
			b, err := json.Marshal(type0)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(b))
			type0.Insert()
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

	type0.Ranking = data[0]
	type0.Company = data[2]
	type0.Volumn = data[3]
	type0.Change = data[5]
	type0.TransactionType = transactionType

	return *type0
}
