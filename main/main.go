package main

import (
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/logmonitor/main/setting"
	"github.com/logmonitor/main/task"
	"github.com/logmonitor/main/views"
)

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	file_path := r.URL.Path
	token := strings.Split(file_path, "/")
	fp := path.Join(token...)

	http.ServeFile(w, r, fp)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	path := "index.html"
	http.ServeFile(w, r, path)
}

func beforeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// add auth, request path validation here
		fn(w, r)
	}
}

func init() {
	if err := setting.Load(); err != nil {
		log.Fatal("load setting failed, task will not start")
		return
	}
	task.Start()
}

func main() {
	http.HandleFunc("/query/realtime/", beforeHandler(views.RealtimeQuery))
	http.HandleFunc("/query/static/", beforeHandler(views.StaticQuery))
	http.HandleFunc("/static/", beforeHandler(staticFileHandler))
	http.HandleFunc("/", beforeHandler(indexHandler))
	http.ListenAndServe(":8001", nil)
}
