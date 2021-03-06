package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/logmonitor/influx"
)

func RealtimeQuery(w http.ResponseWriter, r *http.Request) {
	c := client.GetClientInstance()
	today := time.Now()
	now := today.Format("2006-01-02")
	aWeekAgo := time.Unix(today.Unix()-7*24*3600, 0).Format("2006-01-02")

	queryStr := fmt.Sprintf("select percentile(value, 90) from tm, btm where time > '%s' and time <= '%s' group by time(1d) fill(0)",
		aWeekAgo, now)
	//originQueryStr := "select percentile(value, 90) from tm, btm where time > now() - 1w and time < now() group by time(1d)"
	//fmt.Println(queryStr)
	res, err := c.QueryByRaw("mydb2", queryStr)
	if err != nil {
		log.Fatal(err)
	}
	js, err := json.Marshal(res)
	if err != nil {
		//		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func StaticQuery(w http.ResponseWriter, r *http.Request) {
	database := r.URL.Query().Get("db")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	if len(database) == 0 {
		return
	}
	c := client.GetClientInstance()
	var queryStr string
	if database == "ac" {
		queryStr = fmt.Sprintf("select * from tm, GT, FT, lbt, rGT where time >= '%s' and time <= '%s' fill(0)", from, to)
	}
	if database == "webproxy" {
		queryStr = fmt.Sprintf("select * from tm, btm, ctm, size where time >= '%s' and time <= '%s' fill(0)", from, to)
	}
	if database == "attr" {
		queryStr = fmt.Sprintf("select * from tm, dt, fct, pt where time >= '%s' and time <= '%s' fill(0)", from, to)
	}
	res, err := c.QueryByRaw(database, queryStr)
	if err != nil {
		log.Fatal("query database failed: ", database)
		return
	}
	js, err := json.Marshal(res)
	if err != nil {
		log.Fatal("jsonlize results failed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
