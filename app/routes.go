package app

import (
	"fmt"
	"net/http"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format(time.RFC3339))
}
