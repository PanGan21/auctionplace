package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Trade interface {
	TradeAuctionToOffer(ctx context.Context, client *ethclient.Client, tradeOpts *TradeOpts) error
}

type trade struct {
	privateKey      string
	contractAddress string
}

func NewTradeRunner(privateKey string, contractAddress string) Trade {
	return &trade{
		privateKey:      privateKey,
		contractAddress: contractAddress,
	}
}

func (t *trade) TradeAuctionToOffer(ctx context.Context, client *ethclient.Client, tradeOpts *TradeOpts) error {
	contract, err := GetContract(ctx, client, t.contractAddress)
	if err != nil {
		return err
	}

	signer, err := getSigner(ctx, client)
	if err != nil {
		return err
	}

	tx, err := contract.Trade(signer, tradeOpts.auctionId)
	if err != nil {
		return err
	}

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful || err != nil {
		return err
	}

	return nil
}
