package middleware

import "github.com/danecwalker/structured_go/router"

func HasUserCookieMiddleware(rh router.RouteHandler) router.RouteHandler {
	return func(rc *router.RouteContext) error {
		if _, err := rc.Request.Cookie("user"); err != nil {
			rc.Response.WriteHeader(401)
			rc.Response.Write([]byte("Unauthorized"))
			return nil
		}

		return rh(rc)
	}
}
