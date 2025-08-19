package main

import (
	"comercial-backend/src/config"
	"comercial-backend/src/modules/categoria"

	"github.com/gin-gonic/gin"
)



func main() {
	
	var url string ="mongodb://localhost:27017"
	config.ConnectMongo(url, "ventas")

	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := config.MongoClient.Disconnect(ctx)

	if err != nil {
		fmt.Println("⚠️ Error al desconectar Mongo:", err)
	} else {
		fmt.Println("👋 Desconectado de MongoDB")
	}*/

	router := gin.Default()
	api:=router.Group("api")
	// router categorias
	api.POST("/categorias", categoria.CrearCategoria)
	api.GET("/categorias", categoria.ListarCategoria)
	router.Run(":5000")
}