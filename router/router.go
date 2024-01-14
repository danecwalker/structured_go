package router

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
)

type RouteHandler func(*RouteContext) error
type Middleware func(RouteHandler) RouteHandler

type RouteContext struct {
	store    map[string]interface{}
	Params   map[string]string
	Response http.ResponseWriter
	Request  *http.Request

	logger *slog.Logger
}

func NewRouteContext(r *http.Request, w http.ResponseWriter, params map[string]string, logger *slog.Logger) *RouteContext {
	return &RouteContext{
		store:    make(map[string]interface{}),
		Params:   params,
		Request:  r,
		Response: w,
		logger:   logger,
	}
}

func (rc *RouteContext) Get(key string) interface{} {
	return rc.store[key]
}

func (rc *RouteContext) Set(key string, value interface{}) {
	rc.store[key] = value
}

func (rc *RouteContext) JSON(status int, body interface{}) error {
	json, err := json.Marshal(body)
	if err != nil {
		rc.logger.Error("an error occured while marshalling json", "error", err)
		return err
	}

	rc.Response.Header().Set("Content-Type", "application/json")
	rc.Response.WriteHeader(status)
	rc.Response.Write(json)
	return nil
}

func (rc *RouteContext) HTML(status int, template templ.Component) error {
	rc.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	rc.Response.WriteHeader(status)
	template.Render(context.Background(), rc.Response)
	return nil
}

func (rc *RouteContext) Text(status int, body string) error {
	rc.Response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	rc.Response.WriteHeader(status)
	rc.Response.Write([]byte(body))
	return nil
}
