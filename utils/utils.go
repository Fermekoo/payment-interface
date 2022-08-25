package utils

import (
	"log"

	"github.com/gobeam/stringy"
	"github.com/spf13/viper"
)

func PaymentName(filename string) string {
	str := stringy.New(filename[:len(filename)-3])

	return "New" + str.UcFirst()
}

func Conf(key string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	return viper.GetString(key)
}
