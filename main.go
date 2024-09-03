package main

import (
	"ArticleVista/config"
	"fmt"
	"log"
)

func main() {
	db, err := config.NewConnection()
	if err != nil {
		log.Fatalln("Cannot connect to database: ", err.Error())
	}
	fmt.Println(db)

}
