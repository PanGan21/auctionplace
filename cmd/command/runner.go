package command

import (
	"context"
	"errors"

	"github.com/PanGan21/auctionplace/cmd/command/api"
	"github.com/spf13/cobra"
)

func NewRunnerCommand(ctx context.Context) *cobra.Command {
	runCommand := &cobra.Command{
		Use:   "run",
		Short: "Run smart contract methods",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("please specify a subcommand: [auction]")
		},
	}

	runCommand.AddCommand(api.NewAuctionCommand(ctx))
	runCommand.AddCommand(api.NewOfferCommand(ctx))
	runCommand.AddCommand(api.NewTradeCommand(ctx))

	return runCommand
}
