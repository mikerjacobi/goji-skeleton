package main

import (
	"flag"

	"github.com/hypebeast/gojistatic"
	"github.com/mikerjacobi/goji-skeleton/middleware"
	_ "github.com/mikerjacobi/goji-skeleton/zzz_config"

	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
	"github.com/mikerjacobi/goji-skeleton/controllers"
	"github.com/zenazn/goji"
)

func init() {
	log.Info("starting main middleware init")
}

/*
func alias2ipaddr(alias string) (string, error) {
	cmd := exec.Command("grep", alias, "/etc/hosts")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.Split(out.String(), " ")[0], nil
}
*/

func main() {

	//setup sessions
	/*
		sh := redis.NewSessionHolder()
		redisconfig := viper.GetStringMapString("redis")
		if _, ok := redisconfig["host"]; !ok {
			panic("failed to read redis host")
		}
		if _, ok := redisconfig["port"]; !ok {
			panic("failed to read redis port")
		}

		redisip, err := alias2ipaddr(redisconfig["host"])
		if err != nil {
			panic("failed to lookup redis IP address")
		}
		goji.Use(redis.BuildRedis(fmt.Sprintf("%s:%s", redisip, redisconfig["port"])))
		goji.Use(base.BuildSessionMiddleware(sh))
	*/

	//setup render middleware
	goji.Use(middleware.RenderMiddleware)

	//setup database
	/*
		dbconfig := viper.GetStringMapString("db")
		if _, ok := dbconfig["host"]; !ok {
			panic("failed to read db host")
		}
		if _, ok := dbconfig["name"]; !ok {
			panic("failed to read db name")
		}
		goji.Use(middleware.PostgresMiddleware)

		goji.Use(middleware.AuthMiddleware)
	*/

	//setup routes
	goji.Get("/home", controllers.IndexHandler)
	goji.Get("/healthcheck", controllers.HealthCheckHandler)
	goji.Get("/login", controllers.Login)
	goji.Get("/oauth2callback", controllers.OAuth2Callback)
	//goji.Get("/logout", controllers.Logout)

	//setup static assets
	goji.Use(gojistatic.Static("/go/src/static", gojistatic.StaticOptions{SkipLogging: false, Prefix: "static"}))
	//goji.Use(gojistatic.Static("node_modules", gojistatic.StaticOptions{SkipLogging: false, Prefix: "node_modules"}))

	//begin
	log.Info("Starting App...")

	flag.Set("bind", ":80")
	goji.Serve()
}
