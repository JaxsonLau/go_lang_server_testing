package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 獲取當前 UTC 時間
		nowUTC := time.Now().UTC()

		// 轉換為香港時區時間
		hongKongLocation, _ := time.LoadLocation("Asia/Hong_Kong")
		nowHongKong := nowUTC.In(hongKongLocation)

		data := map[string]interface{}{
			"CurrentTime": nowHongKong.Format("2006-01-02 15:04:05"),
		}

		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
