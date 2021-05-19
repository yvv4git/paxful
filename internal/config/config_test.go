package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	envHostDB     = fmt.Sprintf("%s_DB.HOST", prefixEnvironmet)
	envPortDB     = fmt.Sprintf("%s_DB.PORT", prefixEnvironmet)
	envUserDB     = fmt.Sprintf("%s_DB.USER", prefixEnvironmet)
	envPasswordDB = fmt.Sprintf("%s_DB.PASSWORD", prefixEnvironmet)
	envNameDB     = fmt.Sprintf("%s_DB.NAME", prefixEnvironmet)
	envFileName   = fmt.Sprintf("%s_DB.FILENAMEDB", prefixEnvironmet)
	envPortAPI    = fmt.Sprintf("%s_API.PORT", prefixEnvironmet)
)

func TestConfig(t *testing.T) {
	type environment struct {
		HostDB     string
		PortDB     int
		UserDB     string
		PasswordDB string
		NameDB     string
		PortAPI    int
	}

	setEnvironvent := func(env environment) {
		// key - value
		os.Setenv(envHostDB, env.HostDB)
		os.Setenv(envPortDB, fmt.Sprint(env.PortDB))
		os.Setenv(envUserDB, env.UserDB)
		os.Setenv(envPasswordDB, env.PasswordDB)
		os.Setenv(envNameDB, env.NameDB)
		os.Setenv(envPortAPI, fmt.Sprint(env.PortAPI))
	}

	unsetEnvironment := func() {
		os.Unsetenv(envHostDB)
		os.Unsetenv(envPortDB)
		os.Unsetenv(envUserDB)
		os.Unsetenv(envPasswordDB)
		os.Unsetenv(envNameDB)
		os.Unsetenv(envPortAPI)
	}

	testCases := []struct {
		name           string
		useEnvironment bool
		env            environment
		filePathConfig string
		expect         *Config
		wantError      bool
	}{
		{
			name:           "Config from environment",
			useEnvironment: true,
			env: environment{
				HostDB:     "localhost",
				PortDB:     3306,
				UserDB:     "api",
				PasswordDB: "secret",
				NameDB:     "paxful",
				PortAPI:    8181,
			},
			filePathConfig: "fixtures/test",
			expect: &Config{
				DB: DBConfig{
					Host:     "localhost",
					Port:     3306,
					User:     "api",
					Password: "secret",
					Name:     "paxful",
				},
				API: APIConfig{
					Port: 8181,
				},
			},
			wantError: false,
		},
		{
			name:           "Config from file",
			useEnvironment: false,
			env: environment{
				HostDB:     "localhost",
				PortDB:     3306,
				UserDB:     "api",
				PasswordDB: "secret",
				NameDB:     "paxful",
				PortAPI:    8181,
			},
			filePathConfig: "fixtures/test",
			expect: &Config{
				DB: DBConfig{
					Host:     "localhost",
					Port:     3306,
					User:     "api",
					Password: "secret",
					Name:     "paxful",
				},
				API: APIConfig{
					Port: 8080,
				},
			},
			wantError: false,
		},
		{
			name:           "File not found, then use environment",
			useEnvironment: true,
			env: environment{
				HostDB:     "localhost",
				PortDB:     3306,
				UserDB:     "api",
				PasswordDB: "secret",
				NameDB:     "paxful",
				PortAPI:    8181,
			},
			filePathConfig: "fixtures/unknown",
			expect: &Config{
				DB: DBConfig{
					Host:     "localhost",
					Port:     3306,
					User:     "api",
					Password: "secret",
					Name:     "paxful",
				},
				API: APIConfig{
					Port: 8181,
				},
			},
			wantError: false,
		},
		{
			name:           "Config file not found, environment empty, then use default",
			useEnvironment: false,
			filePathConfig: "fixtures/unknown",
			expect: &Config{
				DB: DBConfig{
					Host:     "localhost",
					Port:     3306,
					User:     "api",
					Password: "secret",
					Name:     "paxful",
				},
				API: APIConfig{
					Port: 1357,
				},
			},
			wantError: false,
		},
		{
			name:           "Bad file stucture, get error",
			useEnvironment: false,
			filePathConfig: "fixtures/bad",
			wantError:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup environment from test case.
			if tc.useEnvironment {
				setEnvironvent(tc.env)
			} else {
				unsetEnvironment()
			}

			// Init config with config file.
			result, err := Init(tc.filePathConfig)
			if tc.wantError {
				assert.NotEmpty(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tc.expect, result)
		})
	}
}

func TestParseFile(t *testing.T) {
	type parsedPath struct {
		dir      string
		fileName string
	}

	testCases := []struct {
		filePath string
		expect   parsedPath
	}{
		{
			filePath: "fixture/test",
			expect: parsedPath{
				dir:      "fixture",
				fileName: "test",
			},
		},
	}

	for _, tc := range testCases {
		dir, filename, err := parseFilePath(tc.filePath)
		assert.Nil(t, err)

		result := parsedPath{
			dir:      dir,
			fileName: filename,
		}
		assert.Equal(t, tc.expect, result, "Chief, we have a problem")
	}
}
