package main

import (
	"avito-test/internal/app"
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	if err := app.RunApp(ctx); err != nil {
		log.Fatalf("error running avito_test service: %v", err)
	}
}
