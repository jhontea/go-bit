package callers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-bit/entities/responses"
	"go-bit/mocks"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewIMDBCaller(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockLogCallerRepository(ctrl)

	imdbCallerObj := NewIMDBCaller(repository)
	if imdbCallerObj == nil {
		t.Errorf("NewIMDBCaller return nil, expected return IMDBCaller")
	}
	if fmt.Sprintf("%T", imdbCallerObj) != "*callers.imdbCaller" {
		t.Errorf("NewIMDBCaller return wrong type, expected '*callers.imdbCaller'")
	}
	if _, ok := imdbCallerObj.(IMDBCaller); !ok {
		t.Errorf("NewIMDBCaller returns not implements IMDBCaller interface")
	}
}

func TestIMDBCallerSearch(t *testing.T) {
	type input struct {
		ctx    context.Context
		search string
		page   int32
	}
	type output struct {
		respBody responses.IMDBSearchResponse
		err      error
	}
	type mockConfig struct {
		in             input
		out            output
		httpClientMock *mocks.MockClient
		repository     *mocks.MockLogCallerRepository
		wg             *sync.WaitGroup
	}

	var (
		baseURL = "http://www.omdbapi.com/"
		apikey  = "apikey"
		search  = "Batman"
		page    = int32(1)

		errDefault = errors.New("error default")
	)

	testCases := []struct {
		name          string
		given         input
		expected      output
		configureMock func(*mockConfig)
	}{
		{
			name: "failed, httpClient.Get default error",
			given: input{
				ctx:    context.Background(),
				search: search,
				page:   page,
			},
			expected: output{err: errDefault},
			configureMock: func(conf *mockConfig) {
				url := fmt.Sprintf("%s?apikey=%s&page=%d&s=%s", baseURL, apikey, page, search)

				conf.httpClientMock.EXPECT().
					Get(url, gomock.Any()).
					Return(&http.Response{}, errDefault)
			},
		},
		{
			name: "success, search imdb movies",
			given: input{
				ctx:    context.Background(),
				search: search,
				page:   page,
			},
			expected: output{
				respBody: responses.IMDBSearchResponse{
					Response: "success",
				},
				err: nil,
			},
			configureMock: func(conf *mockConfig) {
				res := &http.Response{
					Status:     "200 OK",
					StatusCode: http.StatusOK,
				}
				resBodyJSON, _ := json.Marshal(conf.out.respBody)
				res.Body = ioutil.NopCloser(strings.NewReader(string(resBodyJSON)))

				url := fmt.Sprintf("%s?apikey=%s&page=%d&s=%s", baseURL, apikey, page, search)

				conf.httpClientMock.EXPECT().
					Get(url, gomock.Any()).
					Return(res, nil)

				conf.wg.Add(1)
				conf.repository.EXPECT().
					StoreLogCaller(gomock.Any(), url, gomock.Any()).
					Do(func(interface{}, interface{}, interface{}) {
						conf.wg.Done()
					})
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			var wg sync.WaitGroup

			httpClientMock := mocks.NewMockClient(ctrl)
			repository := mocks.NewMockLogCallerRepository(ctrl)

			test.configureMock(&mockConfig{
				in:             test.given,
				out:            test.expected,
				httpClientMock: httpClientMock,
				repository:     repository,
				wg:             &wg,
			})

			c := imdbCaller{
				baseURL:    baseURL,
				apiKey:     apikey,
				httpClient: httpClientMock,
				repository: repository,
			}

			respBody, err := c.Search(test.given.ctx, test.given.search, test.given.page)
			if expected := test.expected.err; (err == nil && expected != nil) || (err != nil && !errors.Is(err, expected)) {
				t.Errorf("error:\nEXPECTED: %v\nGOT: %v\n",
					expected, err)
			}

			wg.Wait()

			if expected := test.expected.respBody; !reflect.DeepEqual(respBody, expected) {
				t.Errorf("respBody:\nEXPECTED:\n%+v\nGOT:\n%+v\n",
					expected, respBody)
			}
		})
	}
}

func TestIMDBCallerGetDetail(t *testing.T) {
	type input struct {
		ctx context.Context
		id  string
	}
	type output struct {
		respBody responses.IMDBGetDetailResponse
		err      error
	}
	type mockConfig struct {
		in             input
		out            output
		httpClientMock *mocks.MockClient
		repository     *mocks.MockLogCallerRepository
		wg             *sync.WaitGroup
	}

	var (
		baseURL = "http://www.omdbapi.com/"
		apikey  = "apikey"
		id      = "tt2011118"

		errDefault = errors.New("error default")
	)

	testCases := []struct {
		name          string
		given         input
		expected      output
		configureMock func(*mockConfig)
	}{
		{
			name: "failed, httpClient.Get default error",
			given: input{
				ctx: context.Background(),
				id:  id,
			},
			expected: output{err: errDefault},
			configureMock: func(conf *mockConfig) {
				url := fmt.Sprintf("%s?apikey=%s&i=%s", baseURL, apikey, id)

				conf.httpClientMock.EXPECT().
					Get(url, gomock.Any()).
					Return(&http.Response{}, errDefault)
			},
		},
		{
			name: "success, search imdb movies",
			given: input{
				ctx: context.Background(),
				id:  id,
			},
			expected: output{
				respBody: responses.IMDBGetDetailResponse{
					Response: "success",
				},
				err: nil,
			},
			configureMock: func(conf *mockConfig) {
				res := &http.Response{
					Status:     "200 OK",
					StatusCode: http.StatusOK,
				}
				resBodyJSON, _ := json.Marshal(conf.out.respBody)
				res.Body = ioutil.NopCloser(strings.NewReader(string(resBodyJSON)))

				url := fmt.Sprintf("%s?apikey=%s&i=%s", baseURL, apikey, id)

				conf.httpClientMock.EXPECT().
					Get(url, gomock.Any()).
					Return(res, nil)

				conf.wg.Add(1)
				conf.repository.EXPECT().
					StoreLogCaller(gomock.Any(), url, gomock.Any()).
					Do(func(interface{}, interface{}, interface{}) {
						conf.wg.Done()
					})
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			var wg sync.WaitGroup

			httpClientMock := mocks.NewMockClient(ctrl)
			repository := mocks.NewMockLogCallerRepository(ctrl)

			test.configureMock(&mockConfig{
				in:             test.given,
				out:            test.expected,
				httpClientMock: httpClientMock,
				repository:     repository,
				wg:             &wg,
			})

			c := imdbCaller{
				baseURL:    baseURL,
				apiKey:     apikey,
				httpClient: httpClientMock,
				repository: repository,
			}

			respBody, err := c.GetDetail(test.given.ctx, test.given.id)
			if expected := test.expected.err; (err == nil && expected != nil) || (err != nil && !errors.Is(err, expected)) {
				t.Errorf("error:\nEXPECTED: %v\nGOT: %v\n",
					expected, err)
			}

			wg.Wait()

			if expected := test.expected.respBody; !reflect.DeepEqual(respBody, expected) {
				t.Errorf("respBody:\nEXPECTED:\n%+v\nGOT:\n%+v\n",
					expected, respBody)
			}
		})
	}
}
