package repositories

import (
	"context"
	"net/http"
)

type LogCallerRepository interface {
	StoreLogCaller(ctx context.Context, url string, response *http.Response)
}
