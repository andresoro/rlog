package postgres

// InsertSite adds a site to postgres database and returns its ID
func (db *DB) InsertSite(name string) (int64, error) {
	id := 0

	q, err := db.conn.Prepare("INSERT INTO sites(name) VALUES ($1) RETURNING id")
	defer q.Close()
	if err != nil {
		return int64(id), err
	}

	err = q.QueryRow(name).Scan(&id)
	if err != nil {
		return int64(id), err
	}

	return int64(id), nil
}
