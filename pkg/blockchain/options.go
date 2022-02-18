package blockchain

import "math/big"

type CreateAuctionOpts struct {
	name        string
	description string
	min         *big.Int
}

func NewCreateAuctionOpts(name string, description string, min int64) *CreateAuctionOpts {
	return &CreateAuctionOpts{
		name:        name,
		description: description,
		min:         big.NewInt(min),
	}
}

type CreateOfferOpts struct {
	auctionId *big.Int
	amount    *big.Int
}

func NewCreateOfferOpts(auctionId int64, amount int64) *CreateOfferOpts {
	return &CreateOfferOpts{
		auctionId: big.NewInt(auctionId),
		amount:    big.NewInt(amount),
	}
}

type TradeOpts struct {
	auctionId *big.Int
}

func NewTradeOpts(auctionId int64) *TradeOpts {
	return &TradeOpts{
		auctionId: big.NewInt(auctionId),
	}
}
