package controllers

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func HealthCheckHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	rend := c.Env["render"].(*render.Render)
	wd, _ := os.Getwd()
	logrus.Info(">>>", wd)
	logrus.Info(viper.GetString("template_path"))
	rend.JSON(w, http.StatusOK, "Success")
}

type IndexPayload struct{}

func IndexHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	rend := c.Env["render"].(*render.Render)

	//ip := IndexPayload{}
	/*if userInfo, ok := c.Env["userinfo"]; ok {
		sc.UserInfo = userInfo.(*models.UserInfo)
	}*/
	//tmpl := viper.GetString("template_path") + "/index"
	rend.HTML(w, http.StatusOK, "index", nil)
	//rend.JSON(w, http.StatusOK, "oo1o")
}
