package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "./data/options.db")
	if err != nil {
		fmt.Println("error:", err)
	}

	db.Close()
}

/**
 * 返回给前端的对象
 *
 **/
type OptionVO struct {
	Name            string `json:"name"`
	Exchange        string `json:"exchange"`
	Ranking         string `json:"ranking"`
	Company         string `json:"company"`
	Volumn          string `json:"volumn"`
	Change          string `json:"change"`
	TransactionType string `json:"transactionType"`
	ContractCode    string `json:"contractCode"`
	TransactionDate string `json:"transactionDate"`
}

/**
 *
 * 数据库对应的表
 **/
type OptionPO struct {
	Ranking         string `json:"ranking"`
	Company         string `json:"company"`
	Volumn          string `json:"volumn"`
	Change          string `json:"change"`
	TransactionType string `json:"transactionType"`
	ContractCode    string `json:"contractCode"`
	TransactionDate string `json:"transactionDate"`
}

func (po *OptionPO) Insert() {
	db.Create(po)
}

func (po *OptionPO) Query() {

}
