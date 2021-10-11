package services

import (
	"context"
	"go-bit/entities/responses"
)

type IMDBService interface {
	Search(ctx context.Context, search string, page int32) (responses.IMDBSearchResponse, error)
	GetDetail(ctx context.Context, id string) (responses.IMDBGetDetailResponse, error)
}
