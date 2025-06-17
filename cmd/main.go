// Copyright (c) 2025 Zedick Otokpa
// Licensed under the MIT License. See LICENSE in the project root for license information.
package main

import (
	"flag"
	"os"

	"github.com/Otiozee/flood-idf-tool/pkg/calculator"
	"github.com/Otiozee/flood-idf-tool/pkg/nasa" // Re-added nasa import
)

func main() {
	cliMode := flag.Bool("cli", false, "Run in command-line mode")
	flag.Parse()

	if *cliMode {
		idfCalc := calculator.NewIDFCalculator()
		calculator.RunCLI(idfCalc)
	} else {
		apiPort := flag.Int("port", 8080, "API server port")
		nasaClient := nasa.NewGPMClient(os.Getenv("NASA_API_KEY"))
		
		// Temporary to avoid unused variable errors
		_ = apiPort
		_ = nasaClient
	}
}
