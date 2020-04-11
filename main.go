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
	trials := 100000
	successSwap := 0
	swap := true
	for i := 0; i < trials; i++ {
		if win(swap) {
			successSwap++
		}
	}
	fmt.Printf("Success with swap: %v is %d%%",swap, successSwap*100.0/trials)
}
