package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hueter57/catalyst/backend/internal/domain"
)

type Mode int

const (
	ModeInvalid Mode = iota
	ModeDiff
	ModeApply
)

func loadMode() (Mode, error) {
	usage := fmt.Errorf("usage: migrations <diff|apply>")

	if len(os.Args) != 2 {
		return ModeInvalid, usage
	}
	switch os.Args[1] {
	case "diff":
		return ModeDiff, nil
	case "apply":
		return ModeApply, nil
	default:
		return ModeInvalid, usage
	}
}

func main() {
	client, err := domain.Connect()
	if err != nil {
		log.Fatal(err)
	}

	mode, err := loadMode()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	switch mode {
	case ModeInvalid:
	case ModeDiff:
		if err := domain.MigrateDiff(ctx, client); err != nil {
			log.Fatal(err)
		}
		log.Println("Computed migration diffs successfully")
	case ModeApply:
		if err := domain.MigrateApply(ctx, client); err != nil {
			log.Fatal(err)
		}
		log.Println("Applied migration successfully")
	}
}
