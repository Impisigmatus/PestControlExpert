package service

import (
	"fmt"
	"net/http"

	"github.com/Impisigmatus/PestControlExpert/price/autogen"
	"github.com/Impisigmatus/PestControlExpert/price/internal/database"
	"github.com/Impisigmatus/PestControlExpert/price/internal/utils"
)

type Transport struct {
	db database.Database
}

func NewTransport(cfg database.PostgresConfig) autogen.ServerInterface {
	return &Transport{
		db: database.NewPostgres(cfg),
	}
}

func (transport *Transport) GetApiPrices(w http.ResponseWriter, r *http.Request) {
	prices, err := transport.db.GetPrices()
	if err != nil {
		utils.WriteString(w, http.StatusInternalServerError, fmt.Errorf("Invalid DB GetPrices(): %s", err), "Не удалось получить цены")
		return
	}

	if len(prices) == 0 {
		utils.WriteNoContent(w)
		return
	}

	utils.WriteObject(w, prices)
}
