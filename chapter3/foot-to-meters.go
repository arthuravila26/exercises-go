//1 pé = 0,3048m

package main

import "fmt"

func main() {
	fmt.Print("Enter the foot to convert: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println(output, "m")
}
