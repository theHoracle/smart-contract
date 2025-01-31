// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bankapi

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BankapiMetaData contains all meta data concerning the Bankapi contract.
var BankapiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600581526020017f302e312e30000000000000000000000000000000000000000000000000000000815250600190816100549190610297565b50610366565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100d557607f821691505b6020821081036100e8576100e7610091565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261014a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261010f565b610154868361010f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61019861019361018e8461016c565b610175565b61016c565b9050919050565b5f819050919050565b6101b18361017e565b6101c56101bd8261019f565b84845461011b565b825550505050565b5f5f905090565b6101dc6101cd565b6101e78184846101a8565b505050565b5b8181101561020a576101ff5f826101d4565b6001810190506101ed565b5050565b601f82111561024f57610220816100ee565b61022984610100565b81016020851015610238578190505b61024c61024485610100565b8301826101ec565b50505b505050565b5f82821c905092915050565b5f61026f5f1984600802610254565b1980831691505092915050565b5f6102878383610260565b9150826002028217905092915050565b6102a08261005a565b67ffffffffffffffff8111156102b9576102b8610064565b5b6102c382546100be565b6102ce82828561020e565b5f60209050601f8311600181146102ff575f84156102ed578287015190505b6102f7858261027c565b86555061035e565b601f19841661030d866100ee565b5f5b828110156103345784890151825560018201915060208501945060208101905061030f565b86831015610351578489015161034d601f891682610260565b8355505b6001600288020188555050505b505050505050565b610f8f806103735f395ff3fe608060405260043610610049575f3560e01c806357ea89b61461004d5780637d7b009914610057578063b4a99a4e14610081578063bb62860d146100ab578063ed21248c146100d5575b5f5ffd5b6100556100df565b005b348015610062575f5ffd5b5061006b610295565b604051610078919061081f565b60405180910390f35b34801561008c575f5ffd5b506100956102b9565b6040516100a2919061081f565b60405180910390f35b3480156100b6575f5ffd5b506100bf6102de565b6040516100cc91906108a8565b60405180910390f35b6100dd61036a565b005b5f3390505f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205403610162576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015990610912565b60405180910390fd5b5f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490508173ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156101e6573d5f5f3e3d5ffd5b505f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a61025333610464565b61025c8361061d565b60405160200161026d9291906109dc565b60405160208183030381529060405260405161028991906108a8565b60405180910390a15050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546102eb90610a59565b80601f016020809104026020016040519081016040528092919081815260200182805461031790610a59565b80156103625780601f1061033957610100808354040283529160200191610362565b820191905f5260205f20905b81548152906001019060200180831161034557829003601f168201915b505050505081565b3460035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546103b69190610abf565b925050819055507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6103e733610464565b61042d60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205461061d565b60405160200161043e929190610b3e565b60405160208183030381529060405260405161045a91906108a8565b60405180910390a1565b60605f602867ffffffffffffffff81111561048257610481610b8e565b5b6040519080825280601f01601f1916602001820160405280156104b45781602001600182028036833780820191505090505b5090505f5f90505b6014811015610613575f8160136104d39190610bbb565b60086104df9190610bee565b60026104eb9190610d5e565b8573ffffffffffffffffffffffffffffffffffffffff1661050c9190610dd5565b60f81b90505f60108260f81c6105229190610e11565b60f81b90505f8160f81c60106105389190610e41565b8360f81c6105469190610e7d565b60f81b90506105548261079b565b858560026105629190610bee565b8151811061057357610572610eb1565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053506105aa8161079b565b8560018660026105ba9190610bee565b6105c49190610abf565b815181106105d5576105d4610eb1565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535050505080806001019150506104bc565b5080915050919050565b60605f8203610663576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050610796565b5f8290505f5b5f821461069257808061067b90610ede565b915050600a8261068b9190610dd5565b9150610669565b5f8167ffffffffffffffff8111156106ad576106ac610b8e565b5b6040519080825280601f01601f1916602001820160405280156106df5781602001600182028036833780820191505090505b5090505f8290505b5f861461078e576001816106fb9190610bbb565b90505f600a808861070c9190610dd5565b6107169190610bee565b876107219190610bbb565b603061072d9190610f25565b90505f8160f81b90508084848151811061074a57610749610eb1565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600a886107859190610dd5565b975050506106e7565b819450505050505b919050565b5f600a8260f81c60ff1610156107c55760308260f81c6107bb9190610f25565b60f81b90506107db565b60578260f81c6107d59190610f25565b60f81b90505b919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610809826107e0565b9050919050565b610819816107ff565b82525050565b5f6020820190506108325f830184610810565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61087a82610838565b6108848185610842565b9350610894818560208601610852565b61089d81610860565b840191505092915050565b5f6020820190508181035f8301526108c08184610870565b905092915050565b7f6e6f7420656e6f7567682062616c616e636500000000000000000000000000005f82015250565b5f6108fc601283610842565b9150610907826108c8565b602082019050919050565b5f6020820190508181035f830152610929816108f0565b9050919050565b7f77697468647261775b0000000000000000000000000000000000000000000000815250565b5f81905092915050565b5f61096a82610838565b6109748185610956565b9350610984818560208601610852565b80840191505092915050565b7f5d20616d6f756e745b0000000000000000000000000000000000000000000000815250565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b5f6109e682610930565b6009820191506109f68285610960565b9150610a0182610990565b600982019150610a118284610960565b9150610a1c826109b6565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610a7057607f821691505b602082108103610a8357610a82610a2c565b5b50919050565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610ac982610a89565b9150610ad483610a89565b9250828201905080821115610aec57610aeb610a92565b5b92915050565b7f6465706f7369745b000000000000000000000000000000000000000000000000815250565b7f5d2062616c616e63655b00000000000000000000000000000000000000000000815250565b5f610b4882610af2565b600882019150610b588285610960565b9150610b6382610b18565b600a82019150610b738284610960565b9150610b7e826109b6565b6001820191508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f610bc582610a89565b9150610bd083610a89565b9250828203905081811115610be857610be7610a92565b5b92915050565b5f610bf882610a89565b9150610c0383610a89565b9250828202610c1181610a89565b91508282048414831517610c2857610c27610a92565b5b5092915050565b5f8160011c9050919050565b5f5f8291508390505b6001851115610c8457808604811115610c6057610c5f610a92565b5b6001851615610c6f5780820291505b8081029050610c7d85610c2f565b9450610c44565b94509492505050565b5f82610c9c5760019050610d57565b81610ca9575f9050610d57565b8160018114610cbf5760028114610cc957610cf8565b6001915050610d57565b60ff841115610cdb57610cda610a92565b5b8360020a915084821115610cf257610cf1610a92565b5b50610d57565b5060208310610133831016604e8410600b8410161715610d2d5782820a905083811115610d2857610d27610a92565b5b610d57565b610d3a8484846001610c3b565b92509050818404811115610d5157610d50610a92565b5b81810290505b9392505050565b5f610d6882610a89565b9150610d7383610a89565b9250610da07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484610c8d565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610ddf82610a89565b9150610dea83610a89565b925082610dfa57610df9610da8565b5b828204905092915050565b5f60ff82169050919050565b5f610e1b82610e05565b9150610e2683610e05565b925082610e3657610e35610da8565b5b828204905092915050565b5f610e4b82610e05565b9150610e5683610e05565b9250828202610e6481610e05565b9150808214610e7657610e75610a92565b5b5092915050565b5f610e8782610e05565b9150610e9283610e05565b9250828203905060ff811115610eab57610eaa610a92565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f610ee882610a89565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f1a57610f19610a92565b5b600182019050919050565b5f610f2f82610e05565b9150610f3a83610e05565b9250828201905060ff811115610f5357610f52610a92565b5b9291505056fea26469706673582212202b51b32ee0b43ca6947a04ddfe7052d3e0fc1af6db180e4d77195af0d3ffbe0564736f6c634300081c0033",
}

