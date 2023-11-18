--PRIMARY KEYS

alter table cliente add constraint cliente_pk primary key (id_usuarie);
alter table direccion_entrega add constraint direccion_entrega_pk primary key (id_usuarie,id_direccion_entrega);
alter table tarifa_entrega add constraint tarifa_entrega_pk primary key (codigo_postal_corto);
alter table producto add constraint producto_pk primary key (id_producto);
alter table pedido add constraint pedido_pk primary key (id_pedido);
alter table pedido_detalle add constraint alquiler_pk primary key (id_pedido,id_producto);
alter table reposicion add constraint reposicion_pk primary key (id_producto,fecha_solicitud);
alter table error add constraint error_pk primary key (id_error);
alter table envio_email add constraint envio_email_pk primary key (id_email);
alter table entrada_trx_pedido add constraint entrada_trx_pedido_pk primary key (id_orden);


-- FOREIGN KEYS 

alter table direccion_entrega  add constraint id_usuarie_fk foreign key (id_usuarie) references cliente (id_usuarie);

alter table pedido add constraint id_usuarie_fk foreign key (id_usuarie) references cliente(id_usuarie);
--alter table pedido add constraint id_direccion_entrega_fk foreign key (id_direccion_entrega) references direccion_entrega(id_direccion_entrega);

alter table pedido_detalle add constraint id_pedido_fk foreign key (id_pedido) references pedido (id_pedido);
alter table pedido_detalle add constraint id_producto_fk foreign key (id_producto) references producto (id_producto);

alter table reposicion add constraint id_producto_fk foreign key (id_producto) references producto (id_producto);

--alter table error add constraint id_usuarie_fk foreign key (id_usuarie) references cliente(id_usuarie);
--alter table error add constraint id_pedido_fk foreign key (id_pedido) references pedido(id_pedido);
--alter table error add constraint id_direccion_entrega_fk foreign key (id_direccion_entrega) references direccion_entrega(id_direccion_entrega);
--alter table error add constraint id_producto_fk foreign key (id_producto) references producto(id_producto);

--alter table entrada_trx_pedido add constraint id_usuarie_fk foreign key (id_usuarie) references cliente(id_usuarie);
--alter table entrada_trx_pedido add constraint id_direccion_entrega_fk foreign key (id_direccion_entrega) references direccion_entrega(id_direccion_entrega);
--alter table entrada_trx_pedido add constraint id_pedido_fk foreign key (id_pedido) references pedido(id_pedido);
--alter table entrada_trx_pedido add constraint id_producto_fk foreign key (id_producto) references producto(id_producto);
