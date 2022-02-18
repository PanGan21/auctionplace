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
var AuctionPlaceBin = "0x60806040526001600455600160055534801561001a57600080fd5b506111268061002a6000396000f3fe6080604052600436106100555760003560e01c8063661d2c2c1461005a578063a269993d14610090578063d7c06919146100a5578063ddf20706146100c7578063df1dd826146100e7578063ff3ad0b414610107575b600080fd5b34801561006657600080fd5b5061007a610075366004610db9565b610127565b6040516100879190610fc2565b60405180910390f35b6100a361009e366004610e5d565b610262565b005b3480156100b157600080fd5b506100ba610418565b6040516100879190610ec3565b3480156100d357600080fd5b506100a36100e2366004610de9565b6106a6565b3480156100f357600080fd5b506100a3610102366004610e5d565b610874565b34801561011357600080fd5b506100ba610122366004610db9565b6109d3565b6001600160a01b038116600090815260036020526040812080546060929067ffffffffffffffff81111561015d5761015d6110da565b6040519080825280602002602001820160405280156101c257816020015b6101af6040518060800160405280600081526020016000815260200160006001600160a01b03168152602001600081525090565b81526020019060019003908161017b5790505b50905060005b825481101561025a57600160006101df838361102f565b8152602080820192909252604090810160002081516080810183528154815260018201549381019390935260028101546001600160a01b031691830191909152600301546060820152825183908390811061023c5761023c6110c4565b6020026020010181905250808061025290611093565b9150506101c8565b509392505050565b80600081118015610274575060045481105b6102be5760405162461bcd60e51b8152602060048201526016602482015275105d58dd1a5bdb88191bd95cc81b9bdd08195e1a5cdd60521b60448201526064015b60405180910390fd5b600082815260208181526040808320600581015484526001909252909120600482015434108015906102f35750806003015434115b6103585760405162461bcd60e51b815260206004820152603060248201527f6d73672e76616c7565206d7573742062652067726561746572207468616e206d60448201526f34b71030b732103132b9ba27b33332b960811b60648201526084016102b5565b6005805483820181905560068401805460018082018355600092835260208084209092019390935560408051608081018252855480825288548285019081523383850181815234606086019081529388528887528588209451855591518489015590516002840180546001600160a01b0319166001600160a01b039092169190911790559051600392830155845282528220845481549485018255908352908220909201919091558154919061040d83611093565b919050555050505050565b60606000600160045461042b9190611047565b67ffffffffffffffff811115610443576104436110da565b60405190808252806020026020018201604052801561047c57816020015b610469610c57565b8152602001906001900390816104615790505b5090506004546001141561048f57919050565b60015b6004548110156106a05760008181526020818152604091829020825160e0810184528154815260018201546001600160a01b03169281019290925260028101805492939192918401916104e49061105e565b80601f01602080910402602001604051908101604052809291908181526020018280546105109061105e565b801561055d5780601f106105325761010080835404028352916020019161055d565b820191906000526020600020905b81548152906001019060200180831161054057829003601f168201915b505050505081526020016003820180546105769061105e565b80601f01602080910402602001604051908101604052809291908181526020018280546105a29061105e565b80156105ef5780601f106105c4576101008083540402835291602001916105ef565b820191906000526020600020905b8154815290600101906020018083116105d257829003601f168201915b5050505050815260200160048201548152602001600582015481526020016006820180548060200260200160405190810160405280929190818152602001828054801561065b57602002820191906000526020600020905b815481526020019060010190808311610647575b505050505081525050826001836106729190611047565b81518110610682576106826110c4565b6020026020010181905250808061069890611093565b915050610492565b50919050565b600081116106e95760405162461bcd60e51b815260206004820152601060248201526f05f6d696e206d757374206265203e20360841b60448201526064016102b5565b604080516000815261010081018252600454602080830191825233838501528351601f8901829004820281018201909452878452919290916060840191908990899081908401838280828437600092019190915250505090825250604080516020601f880181900481028201810190925286815291810191908790879081908401838280828437600092018290525093855250505060208083018690526040808401839052606090930185905260045482528181529082902083518155838201516001820180546001600160a01b0319166001600160a01b039092169190911790559183015180516107e19260028501920190610c9d565b50606082015180516107fd916003840191602090910190610c9d565b506080820151600482015560a0820151600582015560c0820151805161082d916006840191602090910190610d21565b5050336000908152600260209081526040822060048054825460018101845592855292842090910191909155805492509061086783611093565b9190505550505050505050565b80600081118015610886575060045481105b6108cb5760405162461bcd60e51b8152602060048201526016602482015275105d58dd1a5bdb88191bd95cc81b9bdd08195e1a5cdd60521b60448201526064016102b5565b600082815260208181526040808320600581015484526001909252822090915b600683015481101561098b57600083600601828154811061090e5761090e6110c4565b90600052602060002001549050836005015481146109785760008181526001602052604080822060028101546003820154925191936001600160a01b039091169280156108fc02929091818181858888f19350505050158015610975573d6000803e3d6000fd5b50505b508061098381611093565b9150506108eb565b50600182015460038201546040516001600160a01b039092169181156108fc0291906000818181858888f193505050501580156109cc573d6000803e3d6000fd5b5050505050565b6001600160a01b038116600090815260026020526040812080546060929067ffffffffffffffff811115610a0957610a096110da565b604051908082528060200260200182016040528015610a4257816020015b610a2f610c57565b815260200190600190039081610a275790505b50905060005b825481101561025a57600080610a5f83600161102f565b81526020808201929092526040908101600020815160e0810183528154815260018201546001600160a01b0316938101939093526002810180549192840191610aa79061105e565b80601f0160208091040260200160405190810160405280929190818152602001828054610ad39061105e565b8015610b205780601f10610af557610100808354040283529160200191610b20565b820191906000526020600020905b815481529060010190602001808311610b0357829003601f168201915b50505050508152602001600382018054610b399061105e565b80601f0160208091040260200160405190810160405280929190818152602001828054610b659061105e565b8015610bb25780601f10610b8757610100808354040283529160200191610bb2565b820191906000526020600020905b815481529060010190602001808311610b9557829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201805480602002602001604051908101604052809291908181526020018280548015610c1e57602002820191906000526020600020905b815481526020019060010190808311610c0a575b505050505081525050828281518110610c3957610c396110c4565b60200260200101819052508080610c4f90611093565b915050610a48565b6040518060e001604052806000815260200160006001600160a01b0316815260200160608152602001606081526020016000815260200160008152602001606081525090565b828054610ca99061105e565b90600052602060002090601f016020900481019282610ccb5760008555610d11565b82601f10610ce457805160ff1916838001178555610d11565b82800160010185558215610d11579182015b82811115610d11578251825591602001919060010190610cf6565b50610d1d929150610d5b565b5090565b828054828255906000526020600020908101928215610d115791602002820182811115610d11578251825591602001919060010190610cf6565b5b80821115610d1d5760008155600101610d5c565b60008083601f840112610d8257600080fd5b50813567ffffffffffffffff811115610d9a57600080fd5b602083019150836020828501011115610db257600080fd5b9250929050565b600060208284031215610dcb57600080fd5b81356001600160a01b0381168114610de257600080fd5b9392505050565b600080600080600060608688031215610e0157600080fd5b853567ffffffffffffffff80821115610e1957600080fd5b610e2589838a01610d70565b90975095506020880135915080821115610e3e57600080fd5b50610e4b88828901610d70565b96999598509660400135949350505050565b600060208284031215610e6f57600080fd5b5035919050565b6000815180845260005b81811015610e9c57602081850181015186830182015201610e80565b81811115610eae576000602083870101525b50601f01601f19169290920160200192915050565b60006020808301818452808551808352604092508286019150828160051b8701018488016000805b84811015610fb357898403603f19018652825180518552888101516001600160a01b0316898601528781015160e089870181905290610f2c82880182610e76565b91505060608083015187830382890152610f468382610e76565b608085810151908a015260a080860151908a015260c094850151898203958a01959095528451808252948d0194879450908d01925090505b80831015610f9e5783518252928b019260019290920191908b0190610f7e565b50978a01979550505091870191600101610eeb565b50919998505050505050505050565b602080825282518282018190526000919060409081850190868401855b82811015611022578151805185528681015187860152858101516001600160a01b0316868601526060908101519085015260809093019290850190600101610fdf565b5091979650505050505050565b60008219821115611042576110426110ae565b500190565b600082821015611059576110596110ae565b500390565b600181811c9082168061107257607f821691505b602082108114156106a057634e487b7160e01b600052602260045260246000fd5b60006000198214156110a7576110a76110ae565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220f45801b1d75ee7c331c354a8a322de0bf4a5130455e4b95cb4c5a99c55a4447464736f6c63430008060033"

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
