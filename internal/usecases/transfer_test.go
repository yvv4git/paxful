package usecases

import (
	"testing"

	"github.com/yvv4git/paxful/internal/repository"
)

func TestTransferFoundsForm_Transfer(t *testing.T) {
	mockWalletRepo := repository.NewWalletRepositoryMock()

	type fields struct {
		UserFromID     int
		UserToID       int
		Sum            float64
		IdempotenceKey string
	}

	type args struct {
		walletRepository repository.WalletRepository
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Case-1",
			fields: fields{
				UserFromID:     1,
				UserToID:       2,
				Sum:            100.0,
				IdempotenceKey: repository.IdempotenceKeyGood(),
			},
			args: args{
				mockWalletRepo,
			},
			wantErr: false,
		},
		{
			name: "Case-2",
			fields: fields{
				UserFromID:     1,
				UserToID:       2,
				Sum:            500.0,
				IdempotenceKey: repository.IdempotenceKeyGood(),
			},
			args: args{
				mockWalletRepo,
			},
			wantErr: false,
		},
		{
			name: "Case-3",
			fields: fields{
				UserFromID:     1,
				UserToID:       2,
				Sum:            600.0,
				IdempotenceKey: repository.IdempotenceKeyGood(),
			},
			args: args{
				mockWalletRepo,
			},
			wantErr: true,
		},
		{
			name: "Case-4",
			fields: fields{
				UserFromID:     1,
				UserToID:       2,
				Sum:            200.0,
				IdempotenceKey: repository.IdempotenceKeyIncorrect(),
			},
			args: args{
				mockWalletRepo,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &TransferFoundsForm{
				UserFromID:     tt.fields.UserFromID,
				UserToID:       tt.fields.UserToID,
				Sum:            tt.fields.Sum,
				IdempotenceKey: tt.fields.IdempotenceKey,
			}
			if err := f.Transfer(tt.args.walletRepository); (err != nil) != tt.wantErr {
				t.Errorf("TransferFoundsForm.Transfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
