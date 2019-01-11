package main

import (
	"fmt"
	"net/http"

	"github.com/c0z0/go-now/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("content-type", "text/html")
	date := utils.GetTime()
	fmt.Fprintf(w, "Hello from go! The date is %s", date)
}
