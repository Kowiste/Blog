package main

import "flag"

func main() {
	namePrint := flag.String("n", "noName", "Name to print")
	times := flag.Int("number", 5, "Number of time that repeat the name")

	flag.Parse()
	for i := 0; i < *times; i++ {
		println("My name is " + *namePrint)
	}
	println("Finish")
}
