package main

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/middleware"
	"comercial-backend/src/modules/categoria"
	"comercial-backend/src/modules/producto"
	"comercial-backend/src/modules/venta/router"

	"comercial-backend/src/modules/proveedor"
	routerStock "comercial-backend/src/modules/stock/router"
	"comercial-backend/src/modules/usuario"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	var url string = "mongodb://kanna:kanna@localhost:27017/comision?authSource=admin"
	//var url string = "mongodb://localhost:27017"
	config.ConnectMongo(url, "ventas")

	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := config.MongoClient.Disconnect(ctx)

	if err != nil {
		fmt.Println("⚠️ Error al desconectar Mongo:", err)
	} else {
		fmt.Println("👋 Desconectado de MongoDB")
	}*/

	route := gin.Default()
	route.Use(cors.Default())
	route.Use(middleware.ValidarTokenAtenticacion())
	api := route.Group("api")
	// router categorias
	categoria.RouterCategoria(api)
	proveedor.RouterProveedor(api)
	routerStock.RouterStock(api)
	producto.RouterProducto(api)
	usuario.UsuarioRouter(api)
	router.VentaRouter(api)

	//almacen.AlmacenRouter(api)
	route.Run(":3000")
}
