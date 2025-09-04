package structstock

type Producto struct {
	Nombre string `bson:"nombre"`
}

type FiltrosStock struct {
	Codigo         string
	ProductoNombre string
	Categoria      string
	UnidadManejo   string
}
