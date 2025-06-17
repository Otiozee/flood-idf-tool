// Copyright (c) 2025 Zedick Otokpa
// Licensed under the MIT License. See LICENSE in the project root for license information.
package calculator

import (
	"fmt"
	"os"
)

func RunCLI(calc *IDFCalculator) {
	fmt.Println("🌧️ Makurdi IDF Modeling Tool (CLI Mode)")
	fmt.Println("----------------------------------------")

	for {
		var duration, frequency float64
		fmt.Print("\nEnter Duration (hr): ")
		_, err := fmt.Scan(&duration)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		
		fmt.Print("Enter Return Period (yr): ")
		_, err = fmt.Scan(&frequency)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		intensity := calc.CalculateIntensity(duration, frequency, Gumbel)
		fmt.Printf("\n🚀 Calculated Rainfall Intensity: %.2f mm/hr\n", intensity)
		fmt.Println("Distribution: Gumbel (Optimal per research)")
		
		var choice string
		fmt.Print("\nCalculate another? (y/n): ")
		fmt.Scan(&choice)
		if choice != "y" {
			os.Exit(0)
		}
	}
}
