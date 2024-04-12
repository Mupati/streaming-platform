package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/learn-video/dionysia/internal/api"
	"github.com/learn-video/dionysia/internal/config"
	"github.com/learn-video/dionysia/internal/db"
	"github.com/learn-video/dionysia/internal/logging"
	"github.com/learn-video/dionysia/internal/queue"
	"github.com/learn-video/dionysia/internal/service"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Run API server",
		Run: func(*cobra.Command, []string) {
			if err := godotenv.Load(".env"); err != nil {
				log.Println("Could not load env file")
			}

			app := fx.New(
				fx.Provide(config.New),
				fx.Provide(logging.New),
				fx.Provide(db.NewPool),
				fx.Provide(db.NewQuerier),
				fx.Provide(queue.NewClient),
				fx.Provide(service.NewNotificationHandler),
				api.Module,
			)

			app.Run()
		},
	}
}
