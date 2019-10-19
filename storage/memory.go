package storage

// MemoryStorage stores QuoteInfos in memory
type MemoryStorage struct {
	quotes []QuoteInfo
}

// NewMemoryStorage creates a new MemoryStorage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		quotes: []QuoteInfo{},
	}
}

// Get loads all QuoteInfos
func (ms MemoryStorage) Get() []QuoteInfo {
	return ms.quotes
}

// Store inserts one more quote inte to the memory storage
func (ms *MemoryStorage) Store(quote QuoteInfo) {
	ms.quotes = append(ms.quotes, quote)
}
