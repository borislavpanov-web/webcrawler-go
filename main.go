package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/borislavpanov-web/webcrawler-go/crawler"
)

func handler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	data, err := crawler.ScrapeAllPages(1, 150)
	if err != nil {
		http.Error(w, "Failed to scrape data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	elapsedTime := time.Since(startTime)

	response := fmt.Sprintf(
		"<html><body><pre>Time taken in seconds: %v\n\n%s</pre></body></html>",
		elapsedTime.Seconds(),
		data,
	)

	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
