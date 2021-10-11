package grpc

import (
	"context"
	"log"

	"go-bit/entities/proto"
	"go-bit/services"
)

type IMDBGrpcHandler struct {
	service services.IMDBService
}

func NewIMDBGrpcHandler(service services.IMDBService) *IMDBGrpcHandler {
	return &IMDBGrpcHandler{
		service: service,
	}
}

func (h *IMDBGrpcHandler) Search(ctx context.Context, request *proto.SearchRequest) (*proto.SearchResponse, error) {
	log.Printf("have a request with search %s and page %d", request.Search, request.Page)

	if request.Page <= 0 {
		request.Page = 1
	}

	result, err := h.service.Search(ctx, request.Search, request.Page)
	if err != nil {
		return &proto.SearchResponse{}, err
	}

	var imdbSearchData []*proto.IMDBSearchData
	for _, v := range result.Search {
		searchData := proto.IMDBSearchData{
			Title:  v.Title,
			Year:   v.Year,
			ImdbID: v.ImdbID,
			Type:   v.Type,
			Poster: v.Poster,
		}

		imdbSearchData = append(imdbSearchData, &searchData)
	}

	return &proto.SearchResponse{
		List:         imdbSearchData,
		TotalResults: result.TotalResults,
		Response:     result.Response,
	}, nil
}

func (h *IMDBGrpcHandler) GetDetail(ctx context.Context, request *proto.GetDetailRequest) (*proto.GetDetailResponse, error) {
	log.Printf("have a request with id %s", request.Id)

	result, err := h.service.GetDetail(ctx, request.Id)
	if err != nil {
		return &proto.GetDetailResponse{}, err
	}

	var imdbDetailRating []*proto.IMDBGetDetailRating
	for _, v := range result.Ratings {
		sdetailRating := proto.IMDBGetDetailRating{
			Source: v.Source,
			Value:  v.Value,
		}

		imdbDetailRating = append(imdbDetailRating, &sdetailRating)
	}

	return &proto.GetDetailResponse{
		Id:         result.ImdbID,
		Title:      result.Title,
		Year:       result.Year,
		Rated:      result.Rated,
		Released:   result.Released,
		Runtime:    result.Runtime,
		Genre:      result.Genre,
		Director:   result.Director,
		Writer:     result.Writer,
		Actors:     result.Actors,
		Plot:       result.Plot,
		Language:   result.Language,
		Country:    result.Country,
		Awards:     result.Awards,
		Poster:     result.Poster,
		List:       imdbDetailRating,
		Metascore:  result.Metascore,
		ImdbRating: result.ImdbRating,
		ImdbVotes:  result.ImdbVotes,
		ImdbID:     result.ImdbID,
		Type:       result.Type,
		Dvd:        result.Dvd,
		BoxOffice:  result.BoxOffice,
		Production: result.Production,
		Website:    result.Website,
		Response:   result.Response,
		Error:      result.Error,
	}, nil
}
