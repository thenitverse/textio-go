package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case _, ok := <-snapshotTicker:
			if ok {
				takeSnapshot(logChan)
			}
		case _, ok := <-saveAfter:
			if ok {
				saveSnapshot(logChan)
				return

			}
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)

		}
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

func main() {
	logChan := make(chan string)
	go saveBackups(
		time.Tick(1*time.Second),
		time.After(2500*time.Millisecond),
		logChan,
	)

	for msg := range logChan {
		println(msg)
	}
}

/*output:
Nothing to do, waiting...
Nothing to do, waiting...
Taking a backup snapshot...
Nothing to do, waiting...
Nothing to do, waiting...
Taking a backup snapshot...
Nothing to do, waiting...
All backups saved!*/
