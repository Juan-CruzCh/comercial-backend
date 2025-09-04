package structCore

import "go.mongodb.org/mongo-driver/v2/bson"

type ResultadoPaginado struct {
	Data    []bson.M
	Paginas int
}
