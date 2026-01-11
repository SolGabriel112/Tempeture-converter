package main

import (
	"fmt"
)

func main() {
	var choice string
	fmt.Print("C or F\n")
	fmt.Scan(&choice)
	switch {
	case choice == "C":
		C_to_F()
	case choice == "F":
		F_to_C()
	}
}

func C_to_F() {
	var degree float32
	var degre_converted float32
	fmt.Print("how many whole degrees celsius\n")
	fmt.Scan(&degree)
	degre_converted = (degree * 1.8) + 32
	fmt.Print(degree, " celsius is ", degre_converted, " Fahrenheit\n")
	main()
}

func F_to_C() {
	var degree float32
	var degre_converted float32
	fmt.Print("how many whole degrees Fahrenheit\n")
	fmt.Scan(&degree)
	degre_converted = (degree - 32) / 1.8
	fmt.Print(degree, " Fahrenheit is ", degre_converted, " celsius\n")
	main()
}
