package main

import "fmt"

// run command ref - https://yourbasic.org/golang/for-loop/
// https://go.dev/tour/flowcontrol/1
// https://www.educative.io/answers/what-is-the-for-range-loop-in-golang
// only run - go run golang-loops.go
// build and run - go build golang-loops.go
// ./golang-loops

func main() {

	coolList := []string{"i", "am", "reading", "all", "elements", "of", "this", "list"}

	fmt.Println("****************************************")
	fmt.Println("1) Let's start with 'while' loop")
	fmt.Println("****************************************")
	// for is the new while in Golang
	i := 0
	for i < len(coolList) {
		fmt.Println(coolList[i])
		i++
	}

	fmt.Println("****************************************")
	fmt.Println("2) Move on to 'for' loop")
	fmt.Println("****************************************")

	// for loop
	for i := 0; i < len(coolList); i++ {
		fmt.Println(coolList[i])
	}

	fmt.Println("****************************************")
	fmt.Println("3) Another variation of 'for' loop")
	fmt.Println("****************************************")

	// for loop
	for _, ele := range coolList {
		fmt.Println(ele)
	}

}
