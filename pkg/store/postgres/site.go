package postgres

import (
	"github.com/andresoro/rlog/pkg/model"
)

// InsertSite adds a site to postgres database and returns its ID
func (db *DB) InsertSite(name string) (int, error) {
	id := 0

	q, err := db.conn.Prepare("INSERT INTO sites(name) VALUES ($1) RETURNING id")
	defer q.Close()
	if err != nil {
		return int(id), err
	}

	err = q.QueryRow(name).Scan(&id)
	if err != nil {
		return int(id), err
	}

	return int(id), nil
}

// GetSites all sites in db
func (db *DB) GetSites() ([]*model.Site, error) {
	q, err := db.conn.Prepare("SELECT * FROM sites")
	defer q.Close()
	if err != nil {
		return nil, err
	}

	rows, err := q.Query()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	sites := make([]*model.Site, 0)

	for rows.Next() {
		var site model.Site
		err = rows.Scan(
			&site.ID,
			&site.Name,
		)
		if err != nil {
			return sites, err
		}

		sites = append(sites, &site)
	}

	return sites, nil
}
