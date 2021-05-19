package domain

// Wallets is used as the entity of the database structure.
type Wallets struct {
	ID      int
	UsersID int
	Balance float64
}
