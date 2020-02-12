package parse

import (
	"bufio"
	"fmt"
	"github.com/pighead4u/option-parser/src/model"
	"log"
	"os"
	"strings"
)

const NORMAL = "normal"
const BUY = "buy"
const SELL = "sell"

func ParseDataFromShangHai(path string) {
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
	for scanner.Scan() {
		var tmp = scanner.Text()
		if len(tmp) == 0 {
			fmt.Println("==========================================")
			continue
		}
		if len(tmp) != 0 && strings.HasPrefix(tmp, "合约代码 ：") {
			code, date = getCodeAndDateFromShangHai(tmp)
			continue
		}

		if len(tmp) != 0 && strings.HasPrefix(tmp, "名次,") {
			canEnter = true
			continue
		}

		if len(tmp) != 0 && strings.HasPrefix(tmp, "合计,") {
			canEnter = false
			continue
		}

		if canEnter {
			var type0, type1, type2 = getOptionPOFromShangHai(tmp)
			type0.TransactionDate = date
			type0.ContractCode = code
			type0.Insert()

			type1.TransactionDate = date
			type1.ContractCode = code
			type1.Insert()

			type2.TransactionDate = date
			type2.ContractCode = code
			type2.Insert()

		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getCodeAndDateFromShangHai(content string) (string, string) {
	var data = strings.Split(content, " ")
	var code = data[1][3:]
	var tmp = strings.ReplaceAll(data[25], "-", "")
	var date = strings.ReplaceAll(tmp, ",", "")

	return code, date
}

func getOptionPOFromShangHai(content string) (model.Content, model.Content, model.Content) {
	var type0 = new(model.Content)
	var type1 = new(model.Content)
	var type2 = new(model.Content)
	var data = strings.Split(content, ",")
	type0.Ranking = data[0]
	type0.Company = strings.TrimSpace(data[1])
	type0.Volumn = data[2]
	type0.Change = data[3]
	type0.TransactionType = NORMAL

	type1.Ranking = data[4]
	type1.Company = strings.TrimSpace(data[5])
	type1.Volumn = data[6]
	type1.Change = data[7]
	type1.TransactionType = BUY

	type2.Ranking = data[8]
	type2.Company = strings.TrimSpace(data[9])
	type2.Volumn = data[10]
	type2.Change = data[11]
	type2.TransactionType = SELL

	return *type0, *type1, *type2
}
