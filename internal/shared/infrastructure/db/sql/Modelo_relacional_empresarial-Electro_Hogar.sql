-- ============================================================
-- ESQUEMA DIMENSIONAL DE VENTAS (Modelo estrella)
-- ============================================================

CREATE TABLE CLIENTE (
    id_cliente     VARCHAR(40)     PRIMARY KEY,
    nombre         VARCHAR(150)    NOT NULL,
    email          VARCHAR(200)    UNIQUE,
    telefono       VARCHAR(20),
    segmento       VARCHAR(50),
    ciudad         VARCHAR(100),
    sexo           CHAR(1),
    edad           SMALLINT
);

COMMENT ON TABLE CLIENTE IS 'Dimensión que almacena los datos demográficos y de segmentación de cada cliente';
COMMENT ON COLUMN CLIENTE.id_cliente  IS 'Identificador único del cliente (UUID)';
COMMENT ON COLUMN CLIENTE.nombre      IS 'Nombre completo del cliente';
COMMENT ON COLUMN CLIENTE.email       IS 'Correo electrónico de contacto, debe ser único';
COMMENT ON COLUMN CLIENTE.telefono    IS 'Número de teléfono principal del cliente';
COMMENT ON COLUMN CLIENTE.segmento    IS 'Segmento comercial al que pertenece (ej: Premium, Estándar)';
COMMENT ON COLUMN CLIENTE.ciudad      IS 'Ciudad de residencia del cliente';
COMMENT ON COLUMN CLIENTE.sexo        IS 'Sexo del cliente: M=Masculino, F=Femenino, O=Otro';
COMMENT ON COLUMN CLIENTE.edad        IS 'Edad del cliente expresada en años';

CREATE TABLE TIEMPO (
    id_tiempo      VARCHAR(40)     PRIMARY KEY,
    fecha          DATE            NOT NULL UNIQUE,
    dia            SMALLINT,
    mes            SMALLINT,
    trimestre      SMALLINT,
    anio           SMALLINT        NOT NULL,
    dia_semana     VARCHAR(15),
    es_fin_semana  BOOLEAN         DEFAULT FALSE
);

COMMENT ON TABLE TIEMPO IS 'Dimensión que desglosa las fechas en atributos útiles para análisis temporal';
COMMENT ON COLUMN TIEMPO.id_tiempo     IS 'Identificador único de la fecha (UUID)';
COMMENT ON COLUMN TIEMPO.fecha         IS 'Fecha completa en formato YYYY-MM-DD';
COMMENT ON COLUMN TIEMPO.dia           IS 'Día del mes, valores entre 1 y 31';
COMMENT ON COLUMN TIEMPO.mes           IS 'Mes del año, valores entre 1 y 12';
COMMENT ON COLUMN TIEMPO.trimestre     IS 'Trimestre fiscal del año, valores entre 1 y 4';
COMMENT ON COLUMN TIEMPO.anio          IS 'Año de la fecha (ej: 2024)';
COMMENT ON COLUMN TIEMPO.dia_semana    IS 'Nombre del día de la semana (ej: Lunes, Martes)';
COMMENT ON COLUMN TIEMPO.es_fin_semana IS 'Indica si la fecha corresponde a un sábado o domingo';

CREATE TABLE TIENDA (
    id_ciudad      VARCHAR(40)     PRIMARY KEY,
    ciudad         VARCHAR(100)    NOT NULL,
    departamento   VARCHAR(100),
    region         VARCHAR(100)
);

COMMENT ON TABLE TIENDA IS 'Dimensión que almacena la información geográfica de cada tienda';
COMMENT ON COLUMN TIENDA.id_ciudad    IS 'Identificador único de la tienda o ciudad (UUID)';
COMMENT ON COLUMN TIENDA.ciudad       IS 'Ciudad donde está ubicada la tienda';
COMMENT ON COLUMN TIENDA.departamento IS 'Departamento o estado al que pertenece la tienda';
COMMENT ON COLUMN TIENDA.region       IS 'Región geográfica de la tienda';

