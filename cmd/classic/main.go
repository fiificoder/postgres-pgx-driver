package main

import (
	"context"
	"database/sql"
	"log"
	"postgresql-intro/app"
	"postgresql-intro/website"
	"time"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgress:password@localhost:5432/website")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	websiteRepository := website.NewPostgreSQLClassicRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.RunRepositoryDemo(ctx, websiteRepository)
}
