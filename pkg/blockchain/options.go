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
