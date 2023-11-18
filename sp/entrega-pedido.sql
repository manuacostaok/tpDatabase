create or replace function entrega_pedido(p_id_pedido int) returns boolean as $$

declare
	resultado pedido%rowtype;

begin
 -- busca pedido
	select * into resultado
		from pedido
		where id_pedido = p_id_pedido;
-- corrobora que exista pedido , y guarda motivo en tabla error y corta la ejecucion con false.
	if not found then
		insert into error(id_pedido,motivo)
		values(p_id_pedido,'id de pedido no válido');
		return false;
	end if;
-- corrobora si el producto esta cerrado o ya fue entregado. guarda motivo y corta ejecucion.
	if resultado.estado != 'completado' then
		insert into error(id_pedido,motivo)
		values(p_id_pedido,'pedido sin cerrar o ya entregado');
		return false;
	end if;
-- actualiza estado de pedido.
	update pedido
		set estado = 'entregado'
		where id_pedido = p_id_pedido;
--actualiza valores del stock reservado por cada cantidad de producto del pedido entregado.
	update producto
		set stock_reservado = stock_reservado - pd.cantidad
		from pedido_detalle pd
		where pd.id_pedido = p_id_pedido
		and pd.id_producto = producto.id_producto;

	return true;

end;

$$language plpgsql;
