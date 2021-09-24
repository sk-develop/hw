package main

import "fmt"

func main() {
	Hello("sk")
}

func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}