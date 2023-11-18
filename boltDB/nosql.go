package nosql

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
)

var (
	dbbolt            	*bolt.DB
	err               	error
	dataProductos  	    json.RawMessage
	dataClientes 		json.RawMessage
	dataDirecciones 	json.RawMessage
	dataPedidos			json.RawMessage
	dataPedidos_detalle	json.RawMessage
)

type Cliente struct {
	id_usuarie int
	nombre string
	apellido string
	dni int 
	fecha_nacimento string
	telefono string
	email string
}

type Direccion_entrega struct{
	id_usuarie int
	id_direccion_entrega int
	direccion string
	localidad string
	codigo_postal string
}

type Producto struct{
	id_producto int
	nombre string
	precio_unitario float32
	stock_disponible int
	stock_reservado int
	punto_reposicion int
	stock_maximo int
}

type Pedido struct{
	id_pedido int
	f_pedido string
	fecha_entrega string
	hora_entrega_desde string
	hora_entrega_hasta string
	id_usuarie int
	id_direccion_entrega int
	monto_total float32
	costo_envio float32
	estado string
}

type Pedido_detalle struct{
	id_pedido int
	id_producto int
	cantidad int
	precio_unitario float32
}

func Menu(){
	
	var mostrarMenu = true
	var numeroIngresado string
	asciiArt :=

		`
   ___  _                              ___         _  _
  / __|| |_   __ _  _ _   __ _  ___   / _ \  _ _  | |(_) _ _   ___
 | (__ | ' \ / _  ||   \ / _  |/ _ \ | (_) ||   \ | || ||   \ / -_)
  \___||_||_|\__,_||_||_|\__, |\___/  \___/ |_||_||_||_||_||_|\___|
                         |___/         /BOLT_DB_VERSION/                `

	menuDeOpciones := `
	Ingrese 1 si desea generar el archivo JSON
    	Ingrese 2 si desea cargar los datos de prueba
    	Ingrese 3 si desea salir`
 
	for mostrarMenu {
		fmt.Println(asciiArt)
		fmt.Println(menuDeOpciones)
		fmt.Scanf("%s", &numeroIngresado)

		if numeroIngresado == "1" {
			GenerarJson()
			fmt.Println("JSON generado")	
		}
		if numeroIngresado == "2" {
			CargarDatos()
			fmt.Println("Datos cargados")
		}
		
		if numeroIngresado == "3" {
			mostrarMenu = false
			fmt.Println("Saliste de Chango Online")
		}
	}
}

func dbBoltConnection() {
	dbbolt, err = bolt.Open("nosql.db", 0600, nil)
	if err != nil {
		log.Fatalf("Error al conectar a db: %s", err)
	}
}

func CreateUpdate(dbbolt *bolt.DB, bucketName string, key []byte, val []byte) error {
	// abre transaccion de escritura
	tx, err := dbbolt.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
	if err != nil {
		return err
	}

	err = b.Put(key, val)
	if err != nil {
		return err
	}

	// cierra transaccion
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func ReadUnique(dbbolt *bolt.DB, bucketName string, key []byte) ([]byte, error) {
	var buf []byte

	err := dbbolt.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})

	return buf, err
}

