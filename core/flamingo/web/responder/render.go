package responder

import (
	"flamingo/core/flamingo"
	"flamingo/core/flamingo/web"
	"flamingo/core/packages/pug-template"
	"net/http"
)

// RenderAware allows pug-template rendering
type RenderAware struct {
	Router *flamingo.Router `inject:""`
}

// Render returns a web.ContentResponse with status 200 and ContentType text/html
func (r *RenderAware) Render(context web.Context, tpl string, data interface{}) web.Response {
	return web.ContentResponse{
		Status:      http.StatusOK,
		Body:        template.Render(r.Router, context, tpl, data),
		ContentType: "text/html",
	}
}
