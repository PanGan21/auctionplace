package blockchain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/PanGan21/auctionplace/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrInvalidKey             = errors.New("invalid key")
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidContractAddress = errors.New("invalid contract address")
)

func getSigner(ctx context.Context, client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(config.App.Blockchain.PrivateKey)
	if err != nil {
		return nil, err
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, ErrInvalidKey
	}
	address := crypto.PubkeyToAddress(*publicKey)
	nonce, err := client.PendingNonceAt(ctx, address)
	if err != nil {
		return nil, err
	}
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	signer, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	signer.Nonce = big.NewInt(int64(nonce))
	signer.Value = big.NewInt(config.App.ContractConfig.WeiFounds)
	signer.GasLimit = uint64(config.App.ContractConfig.GasLimit)
	signer.GasPrice = big.NewInt(config.App.ContractConfig.GasPrice)

	return signer, nil
}
