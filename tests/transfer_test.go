package tests

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yvv4git/paxful/internal/repository/mysql"
	"github.com/yvv4git/paxful/internal/usecases"
)

func TestMysql_TransferFound(t *testing.T) {
	db, err := PrepareTestDB()
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	type fields struct {
		UserFromID     int
		UserToID       int
		Sum            float64
		IdempotenceKey string
	}

	type args struct {
		db *sql.DB
	}

	testsCases := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		description string
	}{
		{
			name: "case-1",
			fields: fields{
				UserFromID:     1,
				UserToID:       2,
				Sum:            100,
				IdempotenceKey: IdempotenceKey(),
			},
			args: args{
				db: db,
			},
			wantErr:     false,
			description: "The first client sends the second a certain amount",
		},
		{
			name: "case-2",
			fields: fields{
				UserFromID:     1,
				UserToID:       2,
				Sum:            492,
				IdempotenceKey: IdempotenceKey(),
			},
			args: args{
				db: db,
			},
			wantErr:     false,
			description: "The first customer sends the second all their money",
		},
		{
			name: "case-3",
			fields: fields{
				UserFromID:     1,
				UserToID:       2,
				Sum:            600,
				IdempotenceKey: IdempotenceKey(),
			},
			args: args{
				db: db,
			},
			wantErr:     true,
			description: "The first customer doesn't have enough money",
		},
		{
			name: "case-4",
			fields: fields{
				UserFromID:     1,
				UserToID:       2,
				Sum:            200,
				IdempotenceKey: "bad-idempotece-key",
			},
			args: args{
				db: db,
			},
			wantErr:     true,
			description: "Using an invalid idempotence key",
		},
		{
			name: "case-5",
			fields: fields{
				UserFromID:     2,
				UserToID:       1,
				Sum:            200,
				IdempotenceKey: IdempotenceKeyExpiredByAttempts(),
			},
			args: args{
				db: db,
			},
			wantErr:     true,
			description: "Exceeded the number of idempotence key attempts",
		},
	}

	for _, tс := range testsCases {
		t.Run(tс.name, func(t *testing.T) {
			f := &usecases.TransferFoundsForm{
				UserFromID:     tс.fields.UserFromID,
				UserToID:       tс.fields.UserToID,
				Sum:            tс.fields.Sum,
				IdempotenceKey: tс.fields.IdempotenceKey,
			}

			err := ResetTestDB()
			assert.Nil(t, err)

			walletRepository := mysql.NewWalletRepository(tс.args.db)
			if err := f.Transfer(walletRepository); (err != nil) != tс.wantErr {
				t.Errorf("TransferFoundsForm.Transfer() error = %v, wantErr %v", err, tс.wantErr)
			}
		})
	}
}
