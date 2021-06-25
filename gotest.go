package main

//learning GO
import (
	"fmt"
)

func print(str string) {
	fmt.Println(str)
}

func MakeSentence(list [7]string) [7]string {
	var ret [7]string
	for i := range list {
		if i == 0 {
			ret[i] = list[4]
		} else if i == 1 {
			ret[i] = list[6]
		} else if i == 2 {
			ret[i] = list[5]
		} else if i == 3 {
			ret[i] = list[2]
		} else if i == 4 {
			ret[i] = list[0]
		} else if i == 5 {
			ret[i] = list[1]
		} else {
			ret[i] = list[3]
		}
	}
	return ret
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
	yes := "yes"
	no := "no"

	myList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	strList := [7]string{"of", "bold", "array", "stuff", "Coffee", "an", "is"}

	coin := make(map[int]string)
	coin[1] = "Bitcoin"
	coin[2] = "Ethereum"
	coin[3] = "Doge"

	if NumInList(myList, 5) {
		no = yes
	} else {
		yes = no
	}

	sentence := MakeSentence(strList)
	for _, word := range sentence {
		print(word)
	}
}
