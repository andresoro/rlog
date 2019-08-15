package api

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/andresoro/rlog/pkg/model"
)

// Pixel GIF 1x1 1x1+0+0 8-bit sRGB 35 bytes
const Pixel = "\x47\x49\x46\x38\x39\x61\x01\x00\x01\x00\x80\xFF\x00\xFF\xFF\xFF\x00\x00\x00\x2C\x00\x00\x00\x00\x01\x00\x01\x00\x00\x02\x02\x44\x01\x00\x3B"

// Collect endpoint to handle tracking client data
// this is a GET endpoint that embeds an image into the client browser and
// records request data according to the siteID and path provided
// the path after siteID is a way to track specific URLs as that is not possible
// by embeding images alone
// e.x /collect/{siteID}/any/url/path
func (a *API) Collect(w http.ResponseWriter, r *http.Request) {

	// parse URL for site ID and request key, request key is anything after ID (could be nil)

	//trim slashes and slice at each intermediary slash
	trim := strings.Trim(r.URL.Path, "/")
	splitPath := strings.Split(trim, "/")

	// if path does not even contain a possible siteID -> bad request
	if len(splitPath) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error: need siteID"))
		return
	}

	// ignore '/collect/' in payload, siteID should be following param
	payload := splitPath[1:]
	id := payload[0]

	// ensure id is an int aka valid siteID
	siteID, err := strconv.Atoi(id)
	if err != nil {
		if err == strconv.ErrSyntax {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error: siteID is not valid"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var key string
	if len(payload) > 1 {
		// if key is included in url
		key = strings.Join(payload[1:], "/")
	}

	// test if site exists

	a.logHit(siteID, key, r)

	// if request is good, write pixel
	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, private")
	w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	w.Write([]byte(Pixel))
}

func (a *API) logHit(siteID int, key string, r *http.Request) error {

	event := &model.Event{
		SiteID: int(siteID),
		Key:    key,
		Addr:   r.RemoteAddr,
		Date:   time.Now(),
		Unique: true,
	}

	err := a.db.InsertEvent(event)
	if err != nil {
		return err
	}

	return nil
}

func (a *API) cookie(r *http.Request) {

}
