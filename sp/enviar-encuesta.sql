create or replace function enviar_encuesta()
returns void as $$
begin
	-- Enviamos encuesta a quienes recibieron sus productos ayer
	insert into envio_email (
		f_generacion,
		email_cliente,
		asunto,
		cuerpo,
		estado)
	select
		now(),
		c.email,
		'Encuesta de satisfacción',
		'bit.ly/chaonl <- Completa y obten descuentos para la prox compra',
		'pendiente'
	from pedido p, cliente c
	where p.fecha_entrega = current_date - interval '1 day'
	and p.id_usuarie = c.id_usuarie;

end;
$$ language plpgsql;
