package api

import (
	"net/http"
	"strings"
)

// Collect endpoint to handle tracking client data
// this is a GET endpoint that embeds an image into the client browser and
// records request data according to the siteID and path provided
// the path after siteID is a way to track specific URLs as that is not possible
// by embeding images alone
// e.x /collect/{siteID}/any/url/path
func (a *API) Collect(w http.ResponseWriter, r *http.Request) {
	url := strings.Trim(r.URL.Path, "/")
	path := strings.Split(url, "/")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strings.Join(path, "/")))
}
