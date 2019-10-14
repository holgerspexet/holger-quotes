package storage

type QuoteInfo struct {
	CreatedAt string
	CreatedBy string
	Quote     string
	Quoted    string
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
			{CreatedAt: "Igår", CreatedBy: "Ingen", Quote: "Hejsan", Quoted: "Johan"},
			{CreatedAt: "Idag", CreatedBy: "Jag", Quote: "Det var en gång", Quoted: "Astrid"},
			{CreatedAt: "Imon", CreatedBy: "Någon annan", Quote: "All makt åt Tengil", Quoted: "Jonatan"},
		},
	}
}

func (ms MemoryStorage) Get() []QuoteInfo {
	return ms.quotes
}

func (ms *MemoryStorage) Store(quote QuoteInfo) {
	ms.quotes = append(ms.quotes, quote)
}
