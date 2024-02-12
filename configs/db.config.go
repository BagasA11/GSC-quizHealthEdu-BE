package configs

import (
	"BagasA11/GSC-quizHealthEdu-BE/api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// "fmt"
// 	"os"

//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//
// ""
var dbClient *gorm.DB

// var Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//
//	os.Getenv("MYSQL_USERNAME"),
//	os.Getenv("MYSQL_PASSWORD"),
//	os.Getenv("MYSQL_HOST"),
//	os.Getenv("MYSQL_PORT"),
//	os.Getenv("MYSQL_DATABASE"),
//
// )
// "root":""@tcp("127.0.0.1":"3306")/"gsc-quiz-healthedu"?charset=utf8mb4&parseTime=True&loc=Local
var Dsn = "root:@tcp(127.0.0.1:3306)/gsc-quiz-healthedu?charset=utf8mb4&parseTime=True&loc=Local"

func InitDb() error {
	var err error
	// dsn :=
	dbClient, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	dbClient.AutoMigrate(&models.User{}, &models.Quiz{}, &models.Question{}, &models.Option{}, &models.Score{}, &models.TopUp{}, &models.Transaction{}, &models.BlacklistToken{})
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}
