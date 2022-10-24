package main

import (
	"fmt"
	"math/rand"
	"time"
)

// only run - go run golang-conditions.go
// build and run - go build golang-conditions.go
// ./golang-conditions

func main() {

	sun_is_bright, moon_is_cool := true, true

	// in km
	earth_diameter, moon_diameter := 12742, 3475
	// int moon_diameter = 3475

	if sun_is_bright {
		fmt.Println("Shine Shine and Shine")
	}

	if moon_is_cool {
		fmt.Println("Chill vibes")
	}

	if earth_diameter > moon_diameter {
		fmt.Println("Mother Earth is definitely bigger than moon")
		sun_is_bright = false
	} else {
		fmt.Println("Hmmm. Something looks odd")
	}

	if !sun_is_bright {
		fmt.Println("I am white dwarf now")
	}

	min, max := 1, 8

	rand.Seed(time.Now().UnixNano())

	//https://yourbasic.org/golang/switch-statement/
	//https://www.space.com/16080-solar-system-planets.html
	switch planetNo := rand.Intn(max-min+1) + min; planetNo {
	case 1:
		fmt.Println("I am Mercury. Closest to the sun")
	case 2:
		fmt.Println("I am Venus. Second planet from the sun")
	case 3:
		fmt.Println("Hey you know me. You live here. I am Earth")
	case 4:
		fmt.Println("Most popular one for traveling talks. Many are trying to land on me. I am Mars")
	case 5:
		fmt.Println("I am Jupiter. I am the largest plannet.")
	case 6:
		fmt.Println("I am Saturn. I have largest ring.")
	case 7:
		fmt.Println("I am Uranus. I have 27 moons and I have blue-green color.")
	case 8:
		fmt.Println("I am Neptune. I am coldest one. You guys found me by math haha!")
	default:
		fmt.Printf("Don't call me planet. I am not discovered yet.")
	}
}
