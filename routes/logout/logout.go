package logout

import (
	"net/http"
	"net/url"
	"os"
)

// Handler logout handler
func Handler(w http.ResponseWriter, r *http.Request) {
	domain := os.Getenv("AUTH0_DOMAIN")

	var logouturl *url.URL
	logouturl, err := url.Parse("https://" + domain)

	if err != nil {
		panic("boom")
	}

	logouturl.Path += "/v2/logout"
	parameters := url.Values{}
	parameters.Add("returnTo", os.Getenv("AUTH0_BASE_URL"))
	parameters.Add("client_id", "gI6Q1OqTtASpXhIt7SG5DHVAe8RyJgfP")
	logouturl.RawQuery = parameters.Encode()

	http.Redirect(w, r, logouturl.String(), http.StatusTemporaryRedirect)
}
