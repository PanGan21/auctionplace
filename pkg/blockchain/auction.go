package blockchain

import (
	"context"

	contracts "github.com/PanGan21/auctionplace/contracts/interfaces"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Auction interface {
	GetAuctions(ctx context.Context, client *ethclient.Client) ([]contracts.AuctionPlaceAuction, error)
	GetUserAuctions(ctx context.Context, client *ethclient.Client, address string) ([]contracts.AuctionPlaceAuction, error)
	CreateAuction(ctx context.Context, client *ethclient.Client, createAuctionOpts *CreateAuctionOpts) error
}

type auction struct {
	privateKey      string
	contractAddress string
}

func NewAuctionRunner(privateKey string, contractAddress string) Auction {
	return &auction{
		privateKey:      privateKey,
		contractAddress: contractAddress,
	}
}

func (a *auction) GetAuctions(ctx context.Context, client *ethclient.Client) ([]contracts.AuctionPlaceAuction, error) {
	contract, err := GetContract(ctx, client, a.contractAddress)
	if err != nil {
		return nil, err
	}
	auctions, err := contract.GetAuctions(&bind.CallOpts{Pending: false, Context: ctx})
	if err != nil {
		return nil, err
	}
	return auctions, nil
}

func (a *auction) GetUserAuctions(ctx context.Context, client *ethclient.Client, address string) ([]contracts.AuctionPlaceAuction, error) {
	contract, err := GetContract(ctx, client, a.contractAddress)
	if err != nil {
		return nil, err
	}

	addr, err := GetAddressFromString(address)
	if err != nil {
		return nil, err
	}

	auctions, err := contract.GetUserAuctions(&bind.CallOpts{Pending: false, Context: ctx}, addr)
	if err != nil {
		return nil, err
	}
	return auctions, nil
}

func (a *auction) CreateAuction(ctx context.Context, client *ethclient.Client, createAuctionOpts *CreateAuctionOpts) error {
	contract, err := GetContract(ctx, client, a.contractAddress)
	if err != nil {
		return err
	}

	signer, err := getSigner(ctx, client)
	if err != nil {
		return err
	}
	tx, err := contract.CreateAuction(signer, createAuctionOpts.name, createAuctionOpts.description, createAuctionOpts.min)
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