CREATE TABLE PRODUCTO (
    id_producto      VARCHAR(40)     PRIMARY KEY,
    producto_nombre  VARCHAR(200)    NOT NULL,
    marca            VARCHAR(100),
    modelo           VARCHAR(100),
    categoria        VARCHAR(100),
    subcategoria     VARCHAR(100),
    precio           DECIMAL(12, 2)
);

COMMENT ON TABLE PRODUCTO IS 'Dimensión que contiene el catálogo de productos con su jerarquía de clasificación';
COMMENT ON COLUMN PRODUCTO.id_producto      IS 'Identificador único del producto (UUID)';
COMMENT ON COLUMN PRODUCTO.producto_nombre  IS 'Nombre descriptivo completo del producto';
COMMENT ON COLUMN PRODUCTO.marca            IS 'Marca o fabricante del producto';
COMMENT ON COLUMN PRODUCTO.modelo           IS 'Modelo o referencia específica del producto';
COMMENT ON COLUMN PRODUCTO.categoria        IS 'Categoría principal del producto';
COMMENT ON COLUMN PRODUCTO.subcategoria     IS 'Subcategoría detallada';
COMMENT ON COLUMN PRODUCTO.precio           IS 'Precio de lista oficial del producto';

CREATE TABLE VENTA (
    id_venta         VARCHAR(40)     PRIMARY KEY,
    id_tiempo        VARCHAR(40)     NOT NULL,
    id_cliente       VARCHAR(40)     NOT NULL,
    id_producto      VARCHAR(40)     NOT NULL,
    id_tienda        VARCHAR(40)     NOT NULL,
    cantidad         INT,
    precio_unitario  DECIMAL(12, 2),
    total_venta      DECIMAL(14, 2),
    costo_unitario   DECIMAL(12, 2),
    margen_unitario  DECIMAL(12, 2),
    margen_total     DECIMAL(14, 2),

    CONSTRAINT fk_venta_tiempo   FOREIGN KEY (id_tiempo)   REFERENCES TIEMPO   (id_tiempo),
    CONSTRAINT fk_venta_cliente  FOREIGN KEY (id_cliente)  REFERENCES CLIENTE  (id_cliente),
    CONSTRAINT fk_venta_producto FOREIGN KEY (id_producto) REFERENCES PRODUCTO (id_producto),
    CONSTRAINT fk_venta_tienda   FOREIGN KEY (id_tienda)   REFERENCES TIENDA   (id_ciudad)
);

COMMENT ON TABLE VENTA IS 'Tabla de hechos que registra cada transacción de venta con sus métricas numéricas';
COMMENT ON COLUMN VENTA.id_venta         IS 'Identificador único de la transacción (UUID)';
COMMENT ON COLUMN VENTA.id_tiempo        IS 'Referencia a la dimensión TIEMPO';
COMMENT ON COLUMN VENTA.id_cliente       IS 'Referencia a la dimensión CLIENTE';
COMMENT ON COLUMN VENTA.id_producto      IS 'Referencia a la dimensión PRODUCTO';
COMMENT ON COLUMN VENTA.id_tienda        IS 'Referencia a la dimensión TIENDA';
COMMENT ON COLUMN VENTA.cantidad         IS 'Número de unidades vendidas';
COMMENT ON COLUMN VENTA.precio_unitario  IS 'Precio real de venta por unidad';
COMMENT ON COLUMN VENTA.total_venta      IS 'Valor total de la venta (Cantidad x Precio)';
COMMENT ON COLUMN VENTA.costo_unitario   IS 'Costo de adquisición por unidad';
COMMENT ON COLUMN VENTA.margen_unitario  IS 'Ganancia por unidad (Precio - Costo)';
COMMENT ON COLUMN VENTA.margen_total     IS 'Ganancia total de la transacción';

CREATE INDEX idx_venta_tiempo   ON VENTA (id_tiempo);
CREATE INDEX idx_venta_cliente  ON VENTA (id_cliente);
CREATE INDEX idx_venta_producto ON VENTA (id_producto);
CREATE INDEX idx_venta_tienda   ON VENTA (id_tienda);