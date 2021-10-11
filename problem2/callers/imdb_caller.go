package callers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-bit/entities/responses"
	"go-bit/repositories"
	"net/http"
	"net/url"

	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
)

type imdbCaller struct {
	httpClient heimdall.Client
	baseURL    string
	apiKey     string
	repository repositories.LogCallerRepository
}

func NewIMDBCaller(repository repositories.LogCallerRepository) IMDBCaller {

	return &imdbCaller{
		httpClient: httpclient.NewClient(),
		baseURL:    "http://www.omdbapi.com/",
		apiKey:     "faf7e5bb&s",
		repository: repository,
	}
}

func (c *imdbCaller) Search(ctx context.Context, search string, page int32) (responses.IMDBSearchResponse, error) {
	var result responses.IMDBSearchResponse

	u, _ := url.Parse(c.baseURL)
	u.RawQuery = fmt.Sprintf("apikey=%s", c.apiKey)
	q := u.Query()
	q.Set("s", search)
	q.Set("page", fmt.Sprintf("%d", page))
	u.RawQuery = q.Encode()

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	response, err := c.httpClient.Get(u.String(), headers)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	go c.repository.StoreLogCaller(ctx, u.String(), response)

	json.NewDecoder(response.Body).Decode(&result)

	return result, nil
}

func (c *imdbCaller) GetDetail(ctx context.Context, id string) (responses.IMDBGetDetailResponse, error) {
	var result responses.IMDBGetDetailResponse

	u, _ := url.Parse(c.baseURL)
	u.RawQuery = fmt.Sprintf("apikey=%s", c.apiKey)
	q := u.Query()
	q.Set("i", id)
	u.RawQuery = q.Encode()

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	response, err := c.httpClient.Get(u.String(), headers)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	go c.repository.StoreLogCaller(ctx, u.String(), response)

	json.NewDecoder(response.Body).Decode(&result)

	if response.StatusCode != 200 || result.Error != "" {
		return result, errors.New(result.Error)
	}

	return result, nil
}
