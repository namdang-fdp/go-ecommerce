package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.SetConfigFile("config/local.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("JWT:", viper.GetString("security.jwt.key"))
}
