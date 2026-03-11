package main

import (
	"context"
	"os"

	"github.com/cliquefarma/aws-secrets-search/internal/adapter/driver/cli"
)

func main() {
	app, err := cli.NewApp()
	if err != nil {
		os.Exit(1)
	}

	if err := app.Run(context.Background()); err != nil {
		os.Exit(1)
	}
}
