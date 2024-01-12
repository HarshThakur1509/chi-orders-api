package main

import (
	"context"
	"fmt"

	"github.com/HarshThakur1509/chi-orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start the app:", err)
	}
}
