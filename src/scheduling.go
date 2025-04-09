package main

import (
	"time"
)

func scheduleAction(interval time.Duration, action func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		action()
	}
}
