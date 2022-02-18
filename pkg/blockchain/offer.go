package blockchain

import (
	"context"

	contracts "github.com/PanGan21/auctionplace/contracts/interfaces"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Offer interface {
	GetUserOffers(ctx context.Context, client *ethclient.Client, address string) ([]contracts.AuctionPlaceOffer, error)
	CreateOffer(ctx context.Context, client *ethclient.Client, createOfferOpts *CreateOfferOpts) error
}

type offer struct {
	privateKey      string
	contractAddress string
}

func NewOfferRunner(privateKey string, contractAddress string) Offer {
	return &offer{
		privateKey:      privateKey,
		contractAddress: contractAddress,
	}
}

func (o *offer) GetUserOffers(ctx context.Context, client *ethclient.Client, address string) ([]contracts.AuctionPlaceOffer, error) {
	contract, err := GetContract(ctx, client, o.contractAddress)
	if err != nil {
		return nil, err
	}

	addr, err := GetAddressFromString(address)
	if err != nil {
		return nil, err
	}

	offers, err := contract.GetUserOffers(&bind.CallOpts{Pending: false, Context: ctx}, addr)
	if err != nil {
		return nil, err
	}

	return offers, nil
}

func (o *offer) CreateOffer(ctx context.Context, client *ethclient.Client, createOfferOpts *CreateOfferOpts) error {
	contract, err := GetContract(ctx, client, o.contractAddress)
	if err != nil {
		return err
	}

	signer, err := getSigner(ctx, client)
	if err != nil {
		return err
	}

	signer.Value = createOfferOpts.amount

	tx, err := contract.CreateOffer(signer, createOfferOpts.auctionId)
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
