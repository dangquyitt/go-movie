package main

import (
	"log"
	"net/http"

	"github.com/dangquyitt/go-movie/movie/internal/business/movie"
	metadatagateway "github.com/dangquyitt/go-movie/movie/internal/gateway/metadata/http"
	ratinggateway "github.com/dangquyitt/go-movie/movie/internal/gateway/rating/http"
	httphandler "github.com/dangquyitt/go-movie/movie/internal/handler/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	business := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(business)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
