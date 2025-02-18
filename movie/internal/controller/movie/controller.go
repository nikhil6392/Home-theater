package movie

import (
	"context"
	"errors"

	metadatamodel "movieexample.com/metadata/pkg/model"
	"movieexample.com/movie/internal/gateway"
	"movieexample.com/movie/pkg/model"
	ratingmodel "movieexample.com/rating/pkg/model"
	"movieexample.con/movie/internal/gateway"
)

// ErrNotFound is returned when the movie metadata is not
// found

var ErrNotFound = errors.New("movie metadata is not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)

	PutRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType, rating *ratingmodel.Rating) error
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

// Defining the service controller
type Controller struct {
	ratingGateway ratingGateway
	metadataGateway metadataGateway
}

// New Create a new movie service controller
func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway, metadataGateway}
}

// Get returns the movie details including the aggregated
// rating & movie metadata
// Get returns the movie details including the aggregated
// rating and movie metadata

func (c*Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeMovie)

	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		// Just proceed in this case , it's okk not to have rating yet
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}