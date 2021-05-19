package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	cfg "github.com/yvv4git/paxful/internal/config"
	"github.com/yvv4git/paxful/internal/helpers"
	"github.com/yvv4git/paxful/internal/repository"
)

// API is entity of web api.
type API struct {
	config *cfg.Config
	db     *sql.DB
	webApp *fiber.App
}

// NewAPI is used as constructor for web api.
func NewAPI(
	config *cfg.Config,
	walletRepository repository.WalletRepository,
) *API {
	webApp := fiber.New(fiber.Config{})
	transferFoundsHandler := NewTransferFoundHandler(config, walletRepository)

	api := webApp.Group("/api/")
	api.Post("transfer", transferFoundsHandler.Post)

	return &API{
		webApp: webApp,
		config: config,
	}
}

// WebApp is used as getter for fiber.app.
func (a *API) WebApp() *fiber.App {
	return a.webApp
}

// Start is used for run web api server.
func (a *API) Start() error {
	return a.webApp.Listen(
		helpers.ServerAddr(
			a.config.API.Host,
			a.config.API.Port,
		),
	)
}
