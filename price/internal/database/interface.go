package database

import "github.com/Impisigmatus/PestControlExpert/price/autogen"

type Database interface {
	GetPrices() ([]autogen.Price, error)
}
