package models

import (
	"log"
)

func TableCreate() {
	db, err := DbConn()
	if err != nil {
		log.Print(err)
	}
	var stockDetail StockDetail
	if !db.HasTable(&stockDetail) {
		db.CreateTable(&stockDetail)
	}
	defer db.Close()
}

func StockNumbSave([]string, []string) bool {

	return true
}

func StockInfoSave(detail []string) {
	db, err := DbConn()
	if err != nil {
		log.Print(err)
	}
	var stockDetail StockDetail
	stockDetail.Status = detail[0]
	stockDetail.Number = detail[1]
	stockDetail.Name = detail[2]
	stockDetail.Price = detail[3]
	stockDetail.StockRnf = detail[4]
	stockDetail.StockRnfPercent = detail[5]
	stockDetail.CurrentTrade = detail[6]
	stockDetail.TotalTrade = detail[7]
	stockDetail.BuyPrice = detail[8]
	stockDetail.BuyQuantity = detail[9]
	stockDetail.SellPrice = detail[10]
	stockDetail.SellQuantity = detail[11]
	stockDetail.OpenPrice = detail[12]
	stockDetail.Highest = detail[13]
	stockDetail.Lowest = detail[14]
	stockDetail.CreatedAt = detail[15]
	db.Create(&stockDetail)
	defer db.Close()

}
