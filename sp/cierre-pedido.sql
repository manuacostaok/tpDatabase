create or replace function cerrar_pedido(p_id_pedido int, p_fecha_entrega date, p_hora_entrega time)
returns boolean as $$
declare
	v_cantidad_productos int;
	v_hora_hasta time;
begin
 -- busca pedido a cerrar
	select count(*) into v_cantidad_productos
	from pedido_detalle
	where id_pedido = p_id_pedido;

 -- corrobora que existe pedido , en su defecto anota motivo en tabla error y corta ejecucion.
	if v_cantidad_productos = 0 then
		insert into error (id_pedido, operacion, motivo)
		values (p_id_pedido, 'cierre', '?id de pedido no válido');
		return false;
	end if;
 -- corrobora que el pedido entrega  de fecha y hora  no sea anterior a la fecha actual. en su defecto anota motivo en tabla error y corta ejecucion.
	if p_fecha_entrega < current_date or (p_fecha_entrega = current_date and p_hora_entrega <= current_time) then
		insert into error (id_pedido, operacion, motivo)
		values (p_id_pedido, 'cierre', '?fecha de entrega no válida');
		return false;
	end if;
 -- calculo de intervalo de la hora de entrega, desde y hasta
	v_hora_hasta := p_hora_entrega + interval '2 hours';

-- actualiza valores.
	update pedido
	set fecha_entrega = p_fecha_entrega,
		hora_entrega_desde = p_hora_entrega,
		hora_entrega_hasta = v_hora_hasta,
		estado = 'completado'
	where id_pedido = p_id_pedido;

	return true;
end;
$$ language plpgsql;

