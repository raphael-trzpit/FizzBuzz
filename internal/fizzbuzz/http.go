package fizzbuzz

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func FizzBuzzHandler(statsRepo StatisticsRepository) http.HandlerFunc {
	type fizzBuzzRequest struct {
		FizzMultiple int `json:"int1"`
		BuzzMultiple int `json:"int2"`
		Limit int `json:"limit"`
		FizzStr string `json:"str1"`
		BuzzStr string `json:"str2"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var request fizzBuzzRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			// @TODO this could be prettier :)
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fizzBuzzList := FizzBuzz(request.FizzMultiple, request.BuzzMultiple, request.Limit, request.FizzStr, request.BuzzStr)

		err = statsRepo.store(Hit{request.FizzMultiple, request.BuzzMultiple, request.Limit, request.FizzStr, request.BuzzStr})
		if err != nil {
			log.Print(err)
		}

		fmt.Fprintf(w, strings.Join(fizzBuzzList, ","))
	}
}

func StatsHander(statsRepo StatisticsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			FizzMultiple int `json:"int1"`
			BuzzMultiple int `json:"int2"`
			Limit int `json:"limit"`
			FizzStr string `json:"str1"`
			BuzzStr string `json:"str2"`
			Count int `json:"count"`
		}

		hit, count, err := statsRepo.getMostUsedWithCount()
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(response{hit.fizzMultiple, hit.buzzMultiple, hit.limit, hit.fizzStr, hit.buzzStr, count})
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}