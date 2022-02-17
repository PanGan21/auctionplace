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

var allowedActions = map[string]struct{}{
	GetAllAction:  {},
	GetUserAction: {},
}

var (
	ErrInvalidAuctionAction = errors.New("invalid auction action")
	ErrInvalidUserAddress   = errors.New("invalid user address")
)

// NewAuctionCommand created a new auction command
func NewAuctionCommand(ctx context.Context) *cobra.Command {
	var (
		action  string
		account string
	)
	auctionCommand := &cobra.Command{
		Use:   "auction",
		Short: "Get a list of all the auctions or a user's auctions",
		Run: func(cmd *cobra.Command, args []string) {
			err := runAuction(ctx, action, account)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	auctionCommand.Flags().StringVar(&action, "action", "", "Action to perform: all")
	auctionCommand.Flags().StringVar(&account, "account", "", "Account that owns the auctions to be fetched")
	auctionCommand.MarkFlagRequired("action")
	return auctionCommand
}

func runAuction(ctx context.Context, action string, account string) error {
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
		if len(account) == 0 {
			return ErrInvalidUserAddress
		}
		auctions, err := runner.GetUserAuctions(ctx, client, account)
		if err != nil {
			return err
		}

		fmt.Println("Auctions: ", auctions)
	}

	return nil
}
