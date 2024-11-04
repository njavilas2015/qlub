package main

import (
	"log"

	"github.com/njavilas2015/qlub/internal/config"
)

func main() {

	subdomains, err := config.LoadConfig("subdomains.json")

	if err != nil {
		log.Fatalf("Error al cargar la configuración: %v", err)
	}

	if err := config.GenerateNginxConfig(subdomains); err != nil {
		log.Fatalf("Error generating config: %v", err)
	}
}
