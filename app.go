package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type pingResponse struct {
	Status string `json:"status"`
}

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/ping", handlePing)
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// 獲取當前 UTC 時間
	nowUTC := time.Now().UTC()

	// 轉換為香港時區時間
	hongKongLocation, _ := time.LoadLocation("Asia/Hong_Kong")
	nowHongKong := nowUTC.In(hongKongLocation)

	data := map[string]interface{}{
		"CurrentTime": nowHongKong.Format("2006-01-02 15:04:05"),
	}

	t.ExecuteTemplate(w, "index.html.tmpl", data)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	resp := pingResponse{
		Status: "ok",
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

