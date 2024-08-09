package grpc

import (
	"context"
	"errors"

	"github.com/dangquyitt/go-movie/gen"
	"github.com/dangquyitt/go-movie/metadata/internal/business/metadata"
	model "github.com/dangquyitt/go-movie/metadata/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines a movie metadata gRPC handler.
type Handler struct {
	gen.UnimplementedMetadataServiceServer
	biz *metadata.Business
}

// New creates a new movie metadata gRPC handler.
func New(biz *metadata.Business) *Handler {
	return &Handler{
		biz: biz,
	}
}

// GetMetadataByID returns movie metadata by id.
func (h *Handler) GetMetadata(ctx context.Context, req *gen.GetMetadataRequest) (*gen.GetMetadataResponse, error) {
	if req == nil || req.MovieId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.biz.Get(ctx, req.MovieId)
	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetMetadataResponse{Metadata: model.MetadataToProto(m)}, nil
}
