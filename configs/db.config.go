package configs

import (
	"gorm.io/gorm"
)

// "BagasA11/go-POS/api/models"

//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//
// ""
var dbClient *gorm.DB
var Dsn = "root@tcp(127.0.0.1:3306)/golangpos?charset=utf8mb4&parseTime=True&loc=Local"

// func InitDb() error {
// 	var err error
// 	// dsn :=
// 	dbClient, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	dbClient.AutoMigrate(&models.Retail{}, &models.Cashier{}, &models.Item{}, &models.ItemDetail{}, &models.Sale{}, &models.SaleDetail{}, &models.Return{})
// 	return nil
// }

func GetDB() *gorm.DB {
	return dbClient
}
