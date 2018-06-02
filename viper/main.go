package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	key := "title"
	if viper.IsSet(key) {
		fmt.Println(viper.GetString(key))
	}

	fmt.Println(viper.Get("clients.data"))
}
