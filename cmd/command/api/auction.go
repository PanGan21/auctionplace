package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/PanGan21/auctionplace/config"
	"github.com/PanGan21/auctionplace/pkg/blockchain"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

const GetAllAction = "all"
const GetUserAction = "user"
const CreateAuctionAction = "create"

var allowedActions = map[string]struct{}{
	GetAllAction:        {},
	GetUserAction:       {},
	CreateAuctionAction: {},
}

var (
	ErrInvalidAuctionAction = errors.New("invalid auction action")
	ErrInvalidUserAddress   = errors.New("invalid user address")
)

// NewAuctionCommand created a new auction command
func NewAuctionCommand(ctx context.Context) *cobra.Command {
	var (
		action      string
		account     string
		name        string
		description string
		min         int64
	)
	auctionCommand := &cobra.Command{
		Use:   "auction",
		Short: "Get a list of all the auctions or a user's auctions",
		Run: func(cmd *cobra.Command, args []string) {
			err := runAuction(ctx, action, account, name, description, min)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	auctionCommand.Flags().StringVar(&action, "action", "", "Action to perform: all")
	auctionCommand.Flags().StringVar(&account, "account", "", "Account that owns the auctions to be fetched")
	auctionCommand.Flags().StringVar(&name, "name", "", "The name of the auction")
	auctionCommand.Flags().StringVar(&description, "description", "", "The description of the auction")
	auctionCommand.Flags().Int64Var(&min, "min", 0, "The minimum amount of the auction")
	auctionCommand.MarkFlagRequired("action")
	return auctionCommand
}

func runAuction(ctx context.Context, action string, account string, name string, description string, min int64) error {
	if _, ok := allowedActions[action]; !ok {
		return ErrInvalidAuctionAction
	}

	ctxCall, cancel := context.WithTimeout(ctx, config.App.Blockchain.TimeoutIn)
	defer cancel()

	client, err := ethclient.DialContext(ctxCall, config.App.Blockchain.Address)
	if err != nil {
		return err
	}

	runner := blockchain.NewAuctionRunner(config.App.Blockchain.PrivateKey, config.App.Contract.Address)

	switch action {
	case GetAllAction:
		auctions, err := runner.GetAuctions(ctx, client)
		if err != nil {
			return err
		}

		fmt.Println("Auctions: ", auctions)
	case GetUserAction:
		auctions, err := runner.GetUserAuctions(ctx, client, account)
		if err != nil {
			return err
		}

		fmt.Println("Auctions: ", auctions)
	case CreateAuctionAction:
		createAuctionOpts := blockchain.NewCreateAuctionOpts(name, description, min)
		err = runner.CreateAuction(ctx, client, createAuctionOpts)
		if err != nil {
			return err
		}

		fmt.Println("Successfully created auction with name: ", name, " description: ", description, " min: ", min)
	}

	return nil
}
