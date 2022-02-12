package main

import (
	"fmt"

	"github.com/ismailwork/golang-BMI-calculator/info"
)

func main() {

	info.PrintWelcome()

	weight, height := getUserMetrics()

	bmi := weight / (height * height)

	fmt.Printf("Your BMI: %.2f", bmi)

}
