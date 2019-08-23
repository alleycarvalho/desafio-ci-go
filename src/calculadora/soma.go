package main

import (
  "fmt"
)

func soma(x, y int) int {
    return x + y
}

func main () {
    var x, y = 5, 5

	total := soma(x, y)

    fmt.Printf("A soma de %v + %v é igual a %v", x, y, total);
}
