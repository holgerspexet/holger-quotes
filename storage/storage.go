package storage

type QuoteInfo struct {
	Who   string
	When  string
	Quote string
}

type Store interface {
	Get() []QuoteInfo
	Store(quote QuoteInfo)
}

type MemoryStorage struct {
	quotes []QuoteInfo
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		quotes: []QuoteInfo{
			{When: "Igår", Quote: "Hejsan", Who: "Johan"},
			{When: "Idag", Quote: "Det var en gång", Who: "Astrid"},
			{When: "Imon", Quote: "All makt åt Tengil", Who: "Jonatan"},
		},
	}
}

func (ms MemoryStorage) Get() []QuoteInfo {
	return ms.quotes
}

func (ms *MemoryStorage) Store(quote QuoteInfo) {
	ms.quotes = append(ms.quotes, quote)
}
