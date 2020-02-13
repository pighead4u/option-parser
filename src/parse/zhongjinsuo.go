package parse

import (
	"bufio"
	"fmt"
	"github.com/pighead4u/option-parser/src/model"
	"log"
	"os"
	"strings"
)

func ParseDataFromZhongJinSuo(path string) {
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
		if len(tmp) != 0 && strings.HasPrefix(tmp, "合约:") {
			code, date = getCodeAndDateFromZhongJinSuo(tmp)
			continue
		}

		if len(tmp) != 0 && strings.HasPrefix(tmp, "名次") {
			canEnter = true
			count++
			continue
		}

		if len(tmp) != 0 && strings.HasPrefix(tmp, "合计") {
			var type0, type1, type2 = buildTotalFromZhongJinSuo(tmp)
			type0.TransactionDate = date
			type0.ContractCode = code
			type0.Insert()

			type1.TransactionDate = date
			type1.ContractCode = code
			type1.Insert()

			type2.TransactionDate = date
			type2.ContractCode = code
			type2.Insert()
			canEnter = false
			continue
		}

		if canEnter {
			var type0, type1, type2 = buildOptionPOFromZhongJinSuo(tmp)
			type0.TransactionDate = date
			type0.ContractCode = code
			type0.Insert()

			type1.TransactionDate = date
			type1.ContractCode = code
			type1.Insert()

			type2.TransactionDate = date
			type2.ContractCode = code
			type2.Insert()

			//b, err := json.Marshal(type0)
			//if err != nil {
			//	fmt.Println("error:", err)
			//}
			//fmt.Println(string(b))
			//
			//b, err = json.Marshal(type1)
			//if err != nil {
			//	fmt.Println("error:", err)
			//}
			//fmt.Println(string(b))
			//
			//b, err = json.Marshal(type2)
			//if err != nil {
			//	fmt.Println("error:", err)
			//}
			//fmt.Println(string(b))
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func buildTotalFromZhongJinSuo(content string) (model.Content, model.Content, model.Content) {
	var type0 = new(model.Content)
	var type1 = new(model.Content)
	var type2 = new(model.Content)
	var data = strings.Split(content, "	")

	type0.Ranking = strings.ReplaceAll(data[0], " ", "")
	type0.Company = ""
	type0.Volumn = strings.ReplaceAll(data[2], " ", "")
	type0.Change = strings.ReplaceAll(data[3], " ", "")
	type0.TransactionType = NORMAL

	type1.Ranking = type0.Ranking
	type1.Company = ""
	type1.Volumn = strings.ReplaceAll(data[6], " ", "")
	type1.Change = strings.ReplaceAll(data[7], " ", "")
	type1.TransactionType = BUY

	type2.Ranking = type0.Ranking
	type2.Company = ""
	type2.Volumn = strings.ReplaceAll(data[10], " ", "")
	type2.Change = strings.ReplaceAll(data[11], " ", "")
	type2.TransactionType = SELL

	return *type0, *type1, *type2
}

func getCodeAndDateFromZhongJinSuo(content string) (string, string) {
	index := strings.Index(content, "交易日:")
	start := len("合约:")
	code := content[start:index]
	length := len("交易日:")
	date := strings.TrimSpace(content[index+length:])

	return code, date
}

func buildOptionPOFromZhongJinSuo(content string) (model.Content, model.Content, model.Content) {
	var type0 = new(model.Content)
	var type1 = new(model.Content)
	var type2 = new(model.Content)
	var data = strings.Split(content, "	")

	type0.Ranking = strings.ReplaceAll(data[0], " ", "")
	type0.Company = strings.ReplaceAll(data[1], " ", "")
	type0.Volumn = strings.ReplaceAll(data[2], " ", "")
	type0.Change = strings.ReplaceAll(data[3], " ", "")
	type0.TransactionType = NORMAL

	type1.Ranking = type0.Ranking
	type1.Company = strings.ReplaceAll(data[5], " ", "")
	type1.Volumn = strings.ReplaceAll(data[6], " ", "")
	type1.Change = strings.ReplaceAll(data[7], " ", "")
	type1.TransactionType = BUY

	type2.Ranking = type0.Ranking
	type2.Company = strings.ReplaceAll(data[9], " ", "")
	type2.Volumn = strings.ReplaceAll(data[10], " ", "")
	type2.Change = strings.ReplaceAll(data[11], " ", "")
	type2.TransactionType = SELL

	return *type0, *type1, *type2
}
