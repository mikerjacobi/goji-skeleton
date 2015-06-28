package middleware

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"github.com/zenazn/goji/web"
)

var log = logging.MustGetLogger("app")

func PostgresMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("postgres", viper.GetString("db_connection"))
		if err != nil {
			log.Error("db err yo")
		}
		log.Notice(viper.GetString("db_connection"))
		c.Env["db"] = db
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
