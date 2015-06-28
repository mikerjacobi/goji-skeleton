package controllers

import (
	"net/http"

	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func HealthCheckHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	rend := c.Env["render"].(*render.Render)
	rend.JSON(w, http.StatusOK, "Success")
}

type IndexPayload struct{}

func IndexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	rend := c.Env["render"].(*render.Render)

	ip := IndexPayload{}
	/*if userInfo, ok := c.Env["userinfo"]; ok {
		sc.UserInfo = userInfo.(*models.UserInfo)
	}*/
	rend.HTML(w, http.StatusOK, "index", ip)
}
