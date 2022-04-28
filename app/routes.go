package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	if loc, err := time.LoadLocation("UTC"); loc != nil {
		sendJsonResponse(w, map[string]string{"current_time": getNowInLocation(loc)})
	} else {
		sendNotFoundResponse(w, err)
	}
}

func getTimeTz(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	tz := r.URL.Query().Get("tz")
	if tz == "" {
		loc, _ := time.LoadLocation("UTC")
		m["UTC"] = getNowInLocation(loc)
	} else {
		for _, v := range strings.Split(tz, ",") {
			if loc, err := time.LoadLocation(v); loc != nil {
				m[v] = getNowInLocation(loc)
			} else {
				sendNotFoundResponse(w, err)
				fmt.Fprintf(w, "invalid timezone %v", v)
				return
			}
		}
	}
	sendJsonResponse(w, m)
}

func sendNotFoundResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	log.Println(err)
}

func sendJsonResponse(w http.ResponseWriter, m map[string]string) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		sendNotFoundResponse(w, err)
	}
}

func getNowInLocation(loc *time.Location) string {
	return time.Now().In(loc).String()
}
