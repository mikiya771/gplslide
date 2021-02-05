package main

import "fmt"

func main() {
	max := 10000
	fmt.Printf("length of prime numbers under %d is %d\n", max, binarySerach(max, primes))

}
func binarySerach(goal int, primes []int) int {
	l := len(primes)
	if l == 0 {
		return 0
	}
	hp := l / 2
	if primes[hp] <= goal {
		return hp + 1 + binarySerach(goal, primes[hp+1:])
	} else {
		return binarySerach(goal, primes[:hp])
	}
}
