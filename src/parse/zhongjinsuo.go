package parse

import (
	"encoding/csv"
	"fmt"
	"github.com/pighead4u/option-parser/src/model"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func ParseDataFromZhongJinSuo(path string) {
	r, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	record := csv.NewReader(strings.NewReader(string(r)))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := record.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		buildOptionPOFromZhongJinSuo(record)

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
	//type0.Insert()
	fmt.Println(data[3])

	type1.TransactionDate = data[0]
	type1.ContractCode = data[1]
	type1.Ranking = data[2]
	type1.Company = data[6]
	type1.Volumn = data[7]
	type1.Change = data[8]
	type1.TransactionType = BUY
	//type1.Insert()

	type2.TransactionDate = data[0]
	type2.ContractCode = data[1]
	type2.Ranking = data[2]
	type2.Company = data[9]
	type2.Volumn = data[10]
	type2.Change = data[11]
	type2.TransactionType = SELL
	//type2.Insert()
}
