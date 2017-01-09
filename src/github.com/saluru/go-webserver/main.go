package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const url = "https://www.placebear.com/400/400"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	// Simple Get request to pull Dog image.
	// We recieve a *Response and an error
	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(w, " Sorry there is an error occured, ", err)
	}
	// Make sure to close the response after reading all bytes.
	defer resp.Body.Close()

	//Lets  read all the bytes of the image.
	//	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v ", err)
	}

	//	w.Header().Set("Content-Length", fmt.Sprint(resp.ContentLength))
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	if _, err = io.Copy(w, resp.Body); err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}
	elapsed := time.Since(start).Seconds()

	fmt.Printf(" Elapsed Time to get is : %v\n ", elapsed)
}

func main() {

	http.HandleFunc("/", helloHandler)

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

	//  resp, err := http.Get("https://www.placebear.com/400/400")
}
