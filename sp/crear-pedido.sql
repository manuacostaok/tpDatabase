create or replace function crear_pedido(p_id_usuarie int, p_id_direccion_entrega int) returns boolean as $$
declare
	cp char(4);
begin
	-- validacion del id de usuario ingresado
	if p_id_usuarie not in (select id_usuarie from cliente) then
		insert into error(id_usuarie, operacion, motivo)
		values (p_id_usuarie, 'crear', '?id de usuarie no válido');
		return false;
	end if;
	-- se obtiene del codigo_postal solo los numeros para obtener codigo postal corto (cp).
	select substring(codigo_postal, 2, 4) into cp from direccion_entrega
	where id_usuarie = p_id_usuarie and id_direccion_entrega = p_id_direccion_entrega;
	-- valida  si el codigo postal de entrega ingresado es valido, caso negativo almacena en tabla error lo ocurrido
	if not found then
		insert into error (id_usuarie, id_direccion_entrega, operacion, motivo)
		values (p_id_usuarie, p_id_direccion_entrega, 'crear', '?id de dirección no válida');
		return false;
	end if;
	-- valida si el codigo postal esta dentro del area de atencion, caso negativo almacena en tabla error lo ocurrido
	if cp not in (select codigo_postal_corto from tarifa_entrega) then
		insert into error (id_direccion_entrega, operacion, motivo)
		values (p_id_direccion_entrega, 'crear', '?direccion de entrega fuera del área de atención');
		return false;
	end if;
	-- pasadas las validaciones anteriores, se insertan valores en tabla pedido
	insert into pedido(f_pedido,id_usuarie,id_direccion_entrega,monto_total,costo_envio,estado)
	values(
		current_date,
		p_id_usuarie,
		p_id_direccion_entrega,
		0,
		(select costo from tarifa_entrega where codigo_postal_corto = cp),
		'ingresado');
	return true;
end;
$$language plpgsql;
