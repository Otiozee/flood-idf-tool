// Copyright (c) 2025 Zedick Otokpa
// Licensed under the MIT License. See LICENSE in the project root for license information.
package nasa

import (
	"net/http"
	"time"
)

type GPMClient struct {
	apiKey string
	client *http.Client
}

func NewGPMClient(apiKey string) *GPMClient {
	return &GPMClient{
		apiKey: apiKey,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

type GPMPointData struct {
	Lat      float64   `json:"lat"`
	Lon      float64   `json:"lon"`
	Value    float64   `json:"value"`
	DateTime time.Time `json:"datetime"`
}

type IntegrationRequest struct {
	Region      string    `json:"region"`
	Duration    float64   `json:"duration"`
	ReturnPeriod float64   `json:"return_period"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

func (c *GPMClient) GetRainfallData(region string, start, end time.Time) ([]GPMPointData, error) {
	// In production: Call NASA GPM API
	// url := fmt.Sprintf("https://gpm.nasa.gov/api/data?region=%s&start=%s&end=%s&apikey=%s", 
	//     region, start.Format(time.RFC3339), end.Format(time.RFC3339), c.apiKey)
	
	// Mock response for demonstration
	return []GPMPointData{
		{Lat: 7.7322, Lon: 8.5391, Value: 15.4, DateTime: time.Now().Add(-24 * time.Hour)},
		{Lat: 7.7350, Lon: 8.5410, Value: 17.2, DateTime: time.Now().Add(-12 * time.Hour)},
	}, nil
}

func (c *GPMClient) EnhancedCalculation(req IntegrationRequest, baseIntensity float64) (float64, []GPMPointData) {
	data, _ := c.GetRainfallData(req.Region, req.StartDate, req.EndDate)
	
	// Calculate adjustment factor
	adjustment := 1.0
	if len(data) > 0 {
		total := 0.0
		for _, d := range data {
			total += d.Value
		}
		avg := total / float64(len(data))
		adjustment = 1 + (avg-16.0)/100
	}
	
	return baseIntensity * adjustment, data
}
