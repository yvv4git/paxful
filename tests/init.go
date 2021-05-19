package tests

import (
	"database/sql"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/yvv4git/paxful/internal/config"
	"github.com/yvv4git/paxful/internal/repository/mysql"
)

const (
	configFile                      = "config/tests"
	idempotenceKey                  = "1e1c620a-b6cf-11eb-a6da-0242c0a85003"
	idempotenceKeyExpiredByAttempts = "7c708476-b79e-11eb-808a-0242c0a8d005"
	idempotenceKeyIncorrect         = "fkasjdlfjasdfjahdfjhasldjfhlashd"
)

var (
	cfg      *config.Config
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func init() {
	var err error
	// Change pwd.
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	// Init config.
	cfg, err = config.Init(configFile)
	if err != nil {
		log.Fatal(err)
	}
}

func Config() *config.Config {
	return cfg
}

func IdempotenceKey() string {
	return idempotenceKey
}

func IndempotenceKeyIncorrect() string {
	return idempotenceKeyIncorrect
}

func IdempotenceKeyExpiredByAttempts() string {
	return idempotenceKeyExpiredByAttempts
}

func PrepareTestDB() (db *sql.DB, err error) {
	db, err = mysql.NewDB(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)
	if err != nil {
		return nil, err
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("tests/fixtures"),
	)
	if err != nil {
		return nil, err
	}

	if err := fixtures.Load(); err != nil {
		return nil, err
	}

	return db, nil
}

func ResetTestDB() error {
	return fixtures.Load()
}
