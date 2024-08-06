package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dangquyitt/go-movie/metadata/internal/business/metadata"
	"github.com/dangquyitt/go-movie/metadata/internal/repository"
)

// Handler defines a movie metadata HTTP handler.
type Handler struct {
	business *metadata.Business
}

// New creates a new movie metadata HTTP handler.
func New(business *metadata.Business) *Handler {
	return &Handler{business}
}

// GetMetadata handles GET /metadata requests.
func (h *Handler) GetMetadata(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := req.Context()
	m, err := h.business.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
