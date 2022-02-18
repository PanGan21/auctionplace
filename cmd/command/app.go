package command

import (
	"context"
	"errors"

	"github.com/PanGan21/auctionplace/config"
	"github.com/spf13/cobra"
)

func NewRootCommand(ctx context.Context) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:               "auctionplace",
		Short:             "Run the auction place service",
		PersistentPreRunE: config.Setup,
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("please specify a command: [deploy, run, trade]")
		},
	}

	rootCommand.Flags().StringVarP(
		&config.Filename,
		"config",
		"f",
		"config/local.yaml",
		"Relative path to the config file",
	)

	rootCommand.PersistentFlags().StringP("blockchain.pk", "k", "", "Account private key")
	rootCommand.AddCommand(NewDeployCommand(ctx))
	rootCommand.AddCommand(NewRunnerCommand(ctx))

	return rootCommand
}
