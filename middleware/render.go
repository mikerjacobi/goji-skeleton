package middleware

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

var rend *render.Render

func init() {
	logrus.Info("starting render middleware init")
	rend = render.New(render.Options{
		Layout: "layout",
		//Directory: "/go/src/templates",
		Directory: viper.GetString("template_path"),
	})
}

func RenderMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["render"] = rend
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
