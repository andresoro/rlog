package postgres

import (
	"time"

	"github.com/andresoro/rlog/app/model"
)

// InsertEvent adds an interaction with a given page to the db
func (db *DB) InsertEvent(pv *model.Event) error {
	// prepare query
	q, err := db.conn.Prepare("INSERT INTO events VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	defer q.Close()
	if err != nil {
		return err
	}

	_, err = q.Exec(pv.SiteID,
		pv.Duration,
		pv.Host,
		pv.Path,
		pv.Referrer,
		pv.Unique,
		pv.Date)

	return err
}

// RetrieveEvent will return a single interaction given an id
func (db *DB) RetrieveEvent(id int64) (*model.Event, error) {
	// prepare query
	q, err := db.conn.Prepare("SELECT * FROM events WHERE id=$1")
	defer q.Close()
	if err != nil {
		return nil, err
	}

	var event model.Event
	row := q.QueryRow(id)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&event.StatID,
		&event.SiteID,
		&event.Duration,
		&event.Host,
		&event.Path,
		&event.Referrer,
		&event.Date,
		&event.Unique,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// RetrieveAllEvents will return all interactions on a given site
// todo: add time frame?
func (db *DB) RetrieveAllEvents(start, end time.Time, siteID int64) ([]*model.Event, error) {
	return nil, nil
}
