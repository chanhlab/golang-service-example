package main

import (
	"log"
	"os"

	"github.com/chanhlab/golang-service-example/cmd/api"
	"github.com/chanhlab/golang-service-example/cmd/migration"
	"github.com/chanhlab/golang-service-example/cmd/worker"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "golang example service",
		Usage: "golang example service",
		Commands: []*cli.Command{
			{
				Name:  "api",
				Usage: "start api",
				Action: func(ctx *cli.Context) error {
					return api.RunAPI(ctx.Context)
				},
			},
			{
				Name: "migration",
				Action: func(ctx *cli.Context) error {
					return migration.RunMigration(ctx.Context)
				},
			},
			{
				Name: "worker",
				Action: func(ctx *cli.Context) error {
					return worker.RunWorker(ctx.Context)
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
