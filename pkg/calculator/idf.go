// Copyright (c) 2025 Zedick Otokpa
// Licensed under the MIT License. See LICENSE in the project root for license information.
package calculator

import (
	"fmt"
	"math"
)

const (
	Gumbel = iota
	PearsonIII
	LogPearsonIII
)

type IDFConstants struct {
	C, M, E float64
}

var distributions = map[int]IDFConstants{
	Gumbel:       {404.11, 0.1733, 0.66676},
	PearsonIII:   {398.93, 0.1673, 0.6667},
	LogPearsonIII: {399.87, 0.1701, 0.6667},
}

type IDFCalculator struct{}

func NewIDFCalculator() *IDFCalculator {
	return &IDFCalculator{}
}

func (c *IDFCalculator) CalculateIntensity(duration, returnPeriod float64, distType int) float64 {
	constants, exists := distributions[distType]
	if !exists {
		constants = distributions[Gumbel]
	}
	return constants.C * math.Pow(returnPeriod, constants.M) / math.Pow(duration+constants.E, 0.6667)
}

func (c *IDFCalculator) OptimalDistribution() int {
	return Gumbel
}

// Add RunCLI implementation here
func RunCLI(calc *IDFCalculator) {
	fmt.Println("🌧️ Makurdi IDF Modeling Tool (CLI Mode)")
	fmt.Println("----------------------------------------")

	var duration, frequency float64
	fmt.Print("Enter Duration (hr): ")
	fmt.Scan(&duration)
	fmt.Print("Enter Return Period (yr): ")
	fmt.Scan(&frequency)

	intensity := calc.CalculateIntensity(duration, frequency, Gumbel)
	fmt.Printf("\n🚀 Calculated Rainfall Intensity: %.2f mm/hr\n", intensity)
	fmt.Println("Distribution: Gumbel (Optimal per research)")
}
