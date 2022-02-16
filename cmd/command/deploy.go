package command

import (
	"context"
	"fmt"

	"github.com/PanGan21/auctionplace/config"
	"github.com/PanGan21/auctionplace/pkg/blockchain"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

func NewDeployCommand(ctx context.Context) *cobra.Command {
	deployCommand := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy AuctionPlace contract to blockchain",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deploy(ctx)
		},
	}
	return deployCommand
}

func deploy(ctx context.Context) error {
	fmt.Println("Deploying contract")

	ctx, cancel := context.WithTimeout(ctx, config.App.Blockchain.TimeoutIn)
	defer cancel()

	client, err := ethclient.DialContext(ctx, config.App.Blockchain.Address)
	if err != nil {
		return err
	}

	deployer := blockchain.NewDeployer()
	err = deployer.Deploy(ctx, client)
	if err != nil {
		return err
	}

	fmt.Println("contract deployed at address: \n", deployer.ContractAddress())

	return nil
}
