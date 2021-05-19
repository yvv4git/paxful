package domain

// Users is used as the entity of the database structure.
type Users struct {
	ID         int
	FirstName  string
	LastName   string
	MiddleName string
	Email      string
}
