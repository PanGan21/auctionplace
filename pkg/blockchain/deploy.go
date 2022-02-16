package blockchain

import (
	"context"

	contracts "github.com/PanGan21/auctionplace/contracts/interfaces"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Deployer interface
type Deployer interface {
	Deploy(ctx context.Context, client *ethclient.Client) error
	ContractAddress() string
}

type deployer struct {
	address     common.Address
	transaction *types.Transaction
	contract    *contracts.AuctionPlace
}

// NewDepoyer returns a new Deployer instance
func NewDeployer() Deployer {
	return &deployer{}
}

func (d *deployer) Deploy(ctx context.Context, client *ethclient.Client) error {
	signer, err := getSigner(ctx, client)
	if err != nil {
		return err
	}

	address, tx, contract, err := contracts.DeployAuctionPlace(signer, client)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(ctx, client, tx)
	if err != nil {
		return err
	}

	d.address = address
	d.contract = contract
	d.transaction = tx
	return nil
}

// ContractAddress returns the contract address
func (d *deployer) ContractAddress() string {
	return d.address.Hex()
}
