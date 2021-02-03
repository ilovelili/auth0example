package user

import (
	"net/http"

	"github.com/ilovelili/auth0example/app"
	templates "github.com/ilovelili/auth0example/routes"
)

// Handler user handler
func Handler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "user", struct {
		Profile interface{}
		IDToken interface{}
	}{
		Profile: session.Values["profile"],
		IDToken: session.Values["id_token"],
	})
}
