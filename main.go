package main

import (
	"fmt"
	"log"
	"os"

	"github.com/njavilas2015/qlub/internal/config"
)

var version = "1.0.0"

func main() {

	showVersion, configPath := config.ParseFlags()

	if *showVersion {
		fmt.Printf("Versión: %s\n", version)
		os.Exit(0)
	}

	if config.CheckFileExistence(*configPath) {

		subdomains, err := config.LoadConfig(configPath)

		if err != nil {
			log.Fatalf("Error al cargar la configuración: %v", err)
		}

		if err := config.GenerateNginxConfig(subdomains); err != nil {
			log.Fatalf("Error generating config: %v", err)
		}
	} else {
		fmt.Println("Por favor proporciona la ruta al archivo de configuración con --config.")
		os.Exit(1)
	}
}
