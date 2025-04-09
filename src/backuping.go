package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"time"
)

func HandleBackup(dir string, outdir string) error {
	// Résolution du chemin absolu de dir pour garantir la gestion correcte de "."
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	// Nom du fichier zip basé sur le nom du dossier
	zipName := filepath.Base(absDir) + ".zip"
	date := time.Now().Format("2006-01-02 15:04:05") // référence pour le format de date de Go
	zipPath := filepath.Join(outdir, date+"-"+zipName)

	// Création du fichier zip
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Création du writer zip
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Balayage récursif du dossier
	err = filepath.Walk(absDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Si c'est le fichier zip en cours de création, on l'ignore
		if path == zipPath {
			return nil
		}

		// Calcul du chemin relatif dans le zip
		relPath, err := filepath.Rel(absDir, path)
		if err != nil {
			return err
		}

		// Création de l'en-tête
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = relPath

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		// Création de l'entrée dans le zip
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Si c'est un fichier, on copie le contenu
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(writer, file)
			file.Close() // fermeture immédiate
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
