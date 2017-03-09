package main

import (
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gorilla/mux"
)

type Header struct {
	Ipaddress string `json:"ipaddress"`
	Language  string `json:"language"`
	Software  string `json:"software"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", rootHandler)
	http.Handle("/", r)

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	ipaddress := r.RemoteAddr
	language := strings.Split(r.Header["Accept-Language"][0], ",")[0]

	rgx := regexp.MustCompile(`\((.*?)\)`)
	rgxResult := rgx.FindStringSubmatch(r.UserAgent())[0]

	software := rgxResult[1 : len(rgxResult)-1]

	json.NewEncoder(w).Encode(Header{
		ipaddress,
		language,
		software,
	})
}
