package main

import (
	"strings"

	"github.com/mkaiho/go-ecs-batch-sample/logging"
	"github.com/spf13/cobra"
)

var (
	command *cobra.Command
)

func init() {
	logging.InitLoggerWithZap()
	command = newCommand()
}

func main() {
	if err := command.Execute(); err != nil {
		logging.GetLogger().Error(err, "error has occured")
	}
}

func newCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "echo args...",
		Short: "display args",
		Long:  "Display arguments on stdout.",
		RunE:  handle,
	}

	return &command
}

func handle(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		logging.GetLogger().Info("no args")
	} else {
		logging.GetLogger().Info(strings.Join(args, " "))
	}
	return nil
}
