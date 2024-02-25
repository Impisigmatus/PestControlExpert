package database

import "github.com/Impisigmatus/PestControlExpert/microservices/prices/autogen"

type Database interface {
	GetPrices() ([]autogen.Price, error)
}
