create or replace function ejecutar_pruebas() returns void as $$
declare
	v entrada_trx_pedido%rowtype;
begin
	-- recorre con un bucle for las transacciones-pedidas y segun los casos que puede estar el valor de la operacion ejecuta la funcion adecuada.
	for v in select * from entrada_trx_pedido loop
		raise notice 'Corriendo operación: %', v.operacion;
		case v.operacion
			when 'creación' then
				perform crear_pedido(v.id_usuarie, v.id_direccion_entrega);
			when 'producto' then
				perform agregar_producto_al_pedido(v.id_pedido, v.id_producto, v.cantidad);
			when 'cierre' then
				perform cerrar_pedido(v.id_pedido, v.fecha_hora_entrega::date, v.fecha_hora_entrega::time);
			when 'entrega' then
				perform entrega_pedido(v.id_pedido);
			when 'cancelación' then
				perform cancelar_pedido(v.id_pedido);
			else
				raise exception 'operación no válida: %', v.operacion;
		end case;
	end loop;
end;
$$ language plpgsql;
