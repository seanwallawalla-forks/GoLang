/* FizzBuzz Brad Myrick
if number is divisible by 3 print “Fizz”
if number is divisible by 5 print “Buzz”
if number is divisible by both 3 & 5 print “Fizz Buzz”
*/
package main

import (
	"fmt"
)

func fizzbuzz(x int) {
	for i := 1; i < x; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				println("fizzbuzz")
			}
			println("fizz")
		} else if i%5 == 0 {
			println("buzz")
		} else {
			println(i)
		}
	}

}

func main() {
	number := 0
	fmt.Println("till what end?")
	fmt.Scanln(&number)
	fizzbuzz(number)
}
