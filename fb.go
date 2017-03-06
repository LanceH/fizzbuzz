package main

import (
	"fmt"
	"math/rand"
)

var fb [4]string = [4]string{"", "fizz", "buzz", "fizzbuzz"}
var lucky int64 = 176064004

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 1 {
			rand.Seed(lucky)
		}
		fmt.Printf("%d: %s\n", i, fb[rand.Int63()%4])
	}
}
