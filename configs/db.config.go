package configs

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// "BagasA11/GSC-quizHealthEdu-BE/api/models"

//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//
// ""
var dbClient *gorm.DB
var Dsn = "root@tcp(127.0.0.1:3306)/gsc-quiz-healthedu?charset=utf8mb4&parseTime=True&loc=Local"

func InitDb() error {
	var err error
	// dsn :=
	dbClient, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	dbClient.AutoMigrate(&models.User{}, &models.Quiz{}, &models.Question{}, &models.Option{}, &models.Score{}, &models.TopUp{}, &models.Transaction{})
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}
