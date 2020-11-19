package main

import (
	console ".."
)

func main() {
	a := "a"
	console.PrintColoredF("%s, %d, %s : %s : %p", console.Black, "hello", 3, "twelve", "fourteen", &a)
}
