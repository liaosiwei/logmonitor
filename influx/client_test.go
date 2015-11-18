package client

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

var c *InfluxClient = GetClientInstance()

func TestGetTSV(t *testing.T) {
	now := time.Now().Unix()
	aWeekAgo := now - 7*24*3600
	fmt.Println(time.Unix(aWeekAgo, 0).Format("2006-01-02"))
	res, error := c.QueryByRaw("mydb2", "select * from tm where value > 11500")
	if error != nil {
		t.Error("query failed")
	} else {
		columns := res[0].Series[0].Columns
		var resStr string = strings.Join(columns, "\t") + "\n"
		fmt.Println(resStr)
		fmt.Println(reflect.TypeOf(res[0].Series[0].Columns))
		value := res[0].Series[0].Values[0]
		for i := 0; i < len(value); i++ {
			fmt.Print(reflect.TypeOf(value[i]))
			fmt.Print(": ")
			fmt.Println(value[i])
		}
	}
}

func TestWrite(t *testing.T) {
	err := c.CreateDB("test")
	if err != nil {
		t.Error("create database failed")
	}
	tags := map[string]string{}
	fields := map[string]interface{}{
		"value": 10,
	}
	err = c.Write("test", "tm", tags, fields, time.Now())
	if err != nil {
		t.Error("write database failed")
	}
	res, err := c.QueryByRaw("test", "select * from tm")
	if err != nil {
		t.Error("read from test database failed")
	}
	fmt.Println(res)
}
