package database

import (
	"fmt"

	"github.com/Impisigmatus/PestControlExpert/price/autogen"
)

func (pg *Postgres) GetPrices() ([]autogen.Price, error) {
	const query = "SELECT name, description, standart, premium FROM main.price;"

	prices := make([]autogen.Price, 0)
	if err := pg.db.Select(&prices, query); err != nil {
		return nil, fmt.Errorf("Invalid SELECT main.price: %s", err)
	}

	return prices, nil
}