// BankapiABI is the input ABI used to generate the binding from.
// Deprecated: Use BankapiMetaData.ABI instead.
var BankapiABI = BankapiMetaData.ABI

// BankapiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankapiMetaData.Bin instead.
var BankapiBin = BankapiMetaData.Bin

// DeployBankapi deploys a new Ethereum contract, binding an instance of Bankapi to it.
func DeployBankapi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bankapi, error) {
	parsed, err := BankapiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankapiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bankapi{BankapiCaller: BankapiCaller{contract: contract}, BankapiTransactor: BankapiTransactor{contract: contract}, BankapiFilterer: BankapiFilterer{contract: contract}}, nil
}

// Bankapi is an auto generated Go binding around an Ethereum contract.
type Bankapi struct {
	BankapiCaller     // Read-only binding to the contract
	BankapiTransactor // Write-only binding to the contract
	BankapiFilterer   // Log filterer for contract events
}

// BankapiCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankapiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankapiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankapiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankapiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankapiSession struct {
	Contract     *Bankapi          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankapiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankapiCallerSession struct {
	Contract *BankapiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BankapiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankapiTransactorSession struct {
	Contract     *BankapiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BankapiRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankapiRaw struct {
	Contract *Bankapi // Generic contract binding to access the raw methods on
}

// BankapiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankapiCallerRaw struct {
	Contract *BankapiCaller // Generic read-only contract binding to access the raw methods on
}

// BankapiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankapiTransactorRaw struct {
	Contract *BankapiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBankapi creates a new instance of Bankapi, bound to a specific deployed contract.
func NewBankapi(address common.Address, backend bind.ContractBackend) (*Bankapi, error) {
	contract, err := bindBankapi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bankapi{BankapiCaller: BankapiCaller{contract: contract}, BankapiTransactor: BankapiTransactor{contract: contract}, BankapiFilterer: BankapiFilterer{contract: contract}}, nil
}

// NewBankapiCaller creates a new read-only instance of Bankapi, bound to a specific deployed contract.
func NewBankapiCaller(address common.Address, caller bind.ContractCaller) (*BankapiCaller, error) {
	contract, err := bindBankapi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankapiCaller{contract: contract}, nil
}

// NewBankapiTransactor creates a new write-only instance of Bankapi, bound to a specific deployed contract.
func NewBankapiTransactor(address common.Address, transactor bind.ContractTransactor) (*BankapiTransactor, error) {
	contract, err := bindBankapi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankapiTransactor{contract: contract}, nil
}

// NewBankapiFilterer creates a new log filterer instance of Bankapi, bound to a specific deployed contract.
func NewBankapiFilterer(address common.Address, filterer bind.ContractFilterer) (*BankapiFilterer, error) {
	contract, err := bindBankapi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankapiFilterer{contract: contract}, nil
}

// bindBankapi binds a generic wrapper to an already deployed contract.
func bindBankapi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankapiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankapi *BankapiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankapi.Contract.BankapiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankapi *BankapiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.Contract.BankapiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankapi *BankapiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankapi.Contract.BankapiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankapi *BankapiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankapi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankapi *BankapiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankapi *BankapiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankapi.Contract.contract.Transact(opts, method, params...)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiCaller) API(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "API")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiSession) API() (common.Address, error) {
	return _Bankapi.Contract.API(&_Bankapi.CallOpts)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bankapi *BankapiCallerSession) API() (common.Address, error) {
	return _Bankapi.Contract.API(&_Bankapi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiSession) Owner() (common.Address, error) {
	return _Bankapi.Contract.Owner(&_Bankapi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bankapi *BankapiCallerSession) Owner() (common.Address, error) {
	return _Bankapi.Contract.Owner(&_Bankapi.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bankapi.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiSession) Version() (string, error) {
	return _Bankapi.Contract.Version(&_Bankapi.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bankapi *BankapiCallerSession) Version() (string, error) {
	return _Bankapi.Contract.Version(&_Bankapi.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiSession) Deposit() (*types.Transaction, error) {
	return _Bankapi.Contract.Deposit(&_Bankapi.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bankapi *BankapiTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bankapi.Contract.Deposit(&_Bankapi.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankapi.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiSession) Withdraw() (*types.Transaction, error) {
	return _Bankapi.Contract.Withdraw(&_Bankapi.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bankapi *BankapiTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bankapi.Contract.Withdraw(&_Bankapi.TransactOpts)
}

// BankapiEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bankapi contract.
type BankapiEventLogIterator struct {
	Event *BankapiEventLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BankapiEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankapiEventLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BankapiEventLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BankapiEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankapiEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankapiEventLog represents a EventLog event raised by the Bankapi contract.
type BankapiEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankapi *BankapiFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankapiEventLogIterator, error) {

	logs, sub, err := _Bankapi.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankapiEventLogIterator{contract: _Bankapi.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankapi *BankapiFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankapiEventLog) (event.Subscription, error) {

	logs, sub, err := _Bankapi.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankapiEventLog)
				if err := _Bankapi.contract.UnpackLog(event, "EventLog", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventLog is a log parse operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bankapi *BankapiFilterer) ParseEventLog(log types.Log) (*BankapiEventLog, error) {
	event := new(BankapiEventLog)
	if err := _Bankapi.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
