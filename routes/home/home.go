package home

import (
	"net/http"

	templates "github.com/ilovelili/auth0example/routes"
)

// Handler home handler
func Handler(w http.ResponseWriter, r *http.Request) {
	templates.RenderTemplate(w, "home", nil)
}
