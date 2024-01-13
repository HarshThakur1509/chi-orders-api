package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/HarshThakur1509/chi-orders-api/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
