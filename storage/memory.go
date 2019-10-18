package storage

type MemoryStorage struct {
	quotes []QuoteInfo
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		quotes: []QuoteInfo{},
	}
}

func (ms MemoryStorage) Get() []QuoteInfo {
	return ms.quotes
}

func (ms *MemoryStorage) Store(quote QuoteInfo) {
	ms.quotes = append(ms.quotes, quote)
}
