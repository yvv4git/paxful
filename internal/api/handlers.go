package api

import (
	"github.com/gofiber/fiber/v2"
	cfg "github.com/yvv4git/paxful/internal/config"
	"github.com/yvv4git/paxful/internal/repository"
	"github.com/yvv4git/paxful/internal/usecases"
)

type TransferFoundHandler struct {
	config           *cfg.Config
	walletRepository repository.WalletRepository
}

func (h *TransferFoundHandler) Post(c *fiber.Ctx) error {
	form := new(usecases.TransferFoundsForm)

	if err := c.BodyParser(form); err != nil {
		return err
	}

	if err := usecases.Validate(form); err != nil {
		return err
	}

	if err := form.Transfer(h.walletRepository); err != nil {
		return err
	}

	return c.
		Status(200).
		JSON(fiber.Map{
			"status": "success",
		})
}

func NewTransferFoundHandler(
	config *cfg.Config,
	walletRepository repository.WalletRepository,
) *TransferFoundHandler {
	return &TransferFoundHandler{
		config:           config,
		walletRepository: walletRepository,
	}
}
