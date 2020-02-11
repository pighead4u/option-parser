package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./data/options.db")
	if err != nil {
		panic(err)
	}
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
type Content struct {
	gorm.Model
	Ranking         string `json:"ranking"`
	Company         string `json:"company"`
	Volumn          string `json:"volumn"`
	Change          string `json:"change"`
	TransactionType string `json:"transactionType"`
	ContractCode    string `json:"contractCode"`
	TransactionDate string `json:"transactionDate"`
}

func (po *Content) Insert() {
	db.Create(po)
}

func (po *Content) Query() {

}

func CloseDB() {
	db.Close()
}
