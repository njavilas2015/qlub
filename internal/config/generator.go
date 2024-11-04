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
