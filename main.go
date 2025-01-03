package main

import (
	"fmt"
	"log"

	"github.com/JakubKyhos/blogaggregator/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.SetUser("Jakub")
	if err != nil {
		log.Fatal(err)
	}

	updatedCfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedCfg.DBUrl)
}
