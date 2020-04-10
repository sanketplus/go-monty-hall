package main

import (
	"fmt"
	"math/rand"
	"time"
)

func win(swap bool) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prize := r.Int() % 3
	choice := r.Int() % 3
	if prize == choice {
		if swap {
			return false
		} else {
			return true
		}
	}

	if swap {
		return true
	}
	return false
}

func main() {
	fmt.Println("Hello, playground")
	trials := 30000
	successSwap := 0

	for i := 0; i < trials; i++ {
		if win(true) {
			successSwap++
		}
	}
	fmt.Println(successSwap*100.0/trials)
}
