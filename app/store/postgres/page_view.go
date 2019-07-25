package postgres

import "github.com/andresoro/rlog/app/models"

// InsertPageView adds an interaction with a given page to the db
func (db *DB) InsertEvent(pv *models.Event) error {
	query := "INSERT INTO page_views VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	_, err := db.conn.Exec(query,
		pv.SiteID,
		pv.Duration,
		pv.Host,
		pv.Path,
		pv.Referrer,
		pv.Unique,
		pv.Date,
	)

	return err
}
