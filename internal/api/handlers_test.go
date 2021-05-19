package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/yvv4git/paxful/internal/repository"
	"github.com/yvv4git/paxful/internal/usecases"
	"github.com/yvv4git/paxful/tests"
)

func TestTransferFounds(t *testing.T) {
	cfg := tests.Config()
	mockWalletRepo := repository.NewWalletRepositoryMock()
	api := NewAPI(cfg, mockWalletRepo)
	webApp := api.WebApp()

	testCases := []struct {
		name             string
		request          func() (req *http.Request, err error)
		expectStatusCode int
		description      string
	}{
		{
			name: "Case-1",
			request: func() (req *http.Request, err error) {
				form := usecases.TransferFoundsForm{
					UserFromID:     1,
					UserToID:       2,
					Sum:            100.0,
					IdempotenceKey: repository.IdempotenceKeyGood(),
				}

				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("POST", "/api/transfer", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
			description:      "The first client transfers money to the second client",
		},
		{
			name: "Case-2",
			request: func() (req *http.Request, err error) {
				form := usecases.TransferFoundsForm{
					UserFromID:     1,
					UserToID:       2,
					Sum:            500.0,
					IdempotenceKey: repository.IdempotenceKeyGood(),
				}

				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("POST", "/api/transfer", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 200,
			description:      "The first client transfers all their money to the second client",
		},
		{
			name: "Case-3",
			request: func() (req *http.Request, err error) {
				form := usecases.TransferFoundsForm{
					UserFromID:     1,
					UserToID:       2,
					Sum:            600.0,
					IdempotenceKey: repository.IdempotenceKeyGood(),
				}

				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("POST", "/api/transfer", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 500,
			description:      "The first customer tries to transfer more money than they have",
		},
		{
			name: "Case-4",
			request: func() (req *http.Request, err error) {
				form := usecases.TransferFoundsForm{
					UserFromID:     1,
					UserToID:       2,
					Sum:            200.0,
					IdempotenceKey: repository.IdempotenceKeyIncorrect(),
				}

				jsonForm, _ := json.Marshal(form)
				req, err = http.NewRequest("POST", "/api/transfer", bytes.NewBuffer(jsonForm))
				req.Header.Set("Content-Type", "application/json")
				return
			},
			expectStatusCode: 500,
			description:      "The first client uses an invalid idempotency key",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			request, err := tc.request()
			if err != nil {
				t.Fatal(err)
			}

			result, err := webApp.Test(request)
			if err != nil {
				t.Fatal(err)
			}
			defer result.Body.Close()

			t.Log(result.StatusCode)
			bodyBytes, err := ioutil.ReadAll(result.Body)
			if err != nil {
				log.Fatal(err)
			}
			t.Log(string(bodyBytes))
		})
	}
}
