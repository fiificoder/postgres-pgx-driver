package app

import (
	"context"
	"fmt"
	"log"
	"postgresql-intro/website"
)

func RunRepositoryDemo(ctx context.Context, websiteRepository website.Repository) {
	fmt.Println("1. MIGRATE REPOSITORY")

	if err := websiteRepository.Migrate(ctx); err != nil {
		log.Fatal(err)
	}
}
