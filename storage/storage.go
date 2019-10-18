package storage

import (
	"time"
)

type QuoteInfo struct {
	Quote string
	Who   string
	Where string
	When  time.Time
}

type Store interface {
	Get() []QuoteInfo
	Store(quote QuoteInfo)
}
