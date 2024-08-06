package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dangquyitt/go-movie/movie/internal/business/movie"
)

// Handler defines a movie handler.
type Handler struct {
	business *movie.Business
}

// New creates a new movie HTTP handler.
func New(business *movie.Business) *Handler {
	return &Handler{business}
}

// GetMovieDetails handles GET /movie requests.
func (h *Handler) GetMovieDetails(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	details, err := h.business.Get(req.Context(), id)
	if err != nil && errors.Is(err, movie.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
