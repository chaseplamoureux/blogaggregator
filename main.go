package main

import (
	"fmt"
	"github.com/chaseplamoureux/blogaggregator/internal/config"
	"log"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	conf.SetUser("Chase Lamoureux")

	conf, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf.DB_URL)
	fmt.Println(conf.Username)
}
