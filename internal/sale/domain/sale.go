package domain

// Sale representa la entidad Venta en el proceso ETL (Fact Table)
type Sale struct {
	ID             string  `json:"id_venta"`
	IDTiempo       string  `json:"id_tiempo"`
	IDCliente      string  `json:"id_cliente"`
	IDProducto     string  `json:"id_producto"`
	IDTienda       string  `json:"id_tienda"`
	Cantidad       int     `json:"cantidad"`
	PrecioUnitario float64 `json:"precio_unitario"`
	TotalVenta     float64 `json:"total_venta"`
	CostoUnitario  float64 `json:"costo_unitario"`
	MargenUnitario float64 `json:"margen_unitario"`
	MargenTotal    float64 `json:"margen_total"`
}
