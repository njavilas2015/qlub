package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Subdomain struct {
	Name       string   `json:"name"`
	Port       string   `json:"port"`
	Instances  []string `json:"instances"`
	HTTPS      bool     `json:"https"`
	SSLCert    string   `json:"ssl_cert"`
	SSLCertKey string   `json:"ssl_cert_key"`
}

func LoadConfig(filename *string) ([]Subdomain, error) {

	data, err := os.ReadFile(*filename)

	if err != nil {
		return nil, err
	}

	var subdomains []Subdomain

	if err := json.Unmarshal(data, &subdomains); err != nil {
		return nil, err
	}

	return subdomains, nil
}

func ParseFlags() (*bool, *string) {
	showVersion := flag.Bool("version", false, "Muestra la versión de qlub")
	configPath := flag.String("config", "", "Ruta al archivo de configuración JSON")

	flag.Parse()

	if *configPath == "" {
		fmt.Println("Por favor proporciona la ruta al archivo de configuración con --config.")
		os.Exit(1)
	}

	return showVersion, configPath
}

func CheckFileExistence(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
