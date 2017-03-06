package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fb := [...]int64{0, 0, 1, 0, 2, 1, 0, 0, 1, 2, 0, 1, 0, 0, 3}
	//var seed int64 = 2829501022
	var seed int64 = 0

	for {
		rand.Seed(seed)
		s := fmt.Sprintf("seed: %d:", seed)
		for i := int64(0); i < 15; i++ {
			n := rand.Int63()
			if n%int64(4) != fb[i] {
				if i > 13 {
					s = fmt.Sprintf("%s %d ", s, n%4)
					fmt.Println(s)
				}
				break
			} else {
				s = fmt.Sprintf("%s %d ", s, n%4)
			}
			if i == 14 {
				fmt.Println(s)
			}
		}
		seed++
	}
}
