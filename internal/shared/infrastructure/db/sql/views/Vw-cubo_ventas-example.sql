-- ============================================================
-- VISTA LÓGICA DE VENTAS (EJEMPLO CON ESQUEMA)
-- ============================================================
-- Si el esquema es 'electro_hogar'

CREATE OR REPLACE VIEW electro_hogar.Vw_cubo_ventas AS
SELECT 
    t.mes              AS "Mes",
    p.subcategoria     AS "Subcategoría",
    ti.region          AS "Región",
    ti.departamento    AS "Departamento",
    ti.ciudad          AS "Ciudad",
    SUM(v.cantidad)    AS "Cantidad Vendida",
    SUM(v.total_venta) AS "Total Venta",
    SUM(v.costo_unitario) AS "Costo Unitario Total",
    SUM(v.margen_total)  AS "Margen Total"
FROM 
    electro_hogar.VENTA v
JOIN electro_hogar.TIEMPO t   ON v.id_tiempo = t.id_tiempo
JOIN electro_hogar.PRODUCTO p ON v.id_producto = p.id_producto
JOIN electro_hogar.TIENDA ti  ON v.id_tienda = ti.id_ciudad
GROUP BY 
    t.mes, 
    p.subcategoria, 
    ti.region, 
    ti.departamento, 
    ti.ciudad;

COMMENT ON VIEW electro_hogar.Vw_cubo_ventas IS 'Vista de agregación que simula un cubo para análisis de métricas de ventas por mes, subcategoría y ubicación geográfica';
