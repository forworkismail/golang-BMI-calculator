package main

import (
	"fmt"

	"github.com/ismailwork/golang-BMI-calculator/info"
)

func main() {

	info.PrintWelcome()

	weight, height := getUserMetrics()

	bmi := calculateBMI(weight, height)

	fmt.Printf("Your BMI: %.2f", bmi)

}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}
