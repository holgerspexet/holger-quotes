package storage

import (
	"log"
	"time"

	"github.com/holgerspexet/holger-quotes/config"
)

// QuoteInfo holds information about a single quote
type QuoteInfo struct {
	Quote string
	Who   string
	Where string
	When  time.Time
}

// Store is the interface that is used to store and load QuoteInfo's
type Store interface {
	Get() []QuoteInfo
	Store(quote QuoteInfo)
}

// CreateStorage creates and configures a new storage instance from the provided configuration
func CreateStorage(conf config.Config) Store {
	switch conf.StorageType {
	case config.StorageTypeSQLight:
		return NewSQLightStorage(conf.SQLightPath)
	case config.StorageTypeMemory:
		return NewMemoryStorage()
	default:
		log.Fatalf("Unsupported storageType: %s", conf.StorageType)
		return nil
	}
}
