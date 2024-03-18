package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/rvldodo/chi-basic/appllication"
)

func main() {
	app := appllication.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed start the server:", err)
	}
}
