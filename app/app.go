package app

import (
	"encoding/gob"

	"github.com/gorilla/sessions"
)

var (
	// Store cookie store
	Store *sessions.CookieStore
)

// Init cookiestore
func Init() {
	Store = sessions.NewCookieStore([]byte("auth0"))
	// regiser a key value session storage gob
	/*
		callback.go

		var profile map[string]interface{}
		if err = json.NewDecoder(resp.Body).Decode(&profile); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
	gob.Register(map[string]interface{}{})
}
