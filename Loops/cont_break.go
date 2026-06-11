package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue

		} else if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}

		}
		if isPrime {
			fmt.Println(n)

		}
	}
}

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}

/*output:
Primes up to 10:
2
3
5
7
===============================================================
Primes up to 20:
2
3
5
7
11
13
17
19
===============================================================
Primes up to 30:
2
3
5
7
11
13
17
19
23
29
===================================================
*/
