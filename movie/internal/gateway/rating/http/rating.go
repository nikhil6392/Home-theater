package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"movieexample.com/movie/internal/gateway"
	"movieexample.com/rating/pkg/model"
)

// Gateway defines an HTTP gateway for a string service.
type Gateway struct {
	addr string
}

// New creates a new HTTP gateway for a rating service.
func New(addr string) *Gateway {
	return &Gateway{addr}
}

// GetAggregatedRating returns the aggregated rating for a 
// record or ErrNotFound if there are no ratings for it.
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, g.addr+"/rating", nil) 
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", string(recordID))
	values.Add("type", fmt.Sprintf("%v", recordType))
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return 0, gateway.ErrNotFound
	} else if resp.StatusCode/ 100 != 2 {
		return 0, fmt.Errorf("npn- 2xx response: %v", resp)
	}
	var v float64
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return 0, err
	}
	return v, nil
}