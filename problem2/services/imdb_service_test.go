package services

import (
	"context"
	"errors"
	"fmt"
	"go-bit/entities/responses"
	"go-bit/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewIMDBService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	caller := mocks.NewMockIMDBCaller(ctrl)

	obj := NewIMDBService(caller)
	if obj == nil {
		t.Errorf("NewIMDBService return nil, expected return IMDBService")
	}
	if fmt.Sprintf("%T", obj) != "*services.imdbService" {
		t.Errorf("NewIMDBService return wrong type, expected '*services.imdbService'")
	}
	if _, ok := obj.(IMDBService); !ok {
		t.Errorf("NewIMDBService returns not implements IMDBService interface")
	}
}

func TestIMDBServiceSearch(t *testing.T) {
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
		in     input
		out    output
		caller *mocks.MockIMDBCaller
	}

	var (
		search = "Batman"
		page   = int32(1)

		errDefault = errors.New("error default")
	)

	testCases := []struct {
		name          string
		given         input
		expected      output
		configureMock func(*mockConfig)
	}{
		{
			name: "failed, caller.Search default error",
			given: input{
				ctx:    context.Background(),
				search: search,
				page:   page,
			},
			expected: output{err: errDefault},
			configureMock: func(conf *mockConfig) {

				conf.caller.EXPECT().
					Search(conf.in.ctx, conf.in.search, conf.in.page).
					Return(responses.IMDBSearchResponse{}, errDefault)
			},
		},
		{
			name: "success, search movies",
			given: input{
				ctx:    context.Background(),
				search: search,
				page:   page,
			},
			expected: output{err: nil},
			configureMock: func(conf *mockConfig) {

				conf.caller.EXPECT().
					Search(conf.in.ctx, conf.in.search, conf.in.page).
					Return(responses.IMDBSearchResponse{}, nil)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			caller := mocks.NewMockIMDBCaller(ctrl)

			test.configureMock(&mockConfig{
				in:     test.given,
				out:    test.expected,
				caller: caller,
			})

			s := imdbService{
				caller: caller,
			}

			respBody, err := s.Search(test.given.ctx, test.given.search, test.given.page)
			if expected := test.expected.err; (err == nil && expected != nil) || (err != nil && !errors.Is(err, expected)) {
				t.Errorf("error:\nEXPECTED: %v\nGOT: %v\n",
					expected, err)
			}
			if expected := test.expected.respBody; !reflect.DeepEqual(respBody, expected) {
				t.Errorf("respBody:\nEXPECTED:\n%+v\nGOT:\n%+v\n",
					expected, respBody)
			}
		})
	}
}

func TestIMDBServiceGetDetail(t *testing.T) {
	type input struct {
		ctx context.Context
		id  string
	}
	type output struct {
		respBody responses.IMDBGetDetailResponse
		err      error
	}
	type mockConfig struct {
		in     input
		out    output
		caller *mocks.MockIMDBCaller
	}

	var (
		id = "tt2011118"

		errDefault = errors.New("error default")
	)

	testCases := []struct {
		name          string
		given         input
		expected      output
		configureMock func(*mockConfig)
	}{
		{
			name: "failed, caller.GetDetail default error",
			given: input{
				ctx: context.Background(),
				id:  id,
			},
			expected: output{err: errDefault},
			configureMock: func(conf *mockConfig) {

				conf.caller.EXPECT().
					GetDetail(conf.in.ctx, conf.in.id).
					Return(responses.IMDBGetDetailResponse{}, errDefault)
			},
		},
		{
			name: "success, get detail movies",
			given: input{
				ctx: context.Background(),
				id:  id,
			},
			expected: output{err: nil},
			configureMock: func(conf *mockConfig) {

				conf.caller.EXPECT().
					GetDetail(conf.in.ctx, conf.in.id).
					Return(responses.IMDBGetDetailResponse{}, nil)
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			caller := mocks.NewMockIMDBCaller(ctrl)

			test.configureMock(&mockConfig{
				in:     test.given,
				out:    test.expected,
				caller: caller,
			})

			s := imdbService{
				caller: caller,
			}

			respBody, err := s.GetDetail(test.given.ctx, test.given.id)
			if expected := test.expected.err; (err == nil && expected != nil) || (err != nil && !errors.Is(err, expected)) {
				t.Errorf("error:\nEXPECTED: %v\nGOT: %v\n",
					expected, err)
			}
			if expected := test.expected.respBody; !reflect.DeepEqual(respBody, expected) {
				t.Errorf("respBody:\nEXPECTED:\n%+v\nGOT:\n%+v\n",
					expected, respBody)
			}
		})
	}
}
