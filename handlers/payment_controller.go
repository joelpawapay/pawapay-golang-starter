package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/xceejay/pawapay-golang-starter/models"
)

func (server *Server) deposits(w http.ResponseWriter, r *http.Request) {

	var mobileMoney *models.MobileMoneyProvider

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(body, &mobileMoney)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	if mobileMoney.CreateProvider(); mobileMoney.DecimalsInAmount < 0 {
		respondWithError(w, http.StatusInternalServerError, errors.New("Invalid MNO").Error())
	}

	fmt.Println(*mobileMoney)
	apiURL := "https://api.sandbox.pawapay.cloud/deposits"
	jwt := `eyJraWQiOiIxIiwiYWxnIjoiSFMyNTYifQ.eyJqdGkiOiI1YmI0ZmIxOS1kYjczLTQzODAtYjMwNy0wZjdiMzI3NTg2YmQiLCJzdWIiOiIzODQiLCJpYXQiOjE3MDI5ODg3OTYsImV4cCI6MjAxODYwNzk5NiwicG0iOiJEQUYsUEFGIiwidHQiOiJBQVQifQ.UbQgvsRV5Chpr4W10LvVeSAxmLa27iiSJKaoNpMrx2U`

	respbody, err := makeRequest(apiURL, jwt, *mobileMoney)

	if err != nil {
		fmt.Println("Error:", err)
		respondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondWithJSON(w, http.StatusOK, string(respbody))
}

func (server *Server) payouts(w http.ResponseWriter, r *http.Request) {
	var mobileMoney *models.MobileMoneyProvider
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(body, mobileMoney)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

}

func makeRequest(apiURL, jwt string, provider models.MobileMoneyProvider) ([]byte, error) {
	requestBody, err := json.Marshal(provider)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	fmt.Println(string(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read")
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request failed with status code: %s", string(body))
	}

	// Handle the successful response if needed

	return body, nil
}
