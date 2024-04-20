package cmd

import (
	"log"

	"github.com/dionysia-dev/dionysia/internal/config"
	"github.com/dionysia-dev/dionysia/internal/logging"
	"github.com/dionysia-dev/dionysia/internal/queue"
	"github.com/dionysia-dev/dionysia/internal/task"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func NewWorker() *cobra.Command {
	return &cobra.Command{
		Use:   "worker",
		Short: "Run worker server",
		Run: func(*cobra.Command, []string) {
			if err := godotenv.Load(".env"); err != nil {
				log.Println("Could not load env file")
			}

			app := fx.New(
				fx.Provide(config.New),
				fx.Provide(logging.New),
				fx.Provide(queue.NewServer),
				fx.Invoke(task.Run),
			)

			app.Run()
		},
	}
}
