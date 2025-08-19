package config

import (
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

func ConnectMongo(url string, nombreBd string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	
	defer cancel()

	cliente , err := mongo.Connect(options.Client().ApplyURI(url))
	if(err != nil){
		log.Fatalln("Error en la coneccion de la base de datos")
	}
	err = cliente.Ping(ctx, nil)
	if(err != nil){
		log.Fatalln("No se pudo hacer un ping a la base de datos")
	}
	MongoClient = cliente
	MongoDatabase = cliente.Database(nombreBd)
	fmt.Println("Se conecto a la base de datos")

}