package main

import (
	"fmt"
	"reflect"
)

// only run -> go run golang-datatypes.go
// bool: represents a boolean value and is either true or false
// Numeric: represents integer types, floating point values, and complex types
// string: represents a string value

// https://www.w3schools.com/go/go_data_types.php

func main() {

	// Text Type: str
	// strings
	var s string = "theory of Strings"
	fmt.Println(reflect.TypeOf(s), ":", s)

	// Numeric Types: int, float, complex

	var i int = 101
	fmt.Println(reflect.TypeOf(i), ":", i)

	var f32 float32 = 3.141592653589793238
	fmt.Println(reflect.TypeOf(f32), ":", f32)

	var f64 float64 = 3.141592653589793238
	fmt.Println(reflect.TypeOf(f64), ":", f64)

	var c complex64 = 3 + 5i
	fmt.Println(reflect.TypeOf(c), ":", c)

	// Sequence Types: array, slices, range

	arr := [4]string{"Get", "in", "the", "array"}
	fmt.Println(reflect.TypeOf(arr), ":", arr)

	mySlice := []string{"Give", "me", "that", "delicious", "slice"}

	fmt.Println(reflect.TypeOf(mySlice), ":", mySlice)

	// Mapping Type: map

	// Set Types: set, frozenset

	// Boolean Type: bool
	var b bool = true
	fmt.Println(reflect.TypeOf(b), ":", b)

	// Binary Types: bytes, bytearray, memoryview

	// None Type: NoneType

}
