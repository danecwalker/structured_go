package main

import (
	"github.com/danecwalker/structured_go/web/api/account"
	"github.com/danecwalker/structured_go/web/middleware"
	"github.com/danecwalker/structured_go/web/pages/about"
	"github.com/danecwalker/structured_go/web/pages/home"
)

func RegisterRoutes(app *App) {
	app.RegisterRoute("GET", "/", home.Get)
	app.RegisterRoute("GET", "/about", about.Get, middleware.HasUserCookieMiddleware)

	app.RegisterRoute("GET", "/api/v1/account/:id", account.Get)
}
