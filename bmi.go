package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ismailwork/golang-BMI-calculator/info"
)

func main() {

	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)
	fmt.Print(info.WeightPrompt)
	weightInput, _ := reader.ReadString('\n')

	fmt.Print(info.HeightPrompt)
	heightInput, _ := reader.ReadString('\n')

	weightInput = strings.Replace(weightInput, "\r\n", "", -1)
	heightInput = strings.Replace(heightInput, "\r\n", "", -1)

	weight, _ := strconv.ParseFloat(weightInput, 64)
	height, _ := strconv.ParseFloat(heightInput, 64)

	bmi := weight / (height * height)

	fmt.Printf("Your BMI: %.2f", bmi)

}
