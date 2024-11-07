package config

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func write(fileName string, name string, text string, data any) {

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalf("Error creating nginx.conf: %v", err)
	}

	defer file.Close()

	rawTemplate, err := template.New(name).Parse(text)

	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	if err := rawTemplate.Execute(file, data); err != nil {
		log.Fatalf("Error generating config: %v", err)
	}

	fmt.Println("NGINX config generated successfully.")
}

func GenerateNginxConfig(subdomains []Subdomain) error {

	write("nginx.conf", "nginx", DefaultNginxTemplate, "")

	os.RemoveAll("conf.d")

	err := os.MkdirAll("conf.d", os.ModePerm)

	if err != nil {
		log.Fatalf("Error creating directory: %v", err)
	}

	for index, subdomain := range subdomains {

		fileName := fmt.Sprintf("conf.d/%v.conf", index+1)

		write(fileName, "nginx", NginxTemplate, subdomain)

	}

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
