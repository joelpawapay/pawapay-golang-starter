package handlers

import (
	"io"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/xceejay/pawapay-golang-starter/models"
)

/
func (server *Server) deposits(w http.ResponseWriter, r *http.Request) {
	var mobileMoney *models.MobileMoneyProvider
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(body,  mobileMoney)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

}

func (server *Server) payouts(w http.ResponseWriter, r *http.Request) {
	var mobileMoney *models.MobileMoneyProvider
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = json.Unmarshal(body,  mobileMoney)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}

}