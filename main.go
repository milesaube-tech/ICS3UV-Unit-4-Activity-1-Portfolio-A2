/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-05
 * @fileoverview This program prints all even numbers between a starting and ending odd number.
 */

package main

import (
	"fmt"
)

func main() {
  
	// get user input
	var startValue int
	var endValue int

	fmt.Print("Please enter a starting value (it must be an odd number): ")
	fmt.Scan(&startValue)

	fmt.Print("Please enter an ending value (it must be an odd number): ")
	fmt.Scan(&endValue)

	// set variables
	counter := 0
	output := ""

	//set up loop
	for counter = startValue + 1; counter <= endValue; counter = counter +2 {
		output = output + fmt.Sprint(" ", counter)
	}

	// output the results
	fmt.Println(output)

}
