package main

import (
	"fmt"
)

type archivo struct {
	ID        int    "json:id"
	Nombre    string "json:nombre"
	Extension string "json:ext"
}

type listaArchivos []archivo

var lista = listaArchivos{
	{
		ID: 0, Nombre: "archivo1", Extension: ".prueba",
	},
}

func main() {

	fmt.Println("Corriendo...")

}
