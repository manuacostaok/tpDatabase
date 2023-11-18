-- Eliminar claves foráneas

ALTER TABLE direccion_entrega DROP CONSTRAINT id_usuarie_fk;
ALTER TABLE pedido DROP CONSTRAINT id_usuarie_fk;
--ALTER TABLE pedido DROP CONSTRAINT id_direccion_entrega_fk;
ALTER TABLE pedido_detalle DROP CONSTRAINT id_pedido_fk;
ALTER TABLE pedido_detalle DROP CONSTRAINT id_producto_fk;
ALTER TABLE reposicion DROP CONSTRAINT id_producto_fk;
--ALTER TABLE error DROP CONSTRAINT id_usuarie_fk;
--ALTER TABLE error DROP CONSTRAINT id_pedido_fk;
--ALTER TABLE error DROP CONSTRAINT id_direccion_entrega_fk;
--ALTER TABLE error DROP CONSTRAINT id_producto_fk;
--ALTER TABLE entrada_trx_pedido DROP CONSTRAINT id_usuarie_fk;
--ALTER TABLE entrada_trx_pedido DROP CONSTRAINT id_direccion_entrega_fk;
--ALTER TABLE entrada_trx_pedido DROP CONSTRAINT id_pedido_fk;
--ALTER TABLE entrada_trx_pedido DROP CONSTRAINT id_producto_fk;

-- Eliminar claves primarias
ALTER TABLE cliente DROP CONSTRAINT cliente_pk;
ALTER TABLE direccion_entrega DROP CONSTRAINT direccion_entrega_pk;
ALTER TABLE tarifa_entrega DROP CONSTRAINT tarifa_entrega_pk;
ALTER TABLE producto DROP CONSTRAINT producto_pk;
ALTER TABLE pedido DROP CONSTRAINT pedido_pk;
ALTER TABLE pedido_detalle DROP CONSTRAINT alquiler_pk;
ALTER TABLE reposicion DROP CONSTRAINT reposicion_pk;
ALTER TABLE error DROP CONSTRAINT error_pk;
ALTER TABLE envio_email DROP CONSTRAINT envio_email_pk;
ALTER TABLE entrada_trx_pedido DROP CONSTRAINT entrada_trx_pedido_pk;
