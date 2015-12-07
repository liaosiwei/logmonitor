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
	Attr     struct {
		DbName string
		File   string
	}
}

var Config Configuration

func Load(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("open file failed: ", filepath)
		return err
	}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("load config file wrong in setting.Load ", err)
		return err
	}
	return nil
}
