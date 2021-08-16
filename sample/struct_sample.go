package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	p := person{
		"Jay", 20, []string{"kimchi", "ramen"},
	}

	p2 := person{
		name:    "Jay",
		age:     20,
		favFood: []string{"Kimchi"},
	}

	fmt.Println(p)
	fmt.Println(p2)
}
