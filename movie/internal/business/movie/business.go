package movie

import (
	"context"
	"errors"

	metadatamodel "github.com/dangquyitt/go-movie/metadata/pkg"
	"github.com/dangquyitt/go-movie/movie/internal/gateway"
	model "github.com/dangquyitt/go-movie/movie/pkg"
	ratingmodel "github.com/dangquyitt/go-movie/rating/pkg"
)

// ErrNotFound is returned when the movie metadata is not
// found.
var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}
type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

// Business defines a movie service business.
type Business struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New creates a new movie service business.
func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Business {
	return &Business{ratingGateway, metadataGateway}
}

// Get returns the movie details including the aggregated
// rating and movie metadata.
// Get returns the movie details including the aggregated rating and movie metadata.
func (c *Business) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Just proceed in this case, it's ok not to have ratings yet.
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
