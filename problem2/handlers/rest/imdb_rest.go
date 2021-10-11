package rest

import (
	"context"
	"go-bit/entities/responses"
	"go-bit/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type IMDBRestHandler struct {
	service services.IMDBService
}

func NewIMDBRestHandler(service services.IMDBService) *IMDBRestHandler {
	return &IMDBRestHandler{
		service: service,
	}
}

func (h *IMDBRestHandler) SearchIMDBHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		search := strings.TrimSpace(r.URL.Query().Get("search"))
		if search == "" {
			responses.FailedResponse(w, nil, "search must not empty")
			return
		}

		page := strings.TrimSpace(r.URL.Query().Get("page"))
		if page == "" {
			page = "1"
		}
		pageConv, _ := strconv.ParseInt(page, 10, 32)

		result, err := h.service.Search(context.TODO(), search, int32(pageConv))
		if err != nil {
			responses.FailedResponse(w, nil, "failed search movie")
			return
		}

		responses.SuccessResponse(w, result, "success search movie")
	}
}

func (h *IMDBRestHandler) GetIMDBDetailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id == "" {
			responses.FailedResponse(w, nil, "id movie must not empty")
			return
		}

		result, err := h.service.GetDetail(context.TODO(), id)
		if err != nil {
			responses.FailedResponse(w, nil, err.Error())
			return
		}

		responses.SuccessResponse(w, result, "success get detail movie")
	}

}
