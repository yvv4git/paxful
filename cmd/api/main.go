package main

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/yvv4git/paxful/internal/api"
	"github.com/yvv4git/paxful/internal/config"
	"github.com/yvv4git/paxful/internal/repository/mysql"
)

const (
	configFile = "config/production"
)

func main() {
	var err error
	var cfg *config.Config
	var db *sql.DB

	// Initializing the configuration.
	cfg, err = config.Init(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Initializing the database connection.
	db, err = mysql.NewDB(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)
	log.Println(db.Ping())
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Println("Close data base connection")
		db.Close()
	}()

	// For graceful shutdown.
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Start server")
		walletRepository := mysql.NewWalletRepository(db)
		webApi := api.NewAPI(cfg, walletRepository)
		log.Fatal(webApi.Start())
		log.Println("End")
	}()

	// Gracefull exit.
	<-exit
	log.Println("Stopping app...")
}
