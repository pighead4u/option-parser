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

func ParseDataFromZhengZhou(path string) {
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
		if len(tmp) != 0 && strings.HasPrefix(tmp, "合约：") {
			fmt.Println(tmp)
			code, date = getCodeAndDateFromZhengZhou(tmp)
			fmt.Println(code)
			fmt.Println(date)
			continue
		}

		if len(code) != 0 && len(tmp) != 0 && strings.HasPrefix(tmp, "名次") {
			canEnter = true
			continue
		}

		if len(code) != 0 && len(tmp) != 0 && strings.HasPrefix(tmp, "合计") {
			canEnter = false
			continue
		}

		if canEnter {
			var type0, type1, type2 = getOptionPOFromZhengZhou(tmp)
			type0.TransactionDate = date
			type0.ContractCode = code
			type0.Insert()

			type1.TransactionDate = date
			type1.ContractCode = code
			type1.Insert()

			type2.TransactionDate = date
			type2.ContractCode = code
			type2.Insert()

			b, err := json.Marshal(type0)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(b))

			b, err = json.Marshal(type1)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(b))

			b, err = json.Marshal(type2)
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

func getCodeAndDateFromZhengZhou(content string) (string, string) {
	var data = strings.Split(content, " ")
	var lenOfCode = len(data[0])
	var tmp = len("合约：")
	var code = data[0][tmp:lenOfCode]
	var date = strings.ReplaceAll(data[15], "-", "")

	return code, date
}

func getOptionPOFromZhengZhou(content string) (model.Content, model.Content, model.Content) {
	var type0 = new(model.Content)
	var type1 = new(model.Content)
	var type2 = new(model.Content)
	var data = strings.Split(content, "|")

	type0.Ranking = strings.ReplaceAll(data[0], " ", "")
	type0.Company = strings.ReplaceAll(data[1], " ", "")
	type0.Volumn = strings.ReplaceAll(data[2], " ", "")
	type0.Change = strings.ReplaceAll(data[3], " ", "")
	type0.TransactionType = NORMAL

	type1.Ranking = strings.ReplaceAll(data[0], " ", "")
	type1.Company = strings.ReplaceAll(data[4], " ", "")
	type1.Volumn = strings.ReplaceAll(data[5], " ", "")
	type1.Change = strings.ReplaceAll(data[6], " ", "")
	type1.TransactionType = BUY

	type2.Ranking = strings.ReplaceAll(data[0], " ", "")
	type2.Company = strings.ReplaceAll(data[7], " ", "")
	type2.Volumn = strings.ReplaceAll(data[8], " ", "")
	type2.Change = strings.ReplaceAll(data[9], " ", "")
	type2.TransactionType = SELL

	return *type0, *type1, *type2
}
