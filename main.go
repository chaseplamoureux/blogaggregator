package main

import (
	"fmt"
	"github.com/chaseplamoureux/blogaggregator/internal/config"
	"log"
)

func main() {
	conf, err := config.Read(".gatorconfig.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf.DB_URL)
}
