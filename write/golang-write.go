package main

import (
	"log"
	"os"
)

// only run- go run golang-write.go
// https://gobyexample.com/writing-files
// https://zetcode.com/golang/writefile/
// https://screenrant.com/007-iconic-james-bond-quotes/

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// "Bond. James Bond" -Introduction

// "I just showed someone your watch. Really blew their mind." -No Time To Die
// "A martini. Shaken, not stirred." -Drink Order
// "There's a saying in England: Where there's smoke, there's fire." -Russia With Love
// "Why is it that people who can't take advice always insist on giving it?" -Casino Royale
// "I thought Christmas only comes once a year." -The World Is Not Enough
// "Shocking. Positively shocking." -Goldfinger
// "Now there's a name to die for." -Die Another Day

func main() {
	filename := "sample-write.txt"

	quote1 := []byte("\"Bond. James Bond\" -Introduction\n")
	err := os.WriteFile(filename, quote1, 0644)
	// os.WriteFile()
	errorCheck(err)

}
