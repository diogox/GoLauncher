package app

import (
	"github.com/spf13/cobra"
)

func NewCliApp() *CliApp {

	cmd := &cobra.Command{
		Use: "hello",
		Run: func(cmd *cobra.Command, args []string) {
			startLauncher()
		},
	}

	return &CliApp{
		cmd: cmd,
	}
}

type CliApp struct {
	cmd *cobra.Command
}

func (ca *CliApp) Start() error {

	return ca.cmd.Execute()
}