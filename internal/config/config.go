package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
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

	if len(data) > 0 {

		if err := json.Unmarshal(data, &subdomains); err != nil {
			return nil, err
		}
	}

	return subdomains, nil
}

func ParseFlags() (*bool, *string, *bool) {
	showVersion := flag.Bool("version", false, "Show qlub version")
	configPath := flag.String("config", "", "Path to JSON configuration file")
	watcher := flag.Bool("watch", false, "Update configuration realtime")

	flag.Parse()

	if *configPath == "" {
		fmt.Println("Please provide the path to the configuration file with --config")
		os.Exit(1)
	}

	return showVersion, configPath, watcher
}

func CheckFileExistence(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func Watcher(filename *string) {

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	err = watcher.Add(*filename)

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {

					Generate(filename)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Error: %v\n", err)
			}
		}
	}()

	fmt.Println("Observing changes...")
	<-done
}
