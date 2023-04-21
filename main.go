package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"microservices/details"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")

	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running ")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")

	hostname, err := details.GetHostName()
	if err != nil {
		panic(err)
	}

	myIP := details.GetIP()

	response := map[string]string{
		"hostname": hostname,
		"ip":       myIP,
	}

	json.NewEncoder(w).Encode(response)
}

func ZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health")

	vars := mux.Vars(r)

	ZipCode := vars["zip_code"]
	CountryCode := vars["country_code"]
	resp, err := http.Get("https://api.zippopotam.us/" + CountryCode + "/" + ZipCode)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(body))
	fmt.Fprintf(w, string(body))

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	r.HandleFunc("/{country_code}/{zip_code}", ZipCodeHandler)
	// Start the server
	log.Println("Web server has started!!!")
	log.Fatal(http.ListenAndServe(":80", r))
}
