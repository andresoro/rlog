package postgres

import (
	"time"

	"github.com/andresoro/rlog/pkg/model"
)

// InsertEvent adds an interaction with a given page to the db
func (db *DB) InsertEvent(pv *model.Event) error {
	// prepare query
	q, err := db.conn.Prepare("INSERT INTO events VALUES ($1, $2, $3, $4, $5, $6)")
	defer q.Close()
	if err != nil {
		return err
	}

	_, err = q.Exec(pv.SiteID,
		pv.Host,
		pv.Key,
		pv.Addr,
		pv.Date,
		pv.Unique)

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
		&event.ID,
		&event.SiteID,
		&event.Host,
		&event.Key,
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
	pageStats := make(map[string]*model.KeyStats, 0)

	for rows.Next() {
		var event *model.Event
		err = rows.Scan(
			&event.ID,
			&event.SiteID,
			&event.Host,
			&event.Key,
			&event.Date,
			&event.Unique,
		)
		if err != nil {
			return nil, err
		}

		// if this is an event for a new key, create it with correct path name
		if _, ok := pageStats[event.Key]; !ok {
			pageStats[event.Key] = &model.KeyStats{
				Key: event.Key,
			}
		}
		// aggregate
		ps := pageStats[event.Key]
		ps.Add(event)
	}

	// return aggregated site stats

	var siteStats = &model.SiteStats{
		SiteID: siteID,
		Pages:  make([]*model.KeyStats, 0),
	}

	for _, p := range pageStats {
		_ = append(siteStats.Pages, p)
	}

	return siteStats, nil
}
