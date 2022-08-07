package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

func sieveOfAtkin(N int) []string {
	var x, y, n int
	nsqrt := math.Sqrt(float64(N))

	is_prime := make([]bool, N+1)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < N; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	primes := make([]string, 0, 1270606)
	for x = 0; x < len(is_prime)-1; x++ {
		if is_prime[x] {
			primes = append(primes, strconv.Itoa(x))
		}
	}

	// primes is now a slice that contains all primes numbers up to N
	// so let's print them
	return primes
}

func F(n int) int {
	cPrimes := make(chan []string)

	go func() {
		primes := sieveOfAtkin(n)
		cPrimes <- primes
	}()

	primes := <-cPrimes

	lastPrime := primes[len(primes)-1]
	maxNoOfEven := len(lastPrime)
	foundPrime := "-1"

	for i := len(primes) - 1; i >= 0; i-- {
		if len(primes[i]) == maxNoOfEven {
			maxNoOfEven--
		}

		r := regexp.MustCompile(`[02468]`)
		matches := r.FindAllString(primes[i], -1)
		currentPrimeMatches := r.FindAllString(foundPrime, -1)

		if len(currentPrimeMatches) >= maxNoOfEven {
			break
		}

		if len(matches) > len(currentPrimeMatches) {
			foundPrime = primes[i]
		}
	}

	lastPrimeNum, _ := strconv.Atoi(foundPrime)
	return lastPrimeNum
}

func main() {
	timeNow := time.Now()
	arr := []int{1989370, 4626706, 1736937, 3587549, 3381743, 4556737}

	for _, num := range arr {
		fmt.Println(F(num))
	}
	timeSince := time.Since(timeNow)

	fmt.Printf("Execution time: %s", timeSince)
}
