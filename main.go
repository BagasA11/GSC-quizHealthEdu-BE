package main

import (
	"BagasA11/GSC-quizHealthEdu-BE/configs"
	"fmt"
)

func main() {
	fmt.Println("main.go")
	fmt.Print("migrating db: ", configs.InitDb(), "\n")
	fmt.Print("db infos: ", configs.GetDB(), "\n")
}
