package main

import (
	"fmt"
	"time"
)

func pingPong(numPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})
	done := make(chan struct{})
	go ponger(pings, pongs)
	go pinger(pings, numPings)
	go func() {
		i := 0
		for range pongs {
			fmt.Println("got pong", i)
			i++
		}
		fmt.Println("pongs done")
		close(done)
	}()
	<-done

}

func pinger(pings chan struct{}, numPings int) {
	sleepTime := 50 * time.Millisecond
	for i := 0; i < numPings; i++ {
		fmt.Printf("sending ping %v\n", i)
		pings <- struct{}{}
		time.Sleep(sleepTime)
		sleepTime *= 2
	}
	close(pings)
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numPings int) {
	fmt.Println("Starting game...")
	pingPong(numPings)
	fmt.Println("===== Game over =====")
}

func main() {
	test(4)
	test(3)
	test(2)
}

/*output:
Starting game...
sending ping 0
got ping 0, sending pong 0
got pong 0
sending ping 1
got ping 1, sending pong 1
got pong 1
sending ping 2
got ping 2, sending pong 2
got pong 2
sending ping 3
got ping 3, sending pong 3
got pong 3
pings done
pongs done
===== Game over =====
Starting game...
sending ping 0
got ping 0, sending pong 0
got pong 0
sending ping 1
got ping 1, sending pong 1
got pong 1
sending ping 2
got ping 2, sending pong 2
got pong 2
pings done
pongs done
===== Game over =====
Starting game...
sending ping 0
got ping 0, sending pong 0
got pong 0
sending ping 1
got ping 1, sending pong 1
got pong 1
pings done
pongs done
===== Game over =====*/
