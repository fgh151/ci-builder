package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	checkErr(err)
	serve()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
