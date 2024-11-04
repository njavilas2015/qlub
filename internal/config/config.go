package config

import (
	"encoding/json"
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

func LoadConfig(filename string) ([]Subdomain, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	var subdomains []Subdomain

	if err := json.Unmarshal(data, &subdomains); err != nil {
		return nil, err
	}

	return subdomains, nil
}
