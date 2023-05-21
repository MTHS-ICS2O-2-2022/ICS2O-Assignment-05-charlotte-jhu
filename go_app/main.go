// Created by: Charlotte Jhu
// Created on: May 2023
//
// This program calculates the hailstone sequence

package main

import (
	"fmt"
)

func main() {
	// This function calculates the hailstone sequence
	// variables
	var positiveInteger int

	// input
	fmt.Println("This program calculates the hailstone sequence")
	fmt.Println()
	fmt.Print("Enter a positive integer: ")
	fmt.Scanln(&positiveInteger)
	fmt.Println()

	// process
	for positiveInteger > 1 {
		if positiveInteger%2 == 0 {
			positiveInteger = positiveInteger / 2
			fmt.Println(positiveInteger)
		} else {
			positiveInteger = positiveInteger*3 + 1
			fmt.Println(positiveInteger)
		}
	}

	// output
	fmt.Println("\nDone.")
}
