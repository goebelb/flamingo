package responder

import (
	"flamingo/core/flamingo"
	"flamingo/core/flamingo/web"
	"net/http"
)

// RedirectAware allows a controller to issue a 302 redirect
type RedirectAware struct {
	Router *flamingo.Router `inject:""`
}

// Redirect returns a web.RedirectResponse with the proper URL
func (r *RedirectAware) Redirect(name string, args ...string) web.Response {
	url := r.Router.Url(name, args...)

	return web.RedirectResponse{
		Status:   http.StatusFound,
		Location: url.String(),
	}
}
