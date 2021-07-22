/*
Author : Brad Myrick
Title : DOUBLY-LINKED-LIST
Description :
Ask the user in the console if they would like to input a string and how many.
if so ask for the string and add it to the list.
either way, ask the user if they would like to print the list.
save list to a file. if the file already exists, add the strings to the list.
*/
package main

import (
	"fmt"  // for everything
	"os"   // for clear
	"sync" // for waitgroup
	"time" // for sleep just for the accepted message thats just for funzies
)

// simple block consisting of a string and a pointer to the next and previous
type Node struct {
	data string
	next *Node
	prev *Node
}

// struct that contains nodes initialized with pointers to the head and tail of the list
type doublyLinkedList struct {
	head *Node
	tail *Node
}

// add a node to the front of a doubly linked list. right now I don't have an add to back.
func (l *doublyLinkedList) add(data string) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

// print in the order it was saved
func (l *doublyLinkedList) print() {
	for node := l.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

// print in reverse order
func (l *doublyLinkedList) printBackwards() {
	for node := l.tail; node != nil; node = node.prev {
		fmt.Println(node.data)
	}
}

// this is bad TODO fix it
func (l *doublyLinkedList) printMiddle() {
	for node := l.head; node != nil; node = node.next {
		if node.next != nil {
			fmt.Println(node.data)
		}
	}
}

// this is bad TODO fix it
func (l *doublyLinkedList) printMiddleBackwards() {
	for node := l.tail; node != nil; node = node.prev {
		if node.prev != nil {
			fmt.Println(node.data)
		}
	}
}

//TODO os specific checks
//this only works for windows right now
func clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

// returns my run bool telling the program if the user wants to add to the list or move to the print menu
func askTheQuestion(num *int) bool {
	var listinput string = ""
	fmt.Println("Would you like to add a string to the list? (y/n)")
	fmt.Scanln(&listinput)
	clear()
	if listinput == "y" || listinput == "Y" {
		fmt.Println("How many strings would you like to add?")
		fmt.Scanln(num)
		clear()
		return true
	} else {
		return false
	}
}

// TODO os specific checks for making files.
// TODO if the file is already there add the pre existing strings to the list adjust the starting value of i in the loop
func buildList(numToAdd int, l *doublyLinkedList) {
	var listdata string

	f, err := os.Create("list.txt")
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < numToAdd; i++ {
		fmt.Println("What would you like to add to the list?")
		fmt.Scanln(&listdata)
		l.add(listdata)
		f.WriteString(listdata)
		f.WriteString("\n")
		clear()
		fmt.Println("Accepted")
		time.Sleep(250 * time.Millisecond)
		clear()
	}

	defer f.Close()
}

// prints from the list in memory.
func printOptions(l *doublyLinkedList, wg *sync.WaitGroup) {
	input := ""
	fmt.Println("Would you like to print the list? (y/n)")
	fmt.Scanln(&input)
	clear()
	if input == "y" || input == "Y" {
		l.print()
	}
	fmt.Println("Would you like to print the list backwards? (y/n)")
	fmt.Scanln(&input)
	clear()
	if input == "y" || input == "Y" {
		l.printBackwards()
	}
	fmt.Println("Would you like to print the list in the middle? (y/n)")
	fmt.Scanln(&input)
	clear()
	if input == "y" || input == "Y" {
		l.printMiddle()
	}
	fmt.Println("Would you like to print the list backwards in the middle? (y/n)")
	fmt.Scanln(&input)
	clear()
	if input == "y" || input == "Y" {
		l.printMiddleBackwards()
	}

	wg.Done()
}

// needed to pack this together for a Wait Group.  I was getting the print menu in the middle of setting the strings.
func start(run bool, numToAdd int, l *doublyLinkedList, wg *sync.WaitGroup) {
	if run {
		buildList(numToAdd, l)

	} else {
		fmt.Println("run is false")
	}
	wg.Done()
}

// *** START *** MAIN ***
func main() {
	l := &doublyLinkedList{}
	var numToAdd int
	var wg sync.WaitGroup
	var run bool
	run = askTheQuestion(&numToAdd)
	wg.Add(1)
	start(run, numToAdd, l, &wg)
	wg.Wait()
	var input string
	fmt.Println("Would you like to see the print options (y/n)")
	fmt.Scanln(&input)
	clear()
	if input == "y" || input == "Y" {
		wg.Add(1)
		printOptions(l, &wg)
	}
	wg.Wait()
	fmt.Println("Goodbye!")
}
