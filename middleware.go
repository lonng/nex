package nex

import (
	"context"
	"net/http"
)

type BeforeFunc func(context.Context, *http.Request) (context.Context, error)
type AfterFunc func(context.Context, http.ResponseWriter) (context.Context, error)
