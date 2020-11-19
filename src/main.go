package main

import (
	console ".."
)

func main() {
	a := "a"
	console.PrintColoredF("%s, %d, %s : %s : %p\n", console.Black, "hello", 3, "twelve", "fourteen", &a)
	console.PrintColoredRainbow("한글도 되는지 확인하면서 also English")
	console.PrintColoredRainbowAni(1234)
}
