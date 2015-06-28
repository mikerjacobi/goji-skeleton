package middleware

import "strings"

var (
	auth_exempt_whitelist = [...]string{
		"/generate_rss",
		"/login",
		"/logout",
		"/health-check",
		"/oauth2callback",
	}
)

func endpointIsExemptFromAuth(endpoint string) bool {
	for i := 0; i < len(auth_exempt_whitelist); i++ {
		if strings.Contains(endpoint, auth_exempt_whitelist[i]) {
			return true
		}
	}

	return false
}

/*
func AuthMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session := c.Env["session"].(*base.Session)
		if endpointIsExemptFromAuth(string(r.URL.Path)) {
			h.ServeHTTP(w, r)
			return
		}

		if googleID, ok := session.Get("google_id"); ok {
			id := googleID.(string)
			db := c.Env["db"].(*mgo.Database)
			userInfo, err := models.GetUserInfo(db, id)
			if err == nil {
				c.Env["userinfo"] = userInfo
			}
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		h.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}*/
