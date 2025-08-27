package main

import (
	"comercial-backend/src/core/config"
	"comercial-backend/src/core/middleware"
	"comercial-backend/src/modules/categoria"
	"comercial-backend/src/modules/producto"
	ventaRouter "comercial-backend/src/modules/venta/router"

	ingresoRouter "comercial-backend/src/modules/ingreso/router"
	"comercial-backend/src/modules/proveedor"
	routerStock "comercial-backend/src/modules/stock/router"
	sucursalRouter "comercial-backend/src/modules/sucursal/router"
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

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(middleware.ValidarTokenAtenticacion())
	api := router.Group("api")
	// router categorias
	categoria.RouterCategoria(api)
	proveedor.RouterProveedor(api)
	routerStock.RouterStock(api)
	producto.RouterProducto(api)
	usuario.UsuarioRouter(api)
	ventaRouter.VentaRouter(api)
	ventaRouter.DetalleVentaRouter(api)

	sucursalRouter.SucursalRouter(api)
	ingresoRouter.IngresoRouter(api)
	//almacen.AlmacenRouter(api)
	router.Run(":3000")
}
