//Fahrenheit em Celsius
//C = ((F - 32) * 5/9)

package main

import "fmt"

func main() {
	fmt.Print("Enter the temperature(Fahrenheit): ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Println(output, "C°")
}
