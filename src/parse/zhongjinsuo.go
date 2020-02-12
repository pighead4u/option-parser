package parse

import (
	"encoding/csv"
	"fmt"
	"github.com/pighead4u/option-parser/src/model"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"log"
	"os"
)

func ParseDataFromZhongJinSuo(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	reader := csv.NewReader(transform.NewReader(file, simplifiedchinese.GB18030.NewDecoder()))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		buildOptionPOFromZhongJinSuo(record)

		//fmt.Println(record)
		//for index, value := range record {
		//	fmt.Println(index)
		//	fmt.Println(value)
		//}
	}
}

func buildOptionPOFromZhongJinSuo(data []string) {
	var type0 = new(model.Content)
	var type1 = new(model.Content)
	var type2 = new(model.Content)
	type0.TransactionDate = data[0]
	type0.ContractCode = data[1]
	type0.Ranking = data[2]
	type0.Company = data[3]
	type0.Volumn = data[4]
	type0.Change = data[5]
	type0.TransactionType = NORMAL
	type0.Insert()
	//fmt.Println(data[3])

	type1.TransactionDate = data[0]
	type1.ContractCode = data[1]
	type1.Ranking = data[2]
	type1.Company = data[6]
	type1.Volumn = data[7]
	type1.Change = data[8]
	type1.TransactionType = BUY
	type1.Insert()

	type2.TransactionDate = data[0]
	type2.ContractCode = data[1]
	type2.Ranking = data[2]
	type2.Company = data[9]
	type2.Volumn = data[10]
	type2.Change = data[11]
	type2.TransactionType = SELL
	type2.Insert()
}
