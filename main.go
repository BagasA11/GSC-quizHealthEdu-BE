package main

import (
	"BagasA11/GSC-quizHealthEdu-BE/configs"
	"BagasA11/GSC-quizHealthEdu-BE/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("main.go...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	fmt.Print("migrating db: ", configs.InitDb(), "\n")
	r := gin.Default()
	routes.RegisterRoutes(r)

}
