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

	for _, subdomain := range subdomains {
		if err := config.GenerateNginxConfig(subdomain); err != nil {
			log.Printf("Error al generar configuración para %s: %v", subdomain.Name, err)
		}
	}
}
