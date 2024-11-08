/* Make the following concurrent */
package main

import "fmt"

func main() {
	primes := genPrimes(3, 100)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done!")
}

func genPrimes(start, end int) []int {
	primes := make([]int, 0)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
