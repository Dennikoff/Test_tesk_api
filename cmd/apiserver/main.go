package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/Dennikoff/UserTagApi/internal/app/apiserver"
	"log"
)

func main() {
	configPath := "config/config.toml"
	config := apiserver.Config{}
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}
