package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitConfig(files ...string) {
	err := godotenv.Overload(files...)
	if err != nil {
		log.Fatalf("cannot load config with error %s", err.Error())
	}
}

func GetString(key CfgKeys) string {
	val := os.Getenv(string(key))
	return val
}

func GetInt8(key CfgKeys) int8 {
	val := os.Getenv(string(key))

	valInt8, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return int8(valInt8)
}
