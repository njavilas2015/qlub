package main

import (
	"fmt"
	"os"

	"github.com/njavilas2015/qlub/internal/config"
)

var version = "1.0.0"

func main() {

	showVersion, configPath, watcher := config.ParseFlags()

	if *showVersion {
		fmt.Printf("Versión: %s\n", version)
		os.Exit(0)
	}

	if *watcher {
		config.Watcher(configPath)
	} else {
		config.Generate(configPath)
	}

}
