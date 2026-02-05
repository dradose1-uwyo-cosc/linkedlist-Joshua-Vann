//Joshua Vann
//COSC 3750
//2-1-26
//
/*
	Don't forget to run your go mod init command in your terminal
	Review the assignment instructions for running your code
	All the code you need to write should be put in the /ds/ package files
	Uncomment the import statement for your module name
	you can uncomment the tests in main as you go to test
	The code in main is not an extensive test, you should add more and test your code as needed
*/
package main

import (
	"fmt"
	"hw01/ds" //this needs to be changes to your module name and uncommented
)

func main() {
	fmt.Println("Only here so the import doesn't leave an error")

	linkedlist := &ds.LinkedList{}
	linkedlist.InsertAt(0, "first")
	fmt.Println("-------------")
	linkedlist.PrintList()
	fmt.Println("-------------")
	linkedlist.Insert("first")
	linkedlist.Insert("first")
	linkedlist.Insert("second")
	linkedlist.Insert("third")
	linkedlist.Insert("fourth")
	linkedlist.Insert("fifth") //5th is getting accidentally removed when removing third
	linkedlist.RemoveAt(4)
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")
	linkedlist.RemoveAll("first")
	linkedlist.PrintList()
	fmt.Println("-------------")
	linkedlist.Reverse()
	linkedlist.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist.GetSize())
	fmt.Println("-------------")

	linkedlist2 := &ds.LinkedList{}
	linkedlist2.Insert("A")
	linkedlist2.Insert("A")
	linkedlist2.Insert("A")
	linkedlist2.Insert("A")
	linkedlist2.Insert("B")
	linkedlist2.Insert("C")
	linkedlist2.Insert("D")
	linkedlist2.Insert("E")
	linkedlist2.Insert("F")
	linkedlist2.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist2.GetSize())
	linkedlist2.InsertAt(3, "7")
	linkedlist2.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist2.GetSize())
	linkedlist2.RemoveAll("A")
	linkedlist2.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist2.GetSize())
	linkedlist2.Reverse()
	linkedlist2.PrintList()
	fmt.Println("The size of the linked list is:", linkedlist2.GetSize())

	stack := &ds.Stack{}
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")
	data, _ := stack.Pop()
	println("Popped from stack:", data)

	queue := &ds.Queue{}
	queue.Push("first")
	queue.Push("second")
	queue.Push("third")
	data, _ = queue.Pop()
	println("Popped from queue:", data)
}
