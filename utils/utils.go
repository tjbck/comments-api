package utils

import (
	"log"
)

func LogInfo(port string) {
	log.Printf("Starting up on http://localhost:%s", port)
}
