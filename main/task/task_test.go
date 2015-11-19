package task

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"fmt"
	"testing"

	"github.com/logmonitor/influx"
)

func TestCollectWebproxy(t *testing.T) {
	file, err := os.Open(baseDirWebproxy + "myresult.txt")
	if err != nil {
		log.Fatal("open file failed: ", err)
		return
	}
	c := client.GetClientInstance()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// Note: the logic below is strongly dependent of the content of myresult.txt
	// which may be changed later, and this will cause the function failing
	flag := 0
	count := 0
	measure := [...]string{
		"tm", "btm", "ctm", "size",
	}
	tags := map[string]string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "90%" {
			flag = 1
			continue
		}
		if flag == 1 {
			flag = 0
			fmt.Println("get line: ", line)
			value, _ := strconv.ParseFloat(line, 32)
			fields := map[string]interface{}{
				"value": value,
			}
			err := c.Write("webproxy", measure[count], tags, fields, time.Now())
			if err != nil {
				log.Fatal("write database failed")
			}
			count += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("reading webproxy standard input:", err)
	}	
}

