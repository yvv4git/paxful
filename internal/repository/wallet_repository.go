package repository

//go:generate mockery --name=WalletRepository --output=. --filename=mock.go --inpackage
type WalletRepository interface {
	TransferFound(fromUserID int, toUserID int, sum float64, idempotenceKey string) error
}
