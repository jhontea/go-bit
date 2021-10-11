package services

import (
	"context"
	"go-bit/callers"
	"go-bit/entities/responses"
)

type imdbService struct {
	caller callers.IMDBCaller
}

func NewIMDBService(caller callers.IMDBCaller) IMDBService {
	return &imdbService{
		caller: caller,
	}
}

func (s *imdbService) Search(ctx context.Context, search string, page int32) (responses.IMDBSearchResponse, error) {
	imdb, err := s.caller.Search(ctx, search, page)
	if err != nil {
		return imdb, err
	}

	return imdb, nil
}

func (s *imdbService) GetDetail(ctx context.Context, id string) (responses.IMDBGetDetailResponse, error) {
	imdb, err := s.caller.GetDetail(ctx, id)
	if err != nil {
		return imdb, err
	}

	return imdb, nil
}
