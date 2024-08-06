package main

import (
	"log"
	"net/http"

	"github.com/dangquyitt/go-movie/metadata/internal/business/metadata"
	httphandler "github.com/dangquyitt/go-movie/metadata/internal/handler/http"
	"github.com/dangquyitt/go-movie/metadata/internal/repository/memory"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.New()
	business := metadata.New(repo)
	h := httphandler.New(business)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
