create or replace function cancelar_pedido(p_id_pedido int) returns boolean as $$

declare
	estado_pedido char(10);

begin
 -- busca el estado del pedido
	select estado into estado_pedido
	from pedido
	where id_pedido = p_id_pedido;
-- corrobora que exista el pedido
	if estado_pedido is null then
		insert into error (id_pedido, operacion, motivo)
		values (p_id_pedido, 'cancelacion', '?id de pedido no válido');
		return false;
	end if;
-- corrobora que no se pueda cancelar un producto ya entregado o cancelado
	if estado_pedido != 'ingresado' and estado_pedido != 'completado' then
		insert into error (id_pedido, operacion, motivo)
		values (p_id_pedido, 'cancelacion','?pedido ya entregado o cancelado');
		return false;
	end if;
-- pasado las validaciones anteriores se cambia el estado del pedido a cancelado
	update pedido
	set estado = 'cancelado'
	where id_pedido = p_id_pedido;
-- se restaura el stock disponible sumandole el stock reservado del pedido recien cancelado
	update producto
	set stock_disponible = stock_disponible + pd.cantidad,
		stock_reservado = stock_reservado - pd.cantidad
	from pedido_detalle pd
	where pd.id_pedido = p_id_pedido
	and pd.id_producto = producto.id_producto;

	return true;
end;
$$ language plpgsql;

