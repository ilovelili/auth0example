package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/ilovelili/auth0example/routes/callback"
	"github.com/ilovelili/auth0example/routes/home"
	"github.com/ilovelili/auth0example/routes/login"
	"github.com/ilovelili/auth0example/routes/logout"
	"github.com/ilovelili/auth0example/routes/middlewares"
	"github.com/ilovelili/auth0example/routes/user"
)

var (
	port = flag.String("port", ":3000", "port")
)

// Serve serve
func Serve() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", home.Handler)
	r.HandleFunc("/login", login.Handler)
	r.HandleFunc("/logout", logout.Handler)
	r.HandleFunc("/callback", callback.Handler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.Handler)),
	))

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(*port, nil))
}
