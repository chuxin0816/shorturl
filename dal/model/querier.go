package model

type ImpulseSender interface {
	// REPLACE INTO @@table (stub) VALUES ('a')
	ReplaceStub() error

	// SELECT LAST_INSERT_ID()
	LastInsertID() (uint64, error)
}
