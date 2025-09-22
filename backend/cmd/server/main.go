package main

import (
	"github.com/hueter57/catalyst/backend/internal/domain"
)

func main() {
	// Setup ent client
	client, err := domain.Connect()
	if err != nil {
		panic(err)
	}
	defer client.Close()
}
