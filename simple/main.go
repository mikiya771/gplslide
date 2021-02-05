package main

import (
	"fmt"
)

func countPrimeNumber(max int) int {
	furui := map[int]bool{}
	primes := []int{}
	for i := 2; i <= max; i++ {
		if _, ok := furui[i]; !ok {
			primes = append(primes, i)
			furui[i] = true
			StiffOff(i, max, furui)
		}
	}
	return len(primes)
}

func StiffOff(start int, max int, furui map[int]bool) {
	for i := start + 1; i <= max; i++ {
		if _, ok := furui[i]; !ok && i%start == 0 {
			furui[i] = false
		}
	}
}

func main() {
	max := 10000
	fmt.Printf("length of prime numbers under %d is %d", max, countPrimeNumber(max))
}
