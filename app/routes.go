package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

func getTime(w http.ResponseWriter, r *http.Request) {
	if loc, err := LoadLocation(r); loc != nil {
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(TimeResponse{CurrentTime: time.Now().In(loc).Format(time.RFC3339)}); err != nil {
			log.Panic(err)
		}
	} else {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("invalid timezone"))
	}
}

func LoadLocation(r *http.Request) (*time.Location, error) {
	values := r.URL.Query()
	tz := "UTC"
	if values.Has("tz") {
		tz = values.Get("tz")
	}
	return time.LoadLocation(tz)
}
