package helpers

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv(path string){
	if mode := os.Getenv("NODE_ENV"); mode != "production" {
		err := godotenv.Load(path)
		if err != nil {
			log.Fatal("Error load .env file ", err.Error())
		}
	}
}

func GetEnvStr(name string, d string)string{
	var value string = os.Getenv(name)
	if value == "" && d != "" {
		return d
	}

	if value == "" {
		log.Fatal("Env ", name, " key notfound!")
	}

	return value
}

func GetEnvNum(name string, d int)int{
	dbPort, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		if d != 0 {
			return d
		}
		log.Fatal("error ", err.Error())
	}

	return dbPort
}