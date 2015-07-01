package controllers

import (
	"net/http"

	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func HealthCheckHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	rend := c.Env["render"].(*render.Render)
	resp := map[string]string{"Success": "true"}
	rend.JSON(w, http.StatusOK, resp)
}

type IndexPayload struct {
	Action *string
}

func IndexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	rend := c.Env["render"].(*render.Render)

	action := "http://www.jacobra.com:8003/login"
	ip := IndexPayload{
		Action: &action,
	}
	/*if userInfo, ok := c.Env["userinfo"]; ok {
		sc.UserInfo = userInfo.(*models.UserInfo)
	}*/
	rend.HTML(w, http.StatusOK, "index", ip)
}
