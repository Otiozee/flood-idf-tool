// Copyright (c) 2025 Zedick Otokpa
// Licensed under the MIT License. See LICENSE in the project root for license information.
package api

import (
	"encoding/json"
	"fmt" // Added fmt import
	"log" // Added log import
	"net/http"
	"strconv"

	"github.com/Otiozee/flood-idf-tool/pkg/calculator"
	"github.com/Otiozee/flood-idf-tool/pkg/nasa"
)

type Router struct {
	calc       *calculator.IDFCalculator
	nasaClient *nasa.GPMClient
}

func NewRouter(calc *calculator.IDFCalculator, nasaClient *nasa.GPMClient) *Router {
	return &Router{calc, nasaClient}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/idf":
		r.handleIDFCalculation(w, req)
	case "/nasa-integrate":
		r.handleNASAIntegration(w, req)
	case "/health":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	default:
		http.NotFound(w, req)
	}
}

func (r *Router) handleIDFCalculation(w http.ResponseWriter, req *http.Request) {
	duration, _ := strconv.ParseFloat(req.URL.Query().Get("d"), 64)
	returnPeriod, _ := strconv.ParseFloat(req.URL.Query().Get("t"), 64)

	intensity := r.calc.CalculateIntensity(duration, returnPeriod, r.calc.OptimalDistribution())

	response := map[string]interface{}{
		"intensity":    intensity,
		"duration":     duration,
		"returnPeriod": returnPeriod,
		"distribution": "Gumbel",
		"units":        "mm/hr",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (r *Router) handleNASAIntegration(w http.ResponseWriter, req *http.Request) {
	var request nasa.IntegrationRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	baseIntensity := r.calc.CalculateIntensity(request.Duration, request.ReturnPeriod, r.calc.OptimalDistribution())
	enhancedIntensity, pointData := r.nasaClient.EnhancedCalculation(request, baseIntensity)

	response := map[string]interface{}{
		"baseIntensity":     baseIntensity,
		"enhancedIntensity": enhancedIntensity,
		"adjustmentFactor":  enhancedIntensity / baseIntensity,
		"nasaDataPoints":    pointData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RunServer(port int, calc *calculator.IDFCalculator, nasaClient *nasa.GPMClient) {
	router := NewRouter(calc, nasaClient)
	addr := fmt.Sprintf(":%d", port)

	log.Printf("🚀 API Server running on http://localhost%s", addr)
	log.Printf("📊 Endpoints:")
	log.Printf("  GET  /idf?d=<duration>&t=<return_period>")
	log.Printf("  POST /nasa-integrate")
	log.Printf("  GET  /health")
	log.Fatal(http.ListenAndServe(addr, router))
}
