package postgres

import "github.com/andresoro/rlog/app/models"

// InsertEvent adds an interaction with a given page to the db
func (db *DB) InsertEvent(pv *models.Event) error {

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
func (db *DB) RetrieveEvent(id int64) (*models.Event, error) {
	q, err := db.conn.Prepare("SELECT * FROM events WHERE id=$1")
	defer q.Close()
	if err != nil {
		return nil, err
	}

	var event models.Event
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
