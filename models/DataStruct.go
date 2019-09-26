package models

type StockInfo struct {
	id        int
	StockNumb string
	StockName string
}

type DBConnInfo struct {
	USERNAME     string
	USERPASSWORD string
	DBHOST       string
	DBNAME       string
}
type StockDetail struct {
	Id              uint64 `gorm:"type:int auto_increment;primary_key;not null"`
	Status          string `gorm:"type:nvarchar(4);not null"`
	Number          string `gorm:"type:varchar(6);not null"`
	Name            string `gorm:"type:nvarchar(10);not null"`
	Price           string `gorm:"type:varchar(10);not null"`
	StockRnf        string `gorm:"type:varchar(10);not null"`
	StockRnfPercent string `gorm:"type:varchar(10);not null"`
	CurrentTrade    string `gorm:"type:varchar(10);not null"`
	TotalTrade      string `gorm:"type:varchar(10);not null"`
	BuyPrice        string `gorm:"type:varchar(10);not null"`
	BuyQuantity     string `gorm:"type:varchar(10);not null"`
	SellPrice       string `gorm:"type:varchar(10);not null"`
	SellQuantity    string `gorm:"type:varchar(10);not null"`
	OpenPrice       string `gorm:"type:varchar(10);not null"`
	Highest         string `gorm:"type:varchar(10);not null"`
	Lowest          string `gorm:"type:varchar(10);not null"`
	CreatedAt       string `gorm:"type:datetime;not null"`
}
