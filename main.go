package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/danecwalker/structured_go/router"
	"github.com/julienschmidt/httprouter"
)

type App struct {
	Router *httprouter.Router
	Logger *slog.Logger
}

func (app *App) RegisterRoute(method, path string, handler router.RouteHandler, middleware ...router.Middleware) {
	app.Router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, _params httprouter.Params) {
		var params map[string]string
		if _params != nil {
			params = make(map[string]string)
			for _, param := range _params {
				params[param.Key] = param.Value
			}
		}
		ctx := router.NewRouteContext(r, w, params, app.Logger)

		for _, m := range middleware {
			handler = m(handler)
		}

		if err := handler(ctx); err != nil {
			ctx.Response.WriteHeader(500)
			ctx.Response.Write([]byte(err.Error()))
			app.Logger.Error("an error occured while handling the request", "error", err)
		}
	})
}

func main() {
	app := App{
		Router: httprouter.New(),
		Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	RegisterRoutes(&app)

	app.Logger.Info("starting the server")
	if err := http.ListenAndServe(":3000", app.Router); err != nil {
		app.Logger.Error("an error occured while starting the server", "error", err)
	}
}
