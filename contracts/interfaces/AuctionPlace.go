// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AuctionPlaceAuction is an auto generated low-level Go binding around an user-defined struct.
type AuctionPlaceAuction struct {
	Id          *big.Int
	Seller      common.Address
	Name        string
	Description string
	Min         *big.Int
	BestOfferId *big.Int
	OfferIds    []*big.Int
}

// AuctionPlaceOffer is an auto generated low-level Go binding around an user-defined struct.
type AuctionPlaceOffer struct {
	Id        *big.Int
	AuctionId *big.Int
	Buyer     common.Address
	Price     *big.Int
}

// AuctionPlaceABI is the input ABI used to generate the binding from.
const AuctionPlaceABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuctions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bestOfferId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"offerIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structAuctionPlace.Auction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserAuctions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bestOfferId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"offerIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structAuctionPlace.Auction[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserOffers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structAuctionPlace.Offer[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AuctionPlaceFuncSigs maps the 4-byte function signature to its string representation.
var AuctionPlaceFuncSigs = map[string]string{
	"ddf20706": "createAuction(string,string,uint256)",
	"a269993d": "createOffer(uint256)",
	"d7c06919": "getAuctions()",
	"ff3ad0b4": "getUserAuctions(address)",
	"661d2c2c": "getUserOffers(address)",
	"df1dd826": "trade(uint256)",
}

// AuctionPlaceBin is the compiled bytecode used for deploying new contracts.
var AuctionPlaceBin = "0x60806040526001600455600160055534801561001a57600080fd5b506111068061002a6000396000f3fe6080604052600436106100555760003560e01c8063661d2c2c1461005a578063a269993d14610090578063d7c06919146100a5578063ddf20706146100c7578063df1dd826146100e7578063ff3ad0b414610107575b600080fd5b34801561006657600080fd5b5061007a610075366004610db1565b610127565b6040516100879190610fba565b60405180910390f35b6100a361009e366004610e55565b61025c565b005b3480156100b157600080fd5b506100ba610412565b6040516100879190610ebb565b3480156100d357600080fd5b506100a36100e2366004610de1565b6106a0565b3480156100f357600080fd5b506100a3610102366004610e55565b61086e565b34801561011357600080fd5b506100ba610122366004610db1565b6109cd565b6001600160a01b038116600090815260036020526040812080546060929067ffffffffffffffff81111561015d5761015d6110ba565b6040519080825280602002602001820160405280156101c257816020015b6101af6040518060800160405280600081526020016000815260200160006001600160a01b03168152602001600081525090565b81526020019060019003908161017b5790505b50905060015b825481101561025457600081815260016020818152604092839020835160808101855281548152818401549281019290925260028101546001600160a01b031693820193909352600390920154606083015283906102269084611027565b81518110610236576102366110a4565b6020026020010181905250808061024c90611073565b9150506101c8565b509392505050565b8060008111801561026e575060045481105b6102b85760405162461bcd60e51b8152602060048201526016602482015275105d58dd1a5bdb88191bd95cc81b9bdd08195e1a5cdd60521b60448201526064015b60405180910390fd5b600082815260208181526040808320600581015484526001909252909120600482015434108015906102ed5750806003015434115b6103525760405162461bcd60e51b815260206004820152603060248201527f6d73672e76616c7565206d7573742062652067726561746572207468616e206d60448201526f34b71030b732103132b9ba27b33332b960811b60648201526084016102af565b6005805483820181905560068401805460018082018355600092835260208084209092019390935560408051608081018252855480825288548285019081523383850181815234606086019081529388528887528588209451855591518489015590516002840180546001600160a01b0319166001600160a01b039092169190911790559051600392830155845282528220845481549485018255908352908220909201919091558154919061040783611073565b919050555050505050565b6060600060016004546104259190611027565b67ffffffffffffffff81111561043d5761043d6110ba565b60405190808252806020026020018201604052801561047657816020015b610463610c4f565b81526020019060019003908161045b5790505b5090506004546001141561048957919050565b60015b60045481101561069a5760008181526020818152604091829020825160e0810184528154815260018201546001600160a01b03169281019290925260028101805492939192918401916104de9061103e565b80601f016020809104026020016040519081016040528092919081815260200182805461050a9061103e565b80156105575780601f1061052c57610100808354040283529160200191610557565b820191906000526020600020905b81548152906001019060200180831161053a57829003601f168201915b505050505081526020016003820180546105709061103e565b80601f016020809104026020016040519081016040528092919081815260200182805461059c9061103e565b80156105e95780601f106105be576101008083540402835291602001916105e9565b820191906000526020600020905b8154815290600101906020018083116105cc57829003601f168201915b5050505050815260200160048201548152602001600582015481526020016006820180548060200260200160405190810160405280929190818152602001828054801561065557602002820191906000526020600020905b815481526020019060010190808311610641575b5050505050815250508260018361066c9190611027565b8151811061067c5761067c6110a4565b6020026020010181905250808061069290611073565b91505061048c565b50919050565b600081116106e35760405162461bcd60e51b815260206004820152601060248201526f05f6d696e206d757374206265203e20360841b60448201526064016102af565b604080516000815261010081018252600454602080830191825233838501528351601f8901829004820281018201909452878452919290916060840191908990899081908401838280828437600092019190915250505090825250604080516020601f880181900481028201810190925286815291810191908790879081908401838280828437600092018290525093855250505060208083018690526040808401839052606090930185905260045482528181529082902083518155838201516001820180546001600160a01b0319166001600160a01b039092169190911790559183015180516107db9260028501920190610c95565b50606082015180516107f7916003840191602090910190610c95565b506080820151600482015560a0820151600582015560c08201518051610827916006840191602090910190610d19565b5050336000908152600260209081526040822060048054825460018101845592855292842090910191909155805492509061086183611073565b9190505550505050505050565b80600081118015610880575060045481105b6108c55760405162461bcd60e51b8152602060048201526016602482015275105d58dd1a5bdb88191bd95cc81b9bdd08195e1a5cdd60521b60448201526064016102af565b600082815260208181526040808320600581015484526001909252822090915b6006830154811015610985576000836006018281548110610908576109086110a4565b90600052602060002001549050836005015481146109725760008181526001602052604080822060028101546003820154925191936001600160a01b039091169280156108fc02929091818181858888f1935050505015801561096f573d6000803e3d6000fd5b50505b508061097d81611073565b9150506108e5565b50600182015460038201546040516001600160a01b039092169181156108fc0291906000818181858888f193505050501580156109c6573d6000803e3d6000fd5b5050505050565b6001600160a01b038116600090815260026020526040812080546060929067ffffffffffffffff811115610a0357610a036110ba565b604051908082528060200260200182016040528015610a3c57816020015b610a29610c4f565b815260200190600190039081610a215790505b50905060015b82548110156102545760008181526020818152604091829020825160e0810184528154815260018201546001600160a01b0316928101929092526002810180549293919291840191610a939061103e565b80601f0160208091040260200160405190810160405280929190818152602001828054610abf9061103e565b8015610b0c5780601f10610ae157610100808354040283529160200191610b0c565b820191906000526020600020905b815481529060010190602001808311610aef57829003601f168201915b50505050508152602001600382018054610b259061103e565b80601f0160208091040260200160405190810160405280929190818152602001828054610b519061103e565b8015610b9e5780601f10610b7357610100808354040283529160200191610b9e565b820191906000526020600020905b815481529060010190602001808311610b8157829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201805480602002602001604051908101604052809291908181526020018280548015610c0a57602002820191906000526020600020905b815481526020019060010190808311610bf6575b50505050508152505082600183610c219190611027565b81518110610c3157610c316110a4565b60200260200101819052508080610c4790611073565b915050610a42565b6040518060e001604052806000815260200160006001600160a01b0316815260200160608152602001606081526020016000815260200160008152602001606081525090565b828054610ca19061103e565b90600052602060002090601f016020900481019282610cc35760008555610d09565b82601f10610cdc57805160ff1916838001178555610d09565b82800160010185558215610d09579182015b82811115610d09578251825591602001919060010190610cee565b50610d15929150610d53565b5090565b828054828255906000526020600020908101928215610d095791602002820182811115610d09578251825591602001919060010190610cee565b5b80821115610d155760008155600101610d54565b60008083601f840112610d7a57600080fd5b50813567ffffffffffffffff811115610d9257600080fd5b602083019150836020828501011115610daa57600080fd5b9250929050565b600060208284031215610dc357600080fd5b81356001600160a01b0381168114610dda57600080fd5b9392505050565b600080600080600060608688031215610df957600080fd5b853567ffffffffffffffff80821115610e1157600080fd5b610e1d89838a01610d68565b90975095506020880135915080821115610e3657600080fd5b50610e4388828901610d68565b96999598509660400135949350505050565b600060208284031215610e6757600080fd5b5035919050565b6000815180845260005b81811015610e9457602081850181015186830182015201610e78565b81811115610ea6576000602083870101525b50601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b8701018488016000805b84811015610fab57898403603f19018652825180518552888101516001600160a01b0316898601528781015160e089870181905290610f2482880182610e6e565b91505060608083015187830382890152610f3e8382610e6e565b608085810151908a015260a080860151908a015260c094850151898203958a01959095528451808252948d0194879450908d01925090505b80831015610f965783518252928b019260019290920191908b0190610f76565b50978a01979550505091870191600101610ee3565b50919998505050505050505050565b602080825282518282018190526000919060409081850190868401855b8281101561101a578151805185528681015187860152858101516001600160a01b0316868601526060908101519085015260809093019290850190600101610fd7565b5091979650505050505050565b6000828210156110395761103961108e565b500390565b600181811c9082168061105257607f821691505b6020821081141561069a57634e487b7160e01b600052602260045260246000fd5b60006000198214156110875761108761108e565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212201f9f289670f0ba90b9f977f524cc3e19a1379342a24295ffedc1109d7544902c64736f6c63430008060033"

// DeployAuctionPlace deploys a new Ethereum contract, binding an instance of AuctionPlace to it.
func DeployAuctionPlace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AuctionPlace, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionPlaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuctionPlaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuctionPlace{AuctionPlaceCaller: AuctionPlaceCaller{contract: contract}, AuctionPlaceTransactor: AuctionPlaceTransactor{contract: contract}, AuctionPlaceFilterer: AuctionPlaceFilterer{contract: contract}}, nil
}

// AuctionPlace is an auto generated Go binding around an Ethereum contract.
type AuctionPlace struct {
	AuctionPlaceCaller     // Read-only binding to the contract
	AuctionPlaceTransactor // Write-only binding to the contract
	AuctionPlaceFilterer   // Log filterer for contract events
}

// AuctionPlaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionPlaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionPlaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionPlaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionPlaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionPlaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionPlaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionPlaceSession struct {
	Contract     *AuctionPlace     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionPlaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionPlaceCallerSession struct {
	Contract *AuctionPlaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AuctionPlaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionPlaceTransactorSession struct {
	Contract     *AuctionPlaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AuctionPlaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionPlaceRaw struct {
	Contract *AuctionPlace // Generic contract binding to access the raw methods on
}

// AuctionPlaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionPlaceCallerRaw struct {
	Contract *AuctionPlaceCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionPlaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionPlaceTransactorRaw struct {
	Contract *AuctionPlaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionPlace creates a new instance of AuctionPlace, bound to a specific deployed contract.
func NewAuctionPlace(address common.Address, backend bind.ContractBackend) (*AuctionPlace, error) {
	contract, err := bindAuctionPlace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionPlace{AuctionPlaceCaller: AuctionPlaceCaller{contract: contract}, AuctionPlaceTransactor: AuctionPlaceTransactor{contract: contract}, AuctionPlaceFilterer: AuctionPlaceFilterer{contract: contract}}, nil
}

// NewAuctionPlaceCaller creates a new read-only instance of AuctionPlace, bound to a specific deployed contract.
func NewAuctionPlaceCaller(address common.Address, caller bind.ContractCaller) (*AuctionPlaceCaller, error) {
	contract, err := bindAuctionPlace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionPlaceCaller{contract: contract}, nil
}

// NewAuctionPlaceTransactor creates a new write-only instance of AuctionPlace, bound to a specific deployed contract.
func NewAuctionPlaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionPlaceTransactor, error) {
	contract, err := bindAuctionPlace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionPlaceTransactor{contract: contract}, nil
}

// NewAuctionPlaceFilterer creates a new log filterer instance of AuctionPlace, bound to a specific deployed contract.
func NewAuctionPlaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionPlaceFilterer, error) {
	contract, err := bindAuctionPlace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionPlaceFilterer{contract: contract}, nil
}

// bindAuctionPlace binds a generic wrapper to an already deployed contract.
func bindAuctionPlace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionPlaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionPlace *AuctionPlaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionPlace.Contract.AuctionPlaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionPlace *AuctionPlaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionPlace.Contract.AuctionPlaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionPlace *AuctionPlaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionPlace.Contract.AuctionPlaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionPlace *AuctionPlaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionPlace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionPlace *AuctionPlaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionPlace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionPlace *AuctionPlaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionPlace.Contract.contract.Transact(opts, method, params...)
}

// GetAuctions is a free data retrieval call binding the contract method 0xd7c06919.
//
// Solidity: function getAuctions() view returns((uint256,address,string,string,uint256,uint256,uint256[])[])
func (_AuctionPlace *AuctionPlaceCaller) GetAuctions(opts *bind.CallOpts) ([]AuctionPlaceAuction, error) {
	var out []interface{}
	err := _AuctionPlace.contract.Call(opts, &out, "getAuctions")

	if err != nil {
		return *new([]AuctionPlaceAuction), err
	}

	out0 := *abi.ConvertType(out[0], new([]AuctionPlaceAuction)).(*[]AuctionPlaceAuction)

	return out0, err

}

// GetAuctions is a free data retrieval call binding the contract method 0xd7c06919.
//
// Solidity: function getAuctions() view returns((uint256,address,string,string,uint256,uint256,uint256[])[])
func (_AuctionPlace *AuctionPlaceSession) GetAuctions() ([]AuctionPlaceAuction, error) {
	return _AuctionPlace.Contract.GetAuctions(&_AuctionPlace.CallOpts)
}

// GetAuctions is a free data retrieval call binding the contract method 0xd7c06919.
//
// Solidity: function getAuctions() view returns((uint256,address,string,string,uint256,uint256,uint256[])[])
func (_AuctionPlace *AuctionPlaceCallerSession) GetAuctions() ([]AuctionPlaceAuction, error) {
	return _AuctionPlace.Contract.GetAuctions(&_AuctionPlace.CallOpts)
}

// GetUserAuctions is a free data retrieval call binding the contract method 0xff3ad0b4.
//
// Solidity: function getUserAuctions(address _user) view returns((uint256,address,string,string,uint256,uint256,uint256[])[])
func (_AuctionPlace *AuctionPlaceCaller) GetUserAuctions(opts *bind.CallOpts, _user common.Address) ([]AuctionPlaceAuction, error) {
	var out []interface{}
	err := _AuctionPlace.contract.Call(opts, &out, "getUserAuctions", _user)

	if err != nil {
		return *new([]AuctionPlaceAuction), err
	}

	out0 := *abi.ConvertType(out[0], new([]AuctionPlaceAuction)).(*[]AuctionPlaceAuction)

	return out0, err

}

// GetUserAuctions is a free data retrieval call binding the contract method 0xff3ad0b4.
//
// Solidity: function getUserAuctions(address _user) view returns((uint256,address,string,string,uint256,uint256,uint256[])[])
func (_AuctionPlace *AuctionPlaceSession) GetUserAuctions(_user common.Address) ([]AuctionPlaceAuction, error) {
	return _AuctionPlace.Contract.GetUserAuctions(&_AuctionPlace.CallOpts, _user)
}

// GetUserAuctions is a free data retrieval call binding the contract method 0xff3ad0b4.
//
// Solidity: function getUserAuctions(address _user) view returns((uint256,address,string,string,uint256,uint256,uint256[])[])
func (_AuctionPlace *AuctionPlaceCallerSession) GetUserAuctions(_user common.Address) ([]AuctionPlaceAuction, error) {
	return _AuctionPlace.Contract.GetUserAuctions(&_AuctionPlace.CallOpts, _user)
}

// GetUserOffers is a free data retrieval call binding the contract method 0x661d2c2c.
//
// Solidity: function getUserOffers(address _user) view returns((uint256,uint256,address,uint256)[])
func (_AuctionPlace *AuctionPlaceCaller) GetUserOffers(opts *bind.CallOpts, _user common.Address) ([]AuctionPlaceOffer, error) {
	var out []interface{}
	err := _AuctionPlace.contract.Call(opts, &out, "getUserOffers", _user)

	if err != nil {
		return *new([]AuctionPlaceOffer), err
	}

	out0 := *abi.ConvertType(out[0], new([]AuctionPlaceOffer)).(*[]AuctionPlaceOffer)

	return out0, err

}

// GetUserOffers is a free data retrieval call binding the contract method 0x661d2c2c.
//
// Solidity: function getUserOffers(address _user) view returns((uint256,uint256,address,uint256)[])
func (_AuctionPlace *AuctionPlaceSession) GetUserOffers(_user common.Address) ([]AuctionPlaceOffer, error) {
	return _AuctionPlace.Contract.GetUserOffers(&_AuctionPlace.CallOpts, _user)
}

// GetUserOffers is a free data retrieval call binding the contract method 0x661d2c2c.
//
// Solidity: function getUserOffers(address _user) view returns((uint256,uint256,address,uint256)[])
func (_AuctionPlace *AuctionPlaceCallerSession) GetUserOffers(_user common.Address) ([]AuctionPlaceOffer, error) {
	return _AuctionPlace.Contract.GetUserOffers(&_AuctionPlace.CallOpts, _user)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xddf20706.
//
// Solidity: function createAuction(string _name, string _description, uint256 _min) returns()
func (_AuctionPlace *AuctionPlaceTransactor) CreateAuction(opts *bind.TransactOpts, _name string, _description string, _min *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.contract.Transact(opts, "createAuction", _name, _description, _min)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xddf20706.
//
// Solidity: function createAuction(string _name, string _description, uint256 _min) returns()
func (_AuctionPlace *AuctionPlaceSession) CreateAuction(_name string, _description string, _min *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.Contract.CreateAuction(&_AuctionPlace.TransactOpts, _name, _description, _min)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xddf20706.
//
// Solidity: function createAuction(string _name, string _description, uint256 _min) returns()
func (_AuctionPlace *AuctionPlaceTransactorSession) CreateAuction(_name string, _description string, _min *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.Contract.CreateAuction(&_AuctionPlace.TransactOpts, _name, _description, _min)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xa269993d.
//
// Solidity: function createOffer(uint256 _auctionId) payable returns()
func (_AuctionPlace *AuctionPlaceTransactor) CreateOffer(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.contract.Transact(opts, "createOffer", _auctionId)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xa269993d.
//
// Solidity: function createOffer(uint256 _auctionId) payable returns()
func (_AuctionPlace *AuctionPlaceSession) CreateOffer(_auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.Contract.CreateOffer(&_AuctionPlace.TransactOpts, _auctionId)
}

// CreateOffer is a paid mutator transaction binding the contract method 0xa269993d.
//
// Solidity: function createOffer(uint256 _auctionId) payable returns()
func (_AuctionPlace *AuctionPlaceTransactorSession) CreateOffer(_auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.Contract.CreateOffer(&_AuctionPlace.TransactOpts, _auctionId)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(uint256 _auctionId) returns()
func (_AuctionPlace *AuctionPlaceTransactor) Trade(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.contract.Transact(opts, "trade", _auctionId)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(uint256 _auctionId) returns()
func (_AuctionPlace *AuctionPlaceSession) Trade(_auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.Contract.Trade(&_AuctionPlace.TransactOpts, _auctionId)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(uint256 _auctionId) returns()
func (_AuctionPlace *AuctionPlaceTransactorSession) Trade(_auctionId *big.Int) (*types.Transaction, error) {
	return _AuctionPlace.Contract.Trade(&_AuctionPlace.TransactOpts, _auctionId)
}
