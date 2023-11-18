create or replace function agregar_producto_al_pedido(
	p_id_pedido int,
	p_id_producto int,
	p_cantidad int
) returns boolean as $$
declare
	v_estado_pedido char(10);
	v_stock_disponible int;
	v_monto_total_pedido decimal(12, 2);
begin
 -- busca el estado del pedido
	select estado into v_estado_pedido
	from pedido
	where id_pedido = p_id_pedido;

 -- corrobora que el pedido buscado exista , en su defecto guarda motivo en tabla de error y corta
	if v_estado_pedido is null then
		insert into error (id_pedido, operacion, motivo)
		values (p_id_pedido, 'producto', '?id de pedido no válido');
		return false;
	end if;

 -- corrobora que el pedido buscado este en estado ingresado, en su defecto guarda motivo en tabla de error y corta ejecucion
	if v_estado_pedido != 'ingresado' then
		insert into error (id_pedido, operacion, motivo)
		values (p_id_pedido, 'producto', '?pedido cerrado');
		return false;
	end if;

 -- corrobora que el  producto a inpcorporar al pedido exista, en su defecto guarda motivo en tabla de error y corta ejecucion
	if not exists (select 1 from producto where id_producto = p_id_producto) then
		insert into error (id_pedido, id_producto, operacion, motivo)
		values (p_id_pedido, p_id_producto, 'producto', '?id de producto no válido');
		return false;
	end if;
 -- almacena stock disponible actual del producto en v_stock_disponible
	select stock_disponible into v_stock_disponible
	from producto
	where id_producto = p_id_producto;

 -- corrobora que haya stock disponible para la cantidad solicitada , en su defecto guarda motivo en tabla error y corta ejecucion.
	if v_stock_disponible < p_cantidad then
		insert into error (id_pedido, id_producto, cantidad, operacion, motivo)
		values (p_id_pedido, p_id_producto, p_cantidad, 'producto', '?stock no disponible para el producto ');
		return false;
	end if;
  -- inserta los valores del producto en el pedido en la tabla pedido_detalle ,
  -- si ocurre un conflicto de pks ya existentes se realiza una actualizacion del valor de la cantidad unicamente.
	insert into pedido_detalle (id_pedido, id_producto, cantidad, precio_unitario)
	values (p_id_pedido, p_id_producto, p_cantidad, (select precio_unitario from producto where id_producto = p_id_producto))
	on conflict (id_pedido, id_producto) do update
	set cantidad = pedido_detalle.cantidad + excluded.cantidad;

-- actualiza valores del stock restanto la cantidad pedida e incorporandolo al stock reservado
	update producto
	set stock_disponible = stock_disponible - p_cantidad,
		stock_reservado = stock_reservado + p_cantidad
	where id_producto = p_id_producto;

-- calculo de monto total del producto solicitado
	select sum(cantidad * precio_unitario) into v_monto_total_pedido
	from pedido_detalle
	where id_pedido = p_id_pedido;
-- actualiza valor del monto total del pedido
	update pedido
	set monto_total = v_monto_total_pedido
	where id_pedido = p_id_pedido;

	return true;
end;
$$ language plpgsql;
