package domain

// Transactions is used as the entity of the database structure.
type Transactions struct {
	ID             int
	IdempotenceKey string
	Expired        string
	Attempt        int
}
