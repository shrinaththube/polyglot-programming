package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// only run- go run golang-write.go
// https://gobyexample.com/writing-files
// https://zetcode.com/golang/writefile/
// https://www.educative.io/answers/how-to-append-to-a-file-in-golang
// https://onelinerhub.com/golang/how-to-overwrite-content-in-file-if-it-exists
// https://screenrant.com/007-iconic-james-bond-quotes/

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := "sample-write.txt"

	fmt.Println("1) Write single line")
	quote1 := []byte("\"Bond. James Bond\" -Introduction\n\n")
	// overwrite the content.
	err := os.WriteFile(filename, quote1, 0666)
	errorCheck(err)

	// ------------------------------------------------------------------------------------------- //

	fmt.Println("2) Write from list of lines")

	fileOperator, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	errorCheck(err)

	defer fileOperator.Close()

	quotes := []string{"\"A martini. Shaken, not stirred.\" -Drink Order",
		"\"I just showed someone your watch. Really blew their mind.\" -No Time To Die",
		"\"There's a saying in England: Where there's smoke, there's fire.\" -Russia With Love",
	}

	for _, quote := range quotes {

		_, err := fileOperator.WriteString(quote + "\n\n")
		errorCheck(err)
	}

	// ------------------------------------------------------------------------------------------- //

	fmt.Println("3) Write multiline string in one shot")
	bondQuotes := `"Why is it that people who can't take advice always insist on giving it?" -Casino Royale

"I thought Christmas only comes once a year." -The World Is Not Enough

"Shocking. Positively shocking." -Goldfinger

"Now there's a name to die for." -Die Another Day`

	writer := bufio.NewWriter(fileOperator)
	_, err = writer.WriteString(bondQuotes)
	errorCheck(err)
	writer.Flush()
}
