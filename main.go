package main

import (
	"log"

	"github.com/ose-micro/omc-migrate/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("ose-migrate failed: %v", err)
	}
}
