package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var appVersion = "2006.1.2"
var appGitShortSha = "undefined"

type AppInfo struct {
	Version string `json:"version"`
}

type Response struct {
	AppInfo AppInfo `json:"app_info"`
	Method  string  `json:"method"`
	Path    string  `json:"path"`
}

func appInfo() (appInfo AppInfo) {
	var version = fmt.Sprintf("%s+%s", appVersion, appGitShortSha)
	appInfo = AppInfo{
		Version: version,
	}
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		AppInfo: appInfo(),
		Method:  r.Method,
		Path:    r.URL.Path,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := "8080"
	if _port := os.Getenv("PORT"); _port != "" {
		port = _port
	}
	listenAddr := fmt.Sprintf(":%s", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
