package config

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func GenerateNginxConfig(subdomains []Subdomain) error {
	file, err := os.Create("nginx.conf")

	if err != nil {
		log.Fatalf("Error creating nginx.conf: %v", err)
	}

	defer file.Close()

	rawTemplate, err := template.New("nginx").Parse(NginxTemplate)

	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	if err := rawTemplate.Execute(file, subdomains); err != nil {
		log.Fatalf("Error generating config: %v", err)
	}

	fmt.Println("NGINX config generated successfully.")

	return nil
}

func Generate(configPath *string) {

	if CheckFileExistence(*configPath) {

		subdomains, err := LoadConfig(configPath)

		if err != nil {
			log.Fatalf("Error loading configuration: %v", err)
		}

		if err := GenerateNginxConfig(subdomains); err != nil {
			log.Fatalf("Error generating config: %v", err)
		}
	} else {
		fmt.Println("Please provide the path to the configuration file with --config.")
		os.Exit(1)
	}
}
