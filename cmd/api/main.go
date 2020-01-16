package main

import (
	"github.com/raphael-trzpit/FizzBuzz/internal/fizzbuzz"
	"log"
	"net/http"
	"os"
)

func main() {
	repo := fizzbuzz.NewMemoryStatisticsRepository()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/fizzbuzz", fizzbuzz.FizzBuzzHandler(repo))
	http.HandleFunc("/stats", fizzbuzz.StatsHandler(repo))

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
