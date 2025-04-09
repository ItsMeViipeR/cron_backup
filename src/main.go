package main

import (
	"log"
	"time"
)

func main() {
	action := func() {
		err := HandleBackup(".", "/Users/viiper/Desktop/out")
		if err != nil {
			log.Printf("Erreur lors du backup : %v", err)
		}
		log.Println("Exécuté")
	}

	action()

	go scheduleAction(time.Minute, action)

	// Prevent the main function from exiting immediately
	select {}
}