func GenerarJson(){
	dbBoltConnection()
	
	Clientes:= []Cliente{
		{
			1,
			"Ken",
			"Thompson",
			5153057,
			"1995-05-05",
			"15-2889-7948",
			"ken@thompson.org"},
		{
			2,
			"Dennis",
			"Ritchie",
			25610126,
			"1955-04-11",
			"15-7811-5045",
			"dennis@ritchie.org"},
		{
			3,
			"Donald",
			"Knuth",
			9168297,
			"1984-04-05",
			"15-2780-6005",
			"don@knuth.org"},
		{
			4,
			"Rob",
			"Pike",
			4915593,
			"1946-08-16",
			"15-1114-9719",
			"rob@pike.org"},
		{
			5,
			"Douglas",
			"McIlroy",
			33187055,
			"1939-06-09",
			"15-9625-0245",
			"douglas@mcilroy.org"},
		{
			6,
			"Brian",
			"Kernighan",
			13897948,
			"1992-11-22",
			"15-6410-6066",
			"brian@kernighan.org"},
		{
			7,
			"Bill",
			"Joy",
			34115045,
			"1954-02-04",
			"15-4215-8655",
			"bill@joy.org"},
		{
			8,
			"Marshall Kirk",
			"McKusick",
			9806005,
			"1995-12-27",
			"15-5197-4379",
			"marshall_kirk@mckusick.org"},
		{
			9,
			"Theo",
			"de Raadt",
			5149719,
			"1950-02-07",
			"15-6470-9444",
			"theo@deraadt.org"},
		{
			10,
			"Cristina",
			"Kirchner",
			6250245,
			"1990-08-17",
			"15-5291-0113",
			"cfk@fpv.gov.ar"},
		{
			11,
			"Diego",
			"Maradona",
			19158655,
			"1985-02-27",
			"15-3361-4854",
			"diego@dios.com.ar"},
		{
			12,
			"Martín",
			"Palermo",
			5974379,
			"1918-06-09",
			"15-9877-3169",
			"martin@palermo.com.ar"},
		{
			13,
			"Guillermo",
			"Barros Schelotto",
			3910113,
			"1982-05-03",
			"15-5020-5695",
			"guille@melli.com.ar"},
		{
			14,
			"Susú",
			"Pecoraro",
			7547862,
			"1935-04-03",
			"15-6695-9505",
			"susu@pecoraro.com.ar"},
		{
			15,
			"Norma",
			"Aleandro",
			26614854,
			"1992-03-18",
			"15-9155-4115",
			"norma@aleandro.com.ar"},
		{
			16,
			"Soledad",
			"Silveyra",
			7773169,
			"1957-07-28",
			"15-9184-4522",
			"sole@silveyra.com.ar"},
		{
			7,
			"Libertad",
			"Lamarque",
			32205695,
			"1971-03-07",
			"15-6363-9690",
			"libertad@lamarque.com.ar"},
		{
			18,
			"Ana María",
			"Picchio",
			19020903,
			"1946-08-06",
			"15-4819-2117",
			"ana.maria@picchio.com.ar"},
		{
			19,
			"Niní",
			"Marshall",
			10535508,
			"1951-09-07",
			"15-9799-6045",
			"nini@marshall.com"},
		{
			20,
			"Claudia",
			"Lapacó",
			30934609,
			"1961-08-03",
			"15-2005-4879",
			"claudia@lapaco.com.ar"},
		
	}
	dataClientes, err = json.MarshalIndent(Clientes, "", "    ")
	if err != nil {
		log.Fatalf("JSON Error: %s", err)
	}
	
	Direcciones:= []Direccion_entrega{
		{
			1,
			1,
			"Av. Rivadavia 3855, 4A",
			"Almagro",
			"B1636FDA"},
		{
			2,
			1,
			"Av. Las Heras 3215, 2B",
			"Palermo",
			"B7721GBF"},
		{
			3,
			1,
			"Av. Antártida Arg. 703, 5C",
			"Lomas de Zamora",
			"B1821HKD"},
		{
			 4,
			 1,
			"Av. Centenario 72, 1D",
			 "San Isidro",
			 "B6766XJI"},
		{
			 5,
			 1,
			 "Ruta 8 3202, PB",
			 "San Miguel",
			 "B5823AEB"},
		{
			 6,
			 1,
			 "Av. Don Bosco 2680, 1B",
			 "San Justo",
			 "B8591ARU"},
		{
			 7,
			 1,
			 "Av. Santa Fé 2468, 8B",
			 "Recoleta",
			 "B6592ZCH"},
		{
			 7,
			 2,
			 "Av. Corrientes 2558, 3B",
			 "Balvanera",
			 "B2403JAJ"},
		{
			 8,
			 1,
			 "Av. Colón 394, 4B",
			 "Mendoza",
			 "B2510BTM"},
		{
			 9,
			 1,
			 "Perón 989, 2B",
			 "San Miguel",
			 "B8294CXB"},
		{
			 10,
			 1,
			 "Av. Constitución 7570, 5E",
			 "Mar del Plata",
			 "B3332EKY"},
		{
			 11,
			 1,
			 "Av. Rivadavia 5730, 6A",
			 "Caballito",
			 "B8142QWD"},
		{
			 11,
			 2,
			 "Colombres 92, 3D",
			 "Avellaneda",
			 "B1390CPB"},
		{
			 12,
			 1,
			 "Av. Vergara 1900, 1A",
			 "Morón",
			 "B9419XPV"},
		{
			 13,
			 1,
			 "Alvear 2486, 3A",
			 "Villa Ballester",
			 "B5579ABM"},
		{
			 14,
			 1,
			 "H. Yrigoyen 13200, 3C",
			 "Adrogué",
			 "B5030WUE"},
		{
			 15,
			 1,
			 "Brown 266, 4C",
			 "Bahía Blanca",
			 "B5047JAX"},
		{
			 16,
			 1,
			 "Av. Cabildo 2254, 3B",
			 "Belgrano",
			 "B1390CPB"},
		{
			 17,
			 1,
			 "Av. O Higgins 332, 2B",
			 "Córdoba",
			 "B1515DLS"},
		{
			 18,
			 1,
			 "Gral. Urquiza 4785, PB",
			 "Caseros",
			 "B1663QLO"},
		{
			 19,
			 1,
			 "Av. Scalabrini Ortiz 3149, 5B",
			 "Palermo",
			 "B1613ZZT"},
		{
			20,
			 1,
			 "Av. Santa Rosa 62, 4C",
			 "Castelar",
			 "B1615QLO"},
	}
	
	dataDirecciones, err = json.MarshalIndent(Direcciones, "", "    ")
	if err != nil {
		log.Fatalf("JSON Error: %s", err)
	}
	
	Productos:=[]Producto{
		{
			 1,
			 "Shampoo Liso Perfecto Swing x 600 ml",
			 1200,
			 10,
			 0,
			 5,
			 50},
		{
			 2,
			 "Leche Chocolatada Chochoco x 1000 ml",
			 900,
			 15,
			 0,
			 5,
			 50},
		{
			 3,
			 "Salchichas Alemanas Patyviena x 4 un",
			 750,
			 20,
			 0,
			 10,
			 60},
		{
			 4,
			 "Helado DDL, Frutilla y Chocolate Noel x 1 kg",
			 2500,
			 5,
			 0,
			 5,
			 20},
		{
			 5,
			 "Café Molido Torrado al Grano de Café Cabrales x 500 gr",
			 1600,
			 12,
			 0,
			 5,
			 30},
		{
			 6,
			 "Cereales Aritos Frutales Granix x 130 gr",
			 500,
			 18,
			 0,
			 5,
			 30},
		{
			 7,
			 "Azúcar Común Tipo A SIC x 1 kg",
			 700,
			 25,
			 0,
			 10,
			 100},
		{
			 8,
			 "Ñoquis De Papa Aliada x 1 kg",
			 800,
			 8,
			 0,
			 5,
			 30},
		{
			 9,
			 "Harina de Trigo 0/3 Seleccionada x 5 kg",
			 1000,
			 30,
			 0,
			 15,
			 100},
		{
			 10,
			 "Yerba Mate Tradicional Clasica Piporé x 500 gr",
			 950,
			 32,
			 0,
			 10,
			 50},
		{
			 11,
			 "Máquina Para Afeitar Turbo Mach3 Gillette x 4 un",
			 3200,
			 4,
			 0,
			 2,
			 10},
		{
			 12,
			 "Jabón en Polvo Lavado a Mano Granby x 3 kg",
			 1950,
			 23,
			 0,
			 10,
			 50},
		{
			 13,
			 "Queso Cremoso Cuartirolo Reducido c/s trozado Verónica x 1 kg",
			 2000,
			 10,
			 0,
			 5,
			 20},
		{
			 14,
			 "Jugo Listo Naranja Ades x 200 ml",
			 200,
			 15,
			 0,
			 5,
			 50},
		{
			 15,
			 "Rollo de Cocina Flores 60 paños Elite x 3 un",
			 650,
			 40,
			 0,
			 20,
			 100},
		{
			 16,
			 "Vino Tinto Blend Dolores x 750 ml",
			 1450,
			 6,
			 0,
			 5,
			 20},
		{
			 17,
			 "Miel Frasco Taeq x 500 gr",
			 1300,
			 12,
			 0,
			 5,
			 20},
		{
			 18,
			 "Raviol x 1 kilo 4 quesos Mendia x 1 kg",
			 1750,
			 20,
			 0,
			 5,
			 30},
		{
			 19,
			 "Queso Tybo Horma Lactovita x 1 kg",
			 2700,
			 10,
			 0,
			 5,
			 30},
		{
			 20,
			 "Café Molido Torrado Bonafide x 1 kg",
			 3400,
			 4,
			 0,
			 2,
			 10},
	}
	
	dataProductos, err = json.MarshalIndent(Productos, "", "    ")
	if err != nil {
		log.Fatalf("JSON Error: %s", err)
	}
	
	Pedidos:=[]Pedido{
		{
			1,
			"2018-09-12", 
			"2018-09-12",
			"20:18:09:12", 
			"20:18:09:12",
			1,
			8,
			20180,
			912,
			"entregado"},
		{
			2,
			"2018-09-23", 
			"2018-09-24",
			"20:18:09:12", 
			"20:18:09:12",
			2,
			4,
			2018,
			912,
			"cancelado"},
		{
			3,
			"2019-09-01", 
			"2019-09-02",
			"20:18:09:12", 
			"20:18:09:12",
			3,
			6,
			20180,
			912,
			"aceptado"},
		{
			4,
			"2021-01-02", 
			"2021-01-03",
			"20:18:09:12", 
			"20:18:09:12",
			4,
			5,
			20180,
			912,
			"aceptado"},
		{
			5,
			"2021-03-14", 
			"2021-03-15",
			"20:18:09:12", 
			"20:18:09:12",
			5,
			5,
			20180,
			912,
			"entregado"},
	}
	dataPedidos, err = json.MarshalIndent(Pedidos, "", "    ")
	if err != nil {
		log.Fatalf("JSON Error: %s", err)
	}
	
	Pedidos_detalle:=[]Pedido_detalle{
			{1,2,3,9.12},
			{2,5,5,20.18},
			{3,3,4,91.2},
			{3,11,7,201.8912},
			{4,1,11,912.2},
			{4,8,7,2018.2},
			{4,19,6,9.12},
			{4,20,16,2018.2},
			{3,6,4,9.122018},
			{5,8,15,2018.912},
			{5,9,12,912.2018},
			{5,1,10,91.22018},
			{2,15,15,9120.18},
		}
	dataPedidos_detalle, err = json.MarshalIndent(Pedidos_detalle, "", "    ")
	if err != nil {
		log.Fatalf("JSON Error: %s", err)
	}
	defer dbbolt.Close()
}

func CargarDatos() {
	dbBoltConnection()
	
	CreateUpdate(dbbolt, "Clientes", []byte("Clientes"), dataClientes)
	if err != nil {
		log.Fatalf("Cargar Clientes error: %s", err)
	}

	CreateUpdate(dbbolt, "Direcciones", []byte("Direcciones"), dataDirecciones)
	if err != nil {
		log.Fatalf("Cargar Direcciones error: %s", err)
	}

	CreateUpdate(dbbolt, "Productos", []byte("Productos"), dataProductos)
	if err != nil {
		log.Fatalf("Cargar Productos error: %s", err)
	}

	CreateUpdate(dbbolt, "Pedidos", []byte("Pedidos"), dataPedidos)
	if err != nil {
		log.Fatalf("Cargar Pedidos error: %s", err)
	}
	
	CreateUpdate(dbbolt, "Pedidos_detalle", []byte("PedidosDetalle"), dataPedidos_detalle)
	if err != nil {
		log.Fatalf("Cargar Pedidos Detalle error: %s", err)
	}

	defer dbbolt.Close()
}
	
	
