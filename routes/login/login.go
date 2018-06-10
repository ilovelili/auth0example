package login

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/ilovelili/auth0example/app"
	"golang.org/x/oauth2"
)

// Handler login handler
func Handler(w http.ResponseWriter, r *http.Request) {
	domain := os.Getenv("AUTH0_DOMAIN")
	aud := os.Getenv("AUTH0_AUDIENCE")

	conf := &oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://" + domain + "/authorize",
			TokenURL: "https://" + domain + "/oauth/token",
		},
	}

	if aud == "" {
		aud = "https://" + domain + "/userinfo"
	}

	// Generate random state
	b := make([]byte, 32)
	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)

	session, err := app.Store.Get(r, "state")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	audience := oauth2.SetAuthURLParam("audience", aud)
	url := conf.AuthCodeURL(state, audience)

	fmt.Printf("url is %v\n", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
