package blockchain

import (
	"context"

	contracts "github.com/PanGan21/auctionplace/contracts/interfaces"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Auction interface {
	GetAuctions(ctx context.Context, client *ethclient.Client) ([]contracts.AuctionPlaceAuction, error)
	GetUserAuctions(ctx context.Context, client *ethclient.Client, address string) ([]contracts.AuctionPlaceAuction, error)
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
