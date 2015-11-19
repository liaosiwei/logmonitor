package setting

import (
	"encoding/json"
	"log"
	"os"
)

type Assembler struct {
	DbName   string
	File     string
	Cmd      string
	Timeout  int
	Schedule []int
}

type Configuration struct {
	Webproxy Assembler
	Ac       Assembler
}

var config Configuration

func Load() error {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("load config file wrong", err)
		return err
	}
	return nil
}
