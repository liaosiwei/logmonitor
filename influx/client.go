package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/influxdb/influxdb/client/v2"
)

type InfluxClient struct {
	server        string
	port          string
	clientHandler client.Client
}

var influxClient *InfluxClient

func GetClientInstance() *InfluxClient {
	if influxClient == nil {
		influxClient = new(InfluxClient)
		influxClient.server = "http://127.0.0.1"
		influxClient.port = "8086"
		u, _ := url.Parse(influxClient.server + ":" + influxClient.port)
		influxClient.clientHandler = client.NewClient(client.Config{
			URL: u,
		})
	}
	return influxClient
}

func (c *InfluxClient) Close() {
	c.clientHandler.Close()
}

func (c *InfluxClient) QueryByRaw(MyDB string, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: MyDB,
	}
	if response, err := c.clientHandler.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	}
	return res, nil
}

func (c *InfluxClient) CreateDB(database string) error {
	_, err := c.clientHandler.Query(
		client.Query{
			Command: fmt.Sprintf("create database %s", database)})
	return err
}

func (c *InfluxClient) Write(
	database string,
	measurement string,
	tags map[string]string,
	fields map[string]interface{},
	theTime time.Time) error {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: "s",
	})
	pt, err := client.NewPoint(measurement, tags, fields, theTime)
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	c.clientHandler.Write(bp)
	return err
}

func GetTSV(res []client.Result) string {
	var buffer bytes.Buffer
	columns := res[0].Series[0].Columns
	buffer.WriteString(strings.Join(columns, "\t"))
	buffer.WriteString("\n")
	values := res[0].Series[0].Values
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			if j != len(values[i])-1 {
				buffer.WriteString((values[i][j]).(string))
				buffer.WriteString("\t")
			} else {
				if values[i][j] == nil {
					buffer.WriteString("0")
				} else {
					buffer.WriteString(string(values[i][j].(json.Number)))
				}

			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
