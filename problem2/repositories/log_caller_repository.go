package repositories

import (
	"context"
	"net/http"
	"time"
)

type inMemmoryLogCaller struct {
	URL       string         `json:"url"`
	Response  *http.Response `json:"response"`
	CreatedAt time.Time      `json:"created_at"`
}

type logCallerRepository struct {
	storage []inMemmoryLogCaller
}

func NewLogCallerRepository() LogCallerRepository {
	return &logCallerRepository{
		storage: make([]inMemmoryLogCaller, 0),
	}
}

func (r *logCallerRepository) StoreLogCaller(ctx context.Context, url string, response *http.Response) {
	r.storage = append(r.storage, inMemmoryLogCaller{
		URL:       url,
		Response:  response,
		CreatedAt: time.Now(),
	})
}
