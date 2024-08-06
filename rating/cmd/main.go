package main

import (
	"log"
	"net/http"

	"github.com/dangquyitt/go-movie/rating/internal/business/rating"
	httphandler "github.com/dangquyitt/go-movie/rating/internal/handler/http"
	"github.com/dangquyitt/go-movie/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	business := rating.New(repo)
	h := httphandler.New(business)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
