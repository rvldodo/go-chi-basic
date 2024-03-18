package main

import (
	"context"
	"fmt"

	"github.com/rvldodo/chi-basic/appllication"
)

func main() {
	app := appllication.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed start the server:", err)
	}
}
