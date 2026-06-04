package main

import "fmt"

func main() {
	accountAgeFloat := 2.6
	accountAgeInt := int64(accountAgeFloat)
	fmt.Println("your account has existed for", accountAgeInt, "years")
}

//output: your account has existed for 2 years
