create or replace function enviar_email_a_clientes() returns trigger as $$
declare
	asunto char(16);
begin
	--Elejimos el asunto segun el estado
	case new.estado
		when 'completado' then
			asunto := 'Pedido aceptado';
		when 'cancelado' then
			asunto := 'Pedido cancelado';
		when 'entregado' then
			asunto := 'Pedido entregado';
	end case;
	insert into envio_email (
		f_generacion,
		email_cliente,
		asunto,
		cuerpo,
		estado
	)
	values (
		now(),
		(select email from cliente where id_usuarie = new.id_usuarie),
		asunto,
		--Generamos un string del tipo 'Ped 1 16 x2 $1450cu,  14 x4 $200cu'
		--Y limitamos la cantidad de caracteres a 64 para q entre en char(64)
		substring('Ped ' || new.id_pedido ||
			(select string_agg(
			' ' || id_producto ||
			' x' || cantidad ||
			' $' || precio_unitario || 'cu', ',')
			from pedido_detalle
			where id_pedido = new.id_pedido),1,64),
		'pendiente'
	);
	return new;
end;
$$ language plpgsql;

drop trigger if exists enviar_email_trigger on pedido;
create trigger enviar_email_trigger
after update on pedido
for each row
--Solo nos interesa cuando cambia el estado y no el monto
when (old.estado <> new.estado)
execute procedure enviar_email_a_clientes();
