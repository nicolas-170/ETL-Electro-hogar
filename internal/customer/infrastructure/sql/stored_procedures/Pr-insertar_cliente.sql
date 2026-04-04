-- Activa la extensión para UUID, mas en una versión antigua de Postgres
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE PROCEDURE insertar_cliente(
    p_id_cliente VARCHAR,
    p_nombre     VARCHAR,
    p_email      VARCHAR,
    p_telefono   VARCHAR,
    p_segmento   VARCHAR,
    p_ciudad     VARCHAR,
    p_sexo       CHAR,
    p_edad       SMALLINT
)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO CLIENTE (
        id_cliente, 
        nombre, 
        email, 
        telefono, 
        segmento, 
        ciudad, 
        sexo, 
        edad
    )
    VALUES (
        p_id_cliente, 
        p_nombre, 
        p_email, 
        p_telefono, 
        p_segmento, 
        p_ciudad, 
        p_sexo, 
        p_edad
    );
    
    COMMIT;
END;
$$;