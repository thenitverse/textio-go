package main

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan

	}

}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}
func main() {
	testCases := []int{0, 1, 3, 4, 13}

	for _, numDBs := range testCases {
		fmt.Println("---------------------------------")
		fmt.Printf("Testing %v Databases...\n\n", numDBs)

		dbChan, count := getDBsChannel(numDBs)
		waitForDBs(numDBs, dbChan)

		for *count != numDBs {
			fmt.Println("...")
		}

		fmt.Printf("\nexpected count: %v\n", numDBs)
		fmt.Printf("actual count:   %v\n", *count)
		fmt.Printf("channel length: %v\n", len(dbChan))
	}
}

/*output:
---------------------------------
Testing 0 Databases...


expected count: 0
actual count:   0
channel length: 0
---------------------------------
Testing 1 Databases...

Database 1 is online

expected count: 1
actual count:   1
channel length: 0
---------------------------------
Testing 3 Databases...

Database 1 is online
Database 2 is online
Database 3 is online

expected count: 3
actual count:   3
channel length: 0
---------------------------------
Testing 4 Databases...

Database 1 is online
Database 2 is online
Database 3 is online
...
...
...
...
...
...
...
...
...
...
...
...
...
...
...
...
Database 4 is online
...

expected count: 4
actual count:   4
channel length: 0
---------------------------------
Testing 13 Databases...

Database 1 is online
Database 2 is online
Database 3 is online
Database 4 is online
Database 5 is online
Database 6 is online
Database 7 is online
Database 8 is online
Database 9 is online
Database 10 is online
Database 11 is online
Database 12 is online
Database 13 is online

expected count: 13
actual count:   13
channel length: 0
*/
