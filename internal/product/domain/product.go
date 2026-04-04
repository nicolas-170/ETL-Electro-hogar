package domain

// Product define el modelo de dominio para un producto en el proceso ETL
type Product struct {
	ID           string  `json:"id_producto"`
	Nombre       string  `json:"producto_nombre"`
	Marca        string  `json:"marca"`
	Modelo       string  `json:"modelo"`
	Categoria    string  `json:"categoria"`
	Subcategoria string  `json:"subcategoria"`
	Precio       float64 `json:"precio"`
}
