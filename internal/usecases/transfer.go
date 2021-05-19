package usecases

import (
	"github.com/yvv4git/paxful/internal/repository"
)

// TransferFoundsForm used for form validation and some logic level.
type TransferFoundsForm struct {
	UserFromID     int     `valid:"type(int)" json:"idfrom" form:"idfrom"`
	UserToID       int     `valid:"type(int)" json:"idto" form:"idto"`
	Sum            float64 `valid:"type(float64)" json:"sum" form:"sum"`
	IdempotenceKey string  `valid:"length(0|50)" json:"idempotencekey" form:"idempotencekey"`
}

// Transfer is used as method of form.
func (f *TransferFoundsForm) Transfer(walletRepository repository.WalletRepository) (err error) {
	return walletRepository.TransferFound(f.UserFromID, f.UserToID, f.Sum, f.IdempotenceKey)
}
