package mysql

import (
	"database/sql"

	"github.com/yvv4git/paxful/internal/helpers"
	"github.com/yvv4git/paxful/internal/repository"
)

// WalletRepository specific implementation of interface.
type WalletRepository struct {
	db *sql.DB
}

// NewWalletRepository is used as constructor.
func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

// TransferFound is used for transfer Found method for transferring funds
func (w *WalletRepository) TransferFound(
	fromUserID int,
	toUserID int,
	sum float64,
	idempotenceKey string,
) error {
	sumWithPercentage, err := helpers.PercentageOneHalf(sum)
	if err != nil {
		return err
	}

	tx, err := w.db.Begin()
	if err != nil {
		return err
	}

	// Check idempotence-key.
	row := tx.QueryRow(`SELECT t.id, t.attempt FROM transactions t WHERE idempotence_key = ? and attempt = 0`, idempotenceKey)
	var transactionsId, attempt int
	err = row.Scan(&transactionsId, &attempt)
	if err != nil && err == sql.ErrNoRows {
		tx.Rollback()
		return repository.ErrIdempontenceKeyNoFound
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	if attempt > 0 {
		tx.Rollback()
		return repository.ErrIdempotenceKeyAttemptLimit
	}

	// Check user(from) balance.
	row = tx.QueryRow(`SELECT w.balance FROM users u JOIN wallets w ON w.users_id = u.id WHERE u.id = ?`, fromUserID)
	var balanceUserFrom float64
	err = row.Scan(&balanceUserFrom)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Update balance.
	if sumWithPercentage > balanceUserFrom {
		tx.Rollback()
		return repository.ErrInsufficientFunds
	}

	// Reduce the value at the sender.
	_, err = tx.Exec(`UPDATE wallets SET balance = balance - ? WHERE users_id = ?`, sumWithPercentage, fromUserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Increase the value of the recipient.
	_, err = tx.Exec(`UPDATE wallets SET balance = balance + ? WHERE users_id = ?`, sum, toUserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Idempotence key attempt increase.
	_, err = tx.Exec(`UPDATE transactions SET attempt = attempt + 1 WHERE id = ?`, transactionsId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
