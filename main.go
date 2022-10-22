package main

import (
	"log"

	"github.com/JDHM12/twittor/bd"
	"github.com/JDHM12/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
