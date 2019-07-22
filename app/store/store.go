package store

// Store - interface for handling app models against an arbitrary database
// this interface is, ideally, the contract between the api and the database
type Store interface {
	// DB init methods to start and test connection/status of the underlying database
	Connect() error
	Ping() error
	Close() error

	// DB model methods to actually write and retrieve our data
}
