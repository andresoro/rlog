package api

import (
	"net/http"
	"strings"
	"time"
)

// Pixel GIF 1x1 1x1+0+0 8-bit sRGB
const Pixel = "\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\xFF\x00\xFF\xFF\xFF\x00\x00\x00\x2C\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3B"

// Collect endpoint to handle tracking client data
// this is a GET endpoint that embeds an image into the client browser and
// records request data according to the siteID and path provided
// the path after siteID is a way to track specific URLs as that is not possible
// by embeding images alone
// e.x /collect/{siteID}/any/url/path
func (a *API) Collect(w http.ResponseWriter, r *http.Request) {
	// remove beginning and trailing slash
	url := strings.Trim(r.URL.Path, "/")

	// split at every slash
	splitPath := strings.Split(url, "/")

	// if split url does not contain a possible siteID
	if len(splitPath) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload := splitPath[1:]

	go a.logHit(payload, r)

	// add db siteID check

	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	w.Write([]byte(strings.Join(payload, "/")))
}

func (a *API) logHit(payload []string, r *http.Request) {

}

func (a *API) cookie(r *http.Request) {

}
