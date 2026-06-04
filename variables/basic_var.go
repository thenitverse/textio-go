package main

import "fmt"

func main() {
	var username string
	username = "Shubh"

	var is_admin bool
	is_admin = true

	var permission int
	permission = 100

	var costpersms float64
	costpersms = 0.05

	fmt.Println("username:", username)
	fmt.Println("isAdmin:", is_admin)
	fmt.Println("permissions:", permission)
	fmt.Println("costPerSMS:", costpersms)

}

/*output:
username: Shubh
isAdmin: true
permissions: 100
costPerSMS: 0.05*/
