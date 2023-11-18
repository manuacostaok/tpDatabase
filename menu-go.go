package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"acosta_lombardi_quevedo_rumiz_db1/boltDB"
	"time"
)


func runSql(db *sql.DB, path string) {
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal(ioErr)
	}
	sql := string(c)
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

//funcion para activar isolation level en los SP serializables
func runSqlIls(db *sql.DB, path string) {
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal(ioErr)
	}
	sql := string(c)
	
	// se carga el sp en modo isolation level serializable
	_, err := db.Exec("BEGIN ISOLATION LEVEL SERIALIZABLE;"+sql+"COMMIT;")
	if err != nil {
		log.Fatal(err)
	}
}

//funcion para conectar a la base de datos postgres 
func connectDb() (db *sql.DB) {
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//funcion para conectar nuestra base de datos
func connectDbTP() (db *sql.DB) {
    db, err := sql.Open("postgres", "user=postgres host=localhost dbname=acosta_lombardi_quevedo_rumiz_db1 sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    return db
}

//funcion para ejecutar los SP diarios
func runDailyFunctions(db *sql.DB) {
	_, err := db.Exec("select generar_reposicion();")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("select enviar_encuesta();")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db := connectDb()
	dbtp := connectDbTP()

	var mostrarMenu = true
	var numeroIngresado string
	asciiArt :=

		`
   ___  _                              ___         _  _
  / __|| |_   __ _  _ _   __ _  ___   / _ \  _ _  | |(_) _ _   ___
 | (__ | ' \ / _  ||   \ / _  |/ _ \ | (_) ||   \ | || ||   \ / -_)
  \___||_||_|\__,_||_||_|\__, |\___/  \___/ |_||_||_||_||_||_|\___|
                         |___/                                     `

	menuDeOpciones := `
	Ingrese 1 si desea crear la base de datos
    	Ingrese 2 si desea agregar las tablas
    	Ingrese 3 si desea agregar pk y fk
    	Ingrese 4 si desea quitar pk y fk
     	Ingrese 5 si desea agregar datos
    	Ingrese 6 si desea ingresar los stored procedures y triggers
    	Ingrese 7 si desea ejecutar la base noSQL
        Ingrese 8 si desea cargar y ejecutar los datos de prueba
        Ingrese 9 si desea iniciar las tareas diarias
    	Ingrese 0 si desea salir`

	for mostrarMenu {


		fmt.Println(asciiArt)
		fmt.Println(menuDeOpciones)
		fmt.Scanf("%s", &numeroIngresado)

		if numeroIngresado == "1" {
			runSql(db, "dropear-base-de-datos.sql")
			fmt.Println("Base dropeada")
			runSql(db, "crear-base-de-datos.sql")
			fmt.Println("Base creada")
		}
		if numeroIngresado == "2" {
			db.Close()
			runSql(dbtp, "tablas.sql")
			fmt.Println("Tablas creadas con exito")
		}
		if numeroIngresado == "3" {
			runSql(dbtp, "agregar-pk-y-fk.sql")
			fmt.Println("Primary keys y foreign keys agregadas")
		}
		if numeroIngresado == "4" {
		    runSql(dbtp, "quitar-pk-fk.sql")
			fmt.Println("Se han quitado las primary keys y foreign keys")
		}
		if numeroIngresado == "5" {
			runSql(dbtp, "agregar-datos.sql")
			fmt.Println("Los datos han sido agregados")
		}
		if numeroIngresado == "6" {
			// funciones serializables 
			runSqlIls(dbtp,"sp/agregar-producto.sql")
			runSqlIls(dbtp,"sp/cancelar-pedido.sql")
			runSqlIls(dbtp,"sp/entrega-pedido.sql")
			//-------------
			runSql(dbtp,"sp/cierre-pedido.sql")
			runSql(dbtp,"sp/crear-pedido.sql")
			runSql(dbtp,"sp/email-a-clientes.sql")			
			runSql(dbtp,"sp/generar-reposicion.sql")
			runSql(dbtp,"sp/enviar-encuesta.sql")
			runSql(dbtp,"sp/ejecutar-pruebas.sql")
			fmt.Println("Stored Procedures y Triggers ejecutados")
		}
		if numeroIngresado == "7" {
			//acá uso el otro archivo .go que usa boltDB
			nosql.Menu()
		}
		if numeroIngresado == "8" {
			runSql(dbtp, "insertar-datos-en-pedidos.sql")
			fmt.Println("Datos cargados")
			_, err := dbtp.Exec("select ejecutar_pruebas();")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Pruebas ejecutadas")
		}
		if numeroIngresado == "9" {
			//se ejecuta en otro hilo ya que time sleep se queda ejecutando y se trababa en el for
			go func() {
				for {
					runDailyFunctions(dbtp)
					time.Sleep(24 * time.Hour)
				}
			}()
		}
		if numeroIngresado == "0" {
			dbtp.Close()
			mostrarMenu = false
			fmt.Println("chau")
		}
	}
}
