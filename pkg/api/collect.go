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
	// remove beginning and trailing slash
	url := strings.Trim(r.URL.Path, "/")

	// split at every slash
	splitPath := strings.Split(url, "/")

	// if split url does not contain a possible siteID
	if len(splitPath) < 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// path minus the /collect/ prefix i.e siteID followed by params
	var (
		siteID string
		key    string
	)

	payload := splitPath[1:]
	siteID = payload[0]

	// if key is included in url along with siteID
	if len(payload) > 1 {
		key = strings.Join(payload[1:], "/")
	}

	// add db siteID check
	a.logHit(siteID, key, r)

	w.Header().Set("Content-Type", "image/gif")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, private")
	w.Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	w.Write([]byte(Pixel))
}

func (a *API) logHit(siteID, key string, r *http.Request) error {

	id, err := strconv.Atoi(siteID)
	if err != nil {
		return err
	}

	event := &model.Event{
		SiteID: int64(id),
		Host:   r.Host,
		Key:    key,
		Addr:   r.RemoteAddr,
		Date:   time.Now(),
		Unique: true,
	}

	err = a.db.InsertEvent(event)
	if err != nil {
		return err
	}

	return nil
}

func (a *API) cookie(r *http.Request) {

}
