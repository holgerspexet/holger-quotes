package storage

import (
	"time"
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
