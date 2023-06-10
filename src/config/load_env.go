package config

import (
	"fmt"
	"io/ioutil"
	"log"
)

func LoadEnv() (env []map[string]string) {
	fileContent, err := ioutil.ReadFile(".env")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	fmt.Println(text)
	return env
}
