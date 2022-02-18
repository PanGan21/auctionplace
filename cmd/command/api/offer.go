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

const GetUserOfferAction = "user"
const CreateOfferAction = "create"

var allowedOfferActions = map[string]struct{}{
	GetUserOfferAction: {},
	CreateOfferAction:  {},
}

var (
	ErrInvalidOfferAction = errors.New("invalid offer action")
)

// NewOfferCommand created a new offer command
func NewOfferCommand(ctx context.Context) *cobra.Command {
	var (
		action    string
		account   string
		auctionId int64
		amount    int64
	)
	offerCommand := &cobra.Command{
		Use:   "offer",
		Short: "Get a list of a user's offers or create an offer",
		Run: func(cmd *cobra.Command, args []string) {
			err := runOffer(ctx, action, account, auctionId, amount)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	offerCommand.Flags().StringVar(&action, "action", "", "Action to perform: all")
	offerCommand.Flags().StringVar(&account, "account", "", "Account that owns the offers to be fetched")
	offerCommand.Flags().Int64Var(&auctionId, "auctionId", 0, "The auction id for which you want to create an offer")
	offerCommand.Flags().Int64Var(&amount, "amount", 0, "The amount for the offer to be created")
	offerCommand.MarkFlagRequired("action")
	return offerCommand
}

func runOffer(ctx context.Context, action string, account string, auctionId int64, amount int64) error {
	if _, ok := allowedOfferActions[action]; !ok {
		return ErrInvalidOfferAction
	}

	ctxCall, cancel := context.WithTimeout(ctx, config.App.Blockchain.TimeoutIn)
	defer cancel()

	client, err := ethclient.DialContext(ctxCall, config.App.Blockchain.Address)
	if err != nil {
		return err
	}

	runner := blockchain.NewOfferRunner(config.App.Blockchain.PrivateKey, config.App.Contract.Address)

	switch action {
	case GetUserAction:
		offers, err := runner.GetUserOffers(ctx, client, account)
		if err != nil {
			return err
		}

		fmt.Println("Offers: ", offers)
	case CreateOfferAction:
		createOfferOpts := blockchain.NewCreateOfferOpts(auctionId, amount)
		err = runner.CreateOffer(ctx, client, createOfferOpts)
		if err != nil {
			return err
		}

		fmt.Println("Successfully created offer for auction: ", auctionId)
	}

	return nil
}
