package api

import (
	"context"
	"fmt"

	"github.com/PanGan21/auctionplace/config"
	"github.com/PanGan21/auctionplace/pkg/blockchain"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// const TradeAction = "trade"

// var allowedTradeActions = map[string]struct{}{
// 	TradeAction: {},
// }

func NewTradeCommand(ctx context.Context) *cobra.Command {
	var auctionId int64
	tradeCommand := &cobra.Command{
		Use:   "trade",
		Short: "Trade an auction for the best offer",
		Run: func(cmd *cobra.Command, args []string) {
			err := runTrade(ctx, auctionId)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

	tradeCommand.Flags().Int64Var(&auctionId, "auctionId", 0, "The auctionId to perform the trade")
	tradeCommand.MarkFlagRequired("auctionId")
	return tradeCommand
}

func runTrade(ctx context.Context, auctionId int64) error {
	// if _, ok := allowedOfferActions[action]; !ok {
	// 	return ErrInvalidOfferAction
	// }

	ctxCall, cancel := context.WithTimeout(ctx, config.App.Blockchain.TimeoutIn)
	defer cancel()

	client, err := ethclient.DialContext(ctxCall, config.App.Blockchain.Address)
	if err != nil {
		return err
	}

	runner := blockchain.NewTradeRunner(config.App.Blockchain.PrivateKey, config.App.Contract.Address)

	tradeOpts := blockchain.NewTradeOpts(auctionId)

	err = runner.TradeAuctionToOffer(ctx, client, tradeOpts)
	if err != nil {
		return err
	}

	fmt.Println("Successfully traded auction with id: ", auctionId)

	return nil
}
