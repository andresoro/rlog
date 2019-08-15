package postgres

import (
	"time"

	"github.com/andresoro/rlog/pkg/model"
)

// InsertEvent adds an interaction with a given page to the db
func (db *DB) InsertEvent(pv *model.Event) error {
	// prepare query
	q, err := db.conn.Prepare("INSERT INTO events (site_id, event_key, addr, request_date, uniq) VALUES ($1, $2, $3, $4, $5)")
	defer q.Close()
	if err != nil {
		return err
	}

	_, err = q.Exec(pv.SiteID,
		pv.Key,
		pv.Addr,
		pv.Date,
		pv.Unique)

	return err
}

// RetrieveEvent will return a single interaction given an id
func (db *DB) RetrieveEvent(id int) (*model.Event, error) {
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
		&event.Key,
		&event.Addr,
		&event.Date,
		&event.Unique,
	)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// RetrieveSiteStats will return all interactions on a given site
func (db *DB) RetrieveSiteStats(start, end time.Time, siteID int) (*model.SiteStats, error) {
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

	defer rows.Close()
	for rows.Next() {
		var event model.Event
		err = rows.Scan(
			&event.ID,
			&event.SiteID,
			&event.Key,
			&event.Addr,
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
		ps.Add(&event)
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

// RetrieveAll will return all events for a site
func (db *DB) RetrieveAll(siteID int) ([]*model.Event, error) {

	events := make([]*model.Event, 0)

	q, err := db.conn.Prepare("SELECT * FROM events WHERE site_id=$1")
	if err != nil {
		return events, err
	}

	rows, err := q.Query(siteID)
	if err != nil {
		return events, err
	}

	defer rows.Close()
	for rows.Next() {
		var event model.Event
		err = rows.Scan(
			&event.ID,
			&event.SiteID,
			&event.Key,
			&event.Addr,
			&event.Date,
			&event.Unique,
		)
		if err != nil {
			return events, err
		}

		events = append(events, &event)
	}

	return events, nil
}
