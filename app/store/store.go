package store

// Store interface for handling app models with database
type Store interface {
	Connect() error
	Ping() error
}
