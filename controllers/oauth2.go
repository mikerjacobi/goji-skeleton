package controllers

import (
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
	"github.com/philpearl/tt_goji_middleware/base"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
	"golang.org/x/oauth2"
)

func getOAuthConf() oauth2.Config {
	base_url := viper.GetString("oauth_server_base_url")
	client_id := viper.GetString("client_id")
	client_secret := viper.GetString("client_secret")

	conf := oauth2.Config{
		ClientID:     client_id,
		ClientSecret: client_secret,
		Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  base_url + "/authorize",
			TokenURL: base_url + "/token",
		},
	}
	return conf
}

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Infof("%v\n", r.Form)
	conf := getOAuthConf()
	authURL := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	log.Infof("redirecting to: %s", authURL)
	http.Redirect(w, r, authURL, http.StatusFound)
}

func OAuth2Callback(c web.C, w http.ResponseWriter, r *http.Request) {
	rend := c.Env["render"].(*render.Render)

	u, err := url.Parse(r.RequestURI)
	if err != nil {
		log.Error("failed to parse RequestURI:", err)
		rend.JSON(w, http.StatusInternalServerError, nil)
		return
	}
	q := u.Query()

	code := q.Get("code")
	conf := getOAuthConf()

	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Error("failed to exchange for token:", err)
		rend.JSON(w, http.StatusInternalServerError, nil)
		return
	}
	//session.Put("token", tok)
	//session.Put("google_id", info.GoogleID)
	log.Infof("successfully exchanged token: %s.  redirecting to home", token)
	http.Redirect(w, r, "/home", http.StatusFound)

	/*
		rend := c.Env["render"].(*render.Render)
		session := c.Env["session"].(*base.Session)
		db := c.Env["db"].(*mgo.Database)

		u, err := url.Parse(r.RequestURI)
		if err != nil {
			log.Error("[ERROR] failed to parse RequestURI:", err)
			rend.JSON(w, http.StatusInternalServerError, nil)
			return
		}
		q := u.Query()

		code := q.Get("code")
		//tok, err := oauthconf.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Error("[ERROR] failed to exchange for token:", err)
			rend.JSON(w, http.StatusInternalServerError, nil)
			return
		}
			client := oauthconf.Client(oauth2.NoContext, tok)

			//get their email
			resp, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
			if err != nil {
				log.Error("[ERROR] failed to get user info:", err)
				rend.JSON(w, http.StatusInternalServerError, nil)
				return
			}
			defer resp.Body.Close()


			var info models.UserInfo
			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&info); err != nil {
				log.Error("[ERROR] failed to decode json:", err)
				rend.JSON(w, http.StatusInternalServerError, nil)
				return
			}

			if err := info.Upsert(db); err != nil {
				log.Error("[ERROR] failed to upsert user:", err)
				rend.JSON(w, http.StatusInternalServerError, nil)
				return
			}

			session.Put("token", tok)
			session.Put("google_id", info.GoogleID)
			http.Redirect(w, r, "/", http.StatusFound)
	*/
}

func Logout(c web.C, w http.ResponseWriter, r *http.Request) {
	session := c.Env["session"].(*base.Session)
	session.Del("token")
	session.Del("google_id")

	//set a cookie with a negative maxage to delete it
	sessionid_cookie := http.Cookie{
		Name:   "sessionid",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, &sessionid_cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}
