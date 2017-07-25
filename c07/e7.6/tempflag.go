package main

import (
	"flag"
	"fmt"

	"github.com/sszarek/go-programming-lang-excercises/c07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
