package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	checkerr(err)

	//var wg sync.WaitGroup
	//wg.Add(1)
	//go serve(&wg)
	//wg.Wait()

	//clone("./test", "git@github.com:fgh151/shop-engine.git")

	up("./test")
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
