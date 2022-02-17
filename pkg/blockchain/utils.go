package blockchain

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"regexp"

	"github.com/PanGan21/auctionplace/config"
	contracts "github.com/PanGan21/auctionplace/contracts/interfaces"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	signer.Value = big.NewInt(config.App.Contract.WeiFounds)
	signer.GasLimit = uint64(config.App.Contract.GasLimit)
	signer.GasPrice = big.NewInt(config.App.Contract.GasPrice)

	return signer, nil
}

func GetContract(ctx context.Context, client *ethclient.Client, address string) (*contracts.AuctionPlace, error) {
	if err := validateAddress(address); err != nil {
		return nil, err
	}
	contract, err := contracts.NewAuctionPlace(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func GetAddressFromString(address string) (common.Address, error) {
	addr := common.HexToAddress(address)
	if err := validateAddress(address); err != nil {
		return addr, err
	}

	return addr, nil
}

func validateAddress(address string) error {
	regex := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if ok := regex.MatchString(address); !ok {
		return ErrInvalidAddress
	}
	return nil
}
