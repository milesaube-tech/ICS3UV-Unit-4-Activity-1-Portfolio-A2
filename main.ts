/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-05
 * @fileoverview This program prints all even numbers between a starting and ending odd number.
 */

// get user input
let startValue: number = Number(prompt("Please enter a starting value (it must be an odd number): "));
let endValue: number = Number(prompt("Please enter an ending value (it must be an odd number): "));

// set variables
let counter: number = 0;
let output: string = "";

// set up loop
for ( counter = startValue + 1; counter <= endValue; counter = counter + 2 )
{
  output = output + counter + " "
}

console.log(output)