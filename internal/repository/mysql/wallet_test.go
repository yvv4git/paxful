package mysql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestWalletRepository_TransferFound(t *testing.T) {
	type argsTransferFound struct {
		fromUserID     int
		toUserID       int
		sum            float64
		idempotenceKey string
	}

	testCases := []struct {
		name string
		mock func(
			mock sqlmock.Sqlmock,
			args argsTransferFound,
			balanceFromUser float64,
			sumWithPercentage float64,
		)
		args              argsTransferFound
		balanceFromUser   float64
		sumWithPercentage float64
		wantErr           bool
	}{
		{
			name: "case-1",
			mock: func(mock sqlmock.Sqlmock, args argsTransferFound, balanceFromUser float64, sumWithPercentage float64) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"id", "attempt"}).AddRow(1, 0)
				mock.ExpectQuery(`SELECT t.id, t.attempt FROM transactions t WHERE idempotence_key =  \?`).
					WithArgs(args.idempotenceKey).
					WillReturnRows(rows)

				rows = sqlmock.NewRows([]string{"balance"}).AddRow(balanceFromUser)
				mock.ExpectQuery(`SELECT w.balance FROM users u JOIN wallets w ON w.users_id = u.id WHERE u.id = \?`).
					WithArgs(args.fromUserID).
					WillReturnRows(rows)

				mock.ExpectExec(`UPDATE wallets SET balance \= balance \- \? WHERE users_id = \?`).
					WithArgs(sumWithPercentage, args.fromUserID).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec(`UPDATE wallets SET balance \= balance \+ \? WHERE users_id = \?`).
					WithArgs(args.sum, args.toUserID).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectExec(`UPDATE transactions SET attempt \= attempt \+ 1 WHERE id = \?`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			args: argsTransferFound{
				fromUserID:     1,
				toUserID:       2,
				sum:            100.0,
				idempotenceKey: "649e81d6-b729-11eb-b9d7-0242c0a8c002",
			},
			balanceFromUser:   500.0,
			sumWithPercentage: 101.5,
			wantErr:           false,
		},
		{
			name: "case-2",
			mock: func(mock sqlmock.Sqlmock, args argsTransferFound, balanceFromUser float64, sumWithPercentage float64) {
				mock.ExpectBegin()

				rows := sqlmock.NewRows([]string{"balance"}).AddRow(balanceFromUser)
				mock.ExpectQuery(`SELECT w.balance FROM users u JOIN wallets w ON w.users_id = u.id WHERE u.id = \?`).
					WithArgs(args.fromUserID).
					WillReturnRows(rows)
			},
			args: argsTransferFound{
				fromUserID:     1,
				toUserID:       2,
				sum:            600.0,
				idempotenceKey: "649e81d6-b729-11eb-b9d7-0242c0a8c002",
			},
			balanceFromUser:   500.0,
			sumWithPercentage: 609.0,
			wantErr:           true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			defer db.Close()
			assert.Nil(t, err)

			tc.mock(mock, tc.args, tc.balanceFromUser, tc.sumWithPercentage)
			repo := NewWalletRepository(db)
			err = repo.TransferFound(tc.args.fromUserID, tc.args.toUserID, tc.args.sum, tc.args.idempotenceKey)
			if tc.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
