package controllers

import (
	"encoding/gob"
	"encoding/json"
	"net/http"
	"net/url"

	"bitbucket.org/lothsoft/social-wallpaper/models"

	log "github.com/Sirupsen/logrus"
	"github.com/philpearl/tt_goji_middleware/base"
	"github.com/spf13/viper"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2"
)

var oauthconf *oauth2.Config

func init() {
	gob.Register(&oauth2.Token{})
}

func SetupAuth(conf *oauth2.Config) {
	oauthconf = conf
}

func Login(c web.C, w http.ResponseWriter, r *http.Request) {
	//url := oauthconf.AuthCodeURL("state")
	base_url := viper.GetString("oauth_server_base_url")
	client_id := viper.GetString("client_id")
	_ = viper.GetString("client_secret")

	u := url.URL{}
	u.Scheme = "http"
	u.Host = base_url
	u.Path = "/authorize"
	q := u.Query()
	q.Set("response_type", "code")
	q.Set("client_id", client_id)
	q.Set("redirect_url", "http://localhost:14000/appauth/code")
	u.RawQuery = q.Encode()

	log.Infof("redirecting to: %s", u.String())

	http.Redirect(w, r, u.String(), http.StatusFound)
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

func OAuth2Callback(c web.C, w http.ResponseWriter, r *http.Request) {
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
	tok, err := oauthconf.Exchange(oauth2.NoContext, code)
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
}
