package account

import (
	"github.com/danecwalker/structured_go/router"
)

func Get(rc *router.RouteContext) error {
	id := rc.Params["id"]
	return rc.JSON(200, map[string]any{
		"message": "Hello, World!",
		"id":      id,
	})
}
