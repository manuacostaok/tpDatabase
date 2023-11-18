DROP FUNCTION IF EXISTS generar_reposicion();

create or replace function generar_reposicion() returns boolean as $$

declare
	p_loop record;	
	t_cantidad_a_reponer int;	
begin
	  -- corrobora que no haya un pedido de reposicion ya cargado en el dia.
	if exists (select 1 from reposicion
	where estado = 'pendiente' and fecha_solicitud = current_date)
	then 
		return false;
	end if; 
	-- recorre cada producto y corrobora que no este su stock por debajo del punto de reposicion.	
	for p_loop in select * from producto loop
		if p_loop.stock_disponible + p_loop.stock_reservado <= p_loop.punto_reposicion then 
			t_cantidad_a_reponer = p_loop.stock_maximo - (p_loop.stock_disponible + p_loop.stock_reservado);		
			-- se inserta el producto con la cantidad que se necesite reponer. 
			insert into reposicion(id_producto,fecha_solicitud,cantidad_a_reponer,estado)
			values (p_loop.id_producto, current_date,t_cantidad_a_reponer,'pendiente');		
		end if;	
	end loop;
	
	return true;
	
end;

$$ language plpgsql;
