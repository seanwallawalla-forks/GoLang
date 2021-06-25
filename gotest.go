package main

import "fmt"

var yes string = "yes"
var no string = "no"
var myList []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func print(str string) {
	fmt.Println(str)
}

func NumInList(list []int, num int) bool {
	for _, i := range list {
		if i == num {
			return true
		}
	}

	return false
}

func main() {
	if NumInList(myList, 0) {
		print(yes)
	} else {
		print(no)
	}
}
