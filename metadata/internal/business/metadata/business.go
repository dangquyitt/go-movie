package metadata

import (
	"context"
	"errors"
	"github.com/dangquyitt/go-movie/metadata/pkg/model"

	"github.com/dangquyitt/go-movie/metadata/internal/repository"
)

// ErrNotFound is returned when a requested record is not // found.
var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Business defines a metadata service Business.
type Business struct {
	repo metadataRepository
}

// New creates a metadata service business.
func New(repo metadataRepository) *Business {
	return &Business{repo}
}

// Get returns movie metadata by id.
func (c *Business) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
