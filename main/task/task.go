package task

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/logmonitor/influx"
	"github.com/logmonitor/main/setting"
	"github.com/logmonitor/scheduler"
	"github.com/logmonitor/shcmd"
)

var config setting.Configuration

func Start() {
	config = setting.Config
	_, err := scheduler.Schedule(runWebproxyStatic, config.Webproxy.Schedule...)
	if err != nil {
		log.Fatal("start webproxy static task failed")
	}
	_, err = scheduler.Schedule(runAcStatic, config.Ac.Schedule...)
	if err != nil {
		log.Fatal("start ac static task failed")
	}
}

func runWebproxyStatic() {
	_, err := shcmd.RunWithin(config.Webproxy.Cmd, time.Duration(config.Webproxy.Timeout)*time.Hour)
	if err != nil {
		log.Fatal("run webproxy static task failed")
		return
	}
	collectWebproxyStatic()
}

func runAcStatic() {
	_, err := shcmd.RunWithin(config.Ac.Cmd, time.Duration(config.Ac.Timeout)*time.Hour)
	if err != nil {
		log.Fatal("run ac static task failed")
		return
	}
	collectAcStatic()
}

func collectWebproxyStatic() {
	file, err := os.Open(config.Webproxy.File)
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

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	timestap := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.Local)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "90%" {
			flag = 1
			continue
		}
		if flag == 1 {
			flag = 0
			value, _ := strconv.ParseFloat(line, 32)
			fields := map[string]interface{}{
				"value": value,
			}

			err := c.Write(config.Webproxy.DbName, measure[count], tags, fields, timestap)
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

func collectAcStatic() {
	file, err := os.Open(config.Ac.File)
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
		"tm", "GT", "FT",
	}
	tags := map[string]string{}

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	timestap := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.Local)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasSuffix(line, "point") {
			flag = 1
			continue
		}
		if flag == 1 {
			flag = 0
			value, _ := strconv.ParseFloat(line, 32)
			value = value / 1000000
			fields := map[string]interface{}{
				"value": value,
			}
			err := c.Write(config.Ac.DbName, measure[count], tags, fields, timestap)
			if err != nil {
				log.Fatal("write ac database failed")
			}
			count += 1
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("reading standard input:", err)
	}
}
