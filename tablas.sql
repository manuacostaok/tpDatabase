
create table cliente(
	id_usuarie    int,
	nombre   varchar(64),
	apellido varchar(64),
	dni int,
	fecha_nacimiento date,
	telefono char(12),
	email varchar(64)
);

create table direccion_entrega(
	id_usuarie           int,
	id_direccion_entrega  int,
	direccion varchar(64),
	localidad	varchar(64),
	codigo_postal char(8)
);


create table tarifa_entrega(
	codigo_postal_corto char(4),
	costo float
);

create table producto(
	id_producto   int,
	nombre varchar(64),
	precio_unitario float,
	stock_disponible int,
	stock_reservado int,
	punto_reposicion int,
	stock_maximo int
);

create table pedido(
	id_pedido serial,
	f_pedido date,
	fecha_entrega date,
	hora_entrega_desde time,
	hora_entrega_hasta time,
	id_usuarie int,
	id_direccion_entrega int,
	monto_total float,
	costo_envio float,
	estado char(10)
);


create table pedido_detalle(
	id_pedido int,
	id_producto int,
	cantidad int,
	precio_unitario float
);


create table reposicion(
	id_producto int,
	fecha_solicitud date,
	cantidad_a_reponer int,
	fecha_reposicion date,
	estado char(12)
);

create table error(
	id_error serial,
	id_pedido int,
	f_pedido timestamp,
	id_usuarie int,
	id_direccion_entrega int,
	direccion varchar(60),
	localidad varchar(30),
	codigo_postal char(8),
	id_producto int,
	cantidad int,
	operacion char(12),
	f_error timestamp default date_trunc('second', now()),
	motivo varchar(64)
);


create table envio_email(
	id_email serial,
	f_generacion timestamp,
	email_cliente varchar(64),
	asunto varchar(64),
	cuerpo varchar(64),
	f_envio timestamp,
	estado char(10)
);

create table entrada_trx_pedido(
	id_orden int,
	operacion char(12),
	id_usuarie int,
	id_direccion_entrega int,
	id_pedido int,
	id_producto int,
	cantidad int,
	fecha_hora_entrega timestamp
);
