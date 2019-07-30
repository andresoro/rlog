package postgres

import (
	"time"

	"github.com/andresoro/rlog/pkg/model"
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

// RetrieveSiteStats will return all interactions on a given site
func (db *DB) RetrieveSiteStats(start, end time.Time, siteID int64) (*model.SiteStats, error) {
	// get all rows for this site id
	q, err := db.conn.Prepare("SELECT * FROM events WHERE (request_date BETWEEN $1 AND $2) AND (site_id=$3)")
	defer q.Close()
	if err != nil {
		return nil, err
	}

	rows, err := q.Query(start, end, siteID)
	if err != nil {
		return nil, err
	}

	// filter and aggregate individual events into PageStat for each path
	pageStats := make(map[string]*model.PageStats, 0)

	for rows.Next() {
		var event *model.Event
		err = rows.Scan(
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

		// if this is an event for a new path, create it with correct path name
		if _, ok := pageStats[event.Path]; !ok {
			pageStats[event.Path] = &model.PageStats{
				Path: event.Path,
			}
		}
		// aggregate
		ps := pageStats[event.Path]
		ps.Add(event)
	}

	// return aggregated site stats

	var siteStats = &model.SiteStats{
		SiteID: siteID,
		Pages:  make([]*model.PageStats, 0),
	}

	return siteStats, nil
}
