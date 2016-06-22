package utils

import (
	"os"
	"log"
)

func CheckFileExists(fileName string) {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
}
