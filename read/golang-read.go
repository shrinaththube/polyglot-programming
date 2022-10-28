package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// only run - go run golang-read.go
// https://gobyexample.com/reading-files
// https://zetcode.com/golang/readfile/

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	filename := "sample-read.txt"

	fmt.Println("****************************************")
	fmt.Println("1) Read line by line")
	fmt.Println("****************************************")

	fileOperator, err := os.Open(filename)
	errorCheck(err)

	defer fileOperator.Close()

	scanner := bufio.NewScanner(fileOperator)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	errorCheck(scanner.Err())

	fmt.Println("****************************************")
	fmt.Println("2) Read everything using os module")
	fmt.Println("****************************************")

	content, err := os.ReadFile(filename)
	errorCheck(err)
	fmt.Println(string(content))

	fmt.Println("****************************************")
	fmt.Println("3) Read everything using ioutil")
	fmt.Println("****************************************")

	content, err = ioutil.ReadFile(filename)
	errorCheck(err)
	fmt.Println(string(content))

}
