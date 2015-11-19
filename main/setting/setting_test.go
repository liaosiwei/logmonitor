package setting

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	err := Load()
	if err != nil {
		t.Error("load config failed")
	}
	str, _ := json.MarshalIndent(Config, "", "\t")
	fmt.Println(string(str))
}
