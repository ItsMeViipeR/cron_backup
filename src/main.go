package main

import (
	"log"
	"os"
	"time"
)

func main() {
	config, err := ReadConfigFile("settings.json")

	if err != nil {
		log.Fatalf("Erreur lors du chargement de la configuration : %v", err)
	}

	delay, convert_err := convert(config.Delay)

	if convert_err != nil {
		log.Fatalf("Erreur lors de la conversion du délai : %v", convert_err)
	}
	log.Printf("Délai converti : %d secondes", delay)

	action := func() {
		logger := log.New(os.Stdout, "", 0)

		now := time.Now()
		dateFr := now.Format("02/01/2006 15:04:05")

		err := HandleBackup(config.BackupDir, config.OutDir)
		if err != nil {
			logger.Printf("Erreur lors du backup : %v", err)
		}
		logger.Printf("[%s] Exécuté\n", dateFr)
	}

	action()

	go scheduleAction(time.Duration(delay)*time.Second, action)

	// Prevent the main function from exiting immediately
	select {}
}
