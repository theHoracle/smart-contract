// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"EventLog\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"API\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"SetContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506117608061005c5f395ff3fe60806040526004361061007a575f3560e01c8063bb62860d1161004d578063bb62860d14610106578063d2aadb3c14610130578063e63f341f14610158578063ed21248c146101945761007a565b80630ef678871461007e57806357ea89b6146100a85780637d7b0099146100b2578063b4a99a4e146100dc575b5f5ffd5b348015610089575f5ffd5b5061009261019e565b60405161009f9190610b69565b60405180910390f35b6100b06101e2565b005b3480156100bd575f5ffd5b506100c661034c565b6040516100d39190610bc1565b60405180910390f35b3480156100e7575f5ffd5b506100f0610370565b6040516100fd9190610bc1565b60405180910390f35b348015610111575f5ffd5b5061011a610395565b6040516101279190610c4a565b60405180910390f35b34801561013b575f5ffd5b5061015660048036038101906101519190610ca5565b610421565b005b348015610163575f5ffd5b5061017e60048036038101906101799190610ca5565b6106c8565b60405161018b9190610b69565b60405180910390f35b61019c610766565b005b5f60035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905090565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527f57ea89b6000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516102aa9190610d14565b5f60405180830381855af49150503d805f81146102e2576040519150601f19603f3d011682016040523d82523d5f602084013e6102e7565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610315826108d0565b6040516020016103259190610db0565b6040516020818303038152906040526040516103419190610c4a565b60405180910390a150565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600180546103a290610e11565b80601f01602080910402602001604051908101604052809291908181526020018280546103ce90610e11565b80156104195780601f106103f057610100808354040283529160200191610419565b820191905f5260205f20905b8154815290600101906020018083116103fc57829003601f168201915b505050505081565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610479575f5ffd5b805f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fbb62860d000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516105819190610d14565b5f604051808303815f865af19150503d805f81146105ba576040519150601f19603f3d011682016040523d82523d5f602084013e6105bf565b606091505b509150915081156105f257808060200190518101906105de9190610f5f565b600190816105ec9190611146565b50610638565b6040518060400160405280600781526020017f756e6b6e6f776e00000000000000000000000000000000000000000000000000815250600190816106369190611146565b505b7fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a6106825f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610953565b61068b846108d0565b600160405160200161069f93929190611307565b6040516020818303038152906040526040516106bb9190610c4a565b60405180910390a1505050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610721575f5ffd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166040516024016040516020818303038152906040527fed21248c000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161082e9190610d14565b5f60405180830381855af49150503d805f8114610866576040519150601f19603f3d011682016040523d82523d5f602084013e61086b565b606091505b505090507fd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a610899826108d0565b6040516020016108a99190610db0565b6040516020818303038152906040526040516108c59190610c4a565b60405180910390a150565b60608115610915576040518060400160405280600481526020017f7472756500000000000000000000000000000000000000000000000000000000815250905061094e565b6040518060400160405280600581526020017f66616c736500000000000000000000000000000000000000000000000000000081525090505b919050565b60605f602867ffffffffffffffff81111561097157610970610e49565b5b6040519080825280601f01601f1916602001820160405280156109a35781602001600182028036833780820191505090505b5090505f5f90505b6014811015610b02575f8160136109c291906113a0565b60086109ce91906113d3565b60026109da9190611543565b8573ffffffffffffffffffffffffffffffffffffffff166109fb91906115ba565b60f81b90505f60108260f81c610a1191906115f6565b60f81b90505f8160f81c6010610a279190611626565b8360f81c610a359190611662565b60f81b9050610a4382610b0c565b85856002610a5191906113d3565b81518110610a6257610a61611696565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350610a9981610b0c565b856001866002610aa991906113d3565b610ab391906116c3565b81518110610ac457610ac3611696565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535050505080806001019150506109ab565b5080915050919050565b5f600a8260f81c60ff161015610b365760308260f81c610b2c91906116f6565b60f81b9050610b4c565b60578260f81c610b4691906116f6565b60f81b90505b919050565b5f819050919050565b610b6381610b51565b82525050565b5f602082019050610b7c5f830184610b5a565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610bab82610b82565b9050919050565b610bbb81610ba1565b82525050565b5f602082019050610bd45f830184610bb2565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610c1c82610bda565b610c268185610be4565b9350610c36818560208601610bf4565b610c3f81610c02565b840191505092915050565b5f6020820190508181035f830152610c628184610c12565b905092915050565b5f604051905090565b5f5ffd5b5f5ffd5b610c8481610ba1565b8114610c8e575f5ffd5b50565b5f81359050610c9f81610c7b565b92915050565b5f60208284031215610cba57610cb9610c73565b5b5f610cc784828501610c91565b91505092915050565b5f81519050919050565b5f81905092915050565b5f610cee82610cd0565b610cf88185610cda565b9350610d08818560208601610bf4565b80840191505092915050565b5f610d1f8284610ce4565b915081905092915050565b7f737563636573735b000000000000000000000000000000000000000000000000815250565b5f81905092915050565b5f610d6482610bda565b610d6e8185610d50565b9350610d7e818560208601610bf4565b80840191505092915050565b7f5d00000000000000000000000000000000000000000000000000000000000000815250565b5f610dba82610d2a565b600882019150610dca8284610d5a565b9150610dd582610d8a565b60018201915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610e2857607f821691505b602082108103610e3b57610e3a610de4565b5b50919050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e7f82610c02565b810181811067ffffffffffffffff82111715610e9e57610e9d610e49565b5b80604052505050565b5f610eb0610c6a565b9050610ebc8282610e76565b919050565b5f67ffffffffffffffff821115610edb57610eda610e49565b5b610ee482610c02565b9050602081019050919050565b5f610f03610efe84610ec1565b610ea7565b905082815260208101848484011115610f1f57610f1e610e45565b5b610f2a848285610bf4565b509392505050565b5f82601f830112610f4657610f45610e41565b5b8151610f56848260208601610ef1565b91505092915050565b5f60208284031215610f7457610f73610c73565b5b5f82015167ffffffffffffffff811115610f9157610f90610c77565b5b610f9d84828501610f32565b91505092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026110027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610fc7565b61100c8683610fc7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61104761104261103d84610b51565b611024565b610b51565b9050919050565b5f819050919050565b6110608361102d565b61107461106c8261104e565b848454610fd3565b825550505050565b5f5f905090565b61108b61107c565b611096818484611057565b505050565b5b818110156110b9576110ae5f82611083565b60018101905061109c565b5050565b601f8211156110fe576110cf81610fa6565b6110d884610fb8565b810160208510156110e7578190505b6110fb6110f385610fb8565b83018261109b565b50505b505050565b5f82821c905092915050565b5f61111e5f1984600802611103565b1980831691505092915050565b5f611136838361110f565b9150826002028217905092915050565b61114f82610bda565b67ffffffffffffffff81111561116857611167610e49565b5b6111728254610e11565b61117d8282856110bd565b5f60209050601f8311600181146111ae575f841561119c578287015190505b6111a6858261112b565b86555061120d565b601f1984166111bc86610fa6565b5f5b828110156111e3578489015182556001820191506020850194506020810190506111be565b8683101561120057848901516111fc601f89168261110f565b8355505b6001600288020188555050505b505050505050565b7f636f6e74726163745b0000000000000000000000000000000000000000000000815250565b7f5d20737563636573735b00000000000000000000000000000000000000000000815250565b7f5d2076657273696f6e5b00000000000000000000000000000000000000000000815250565b5f815461129381610e11565b61129d8186610d50565b9450600182165f81146112b757600181146112cc576112fe565b60ff19831686528115158202860193506112fe565b6112d585610fa6565b5f5b838110156112f6578154818901526001820191506020810190506112d7565b838801955050505b50505092915050565b5f61131182611215565b6009820191506113218286610d5a565b915061132c8261123b565b600a8201915061133c8285610d5a565b915061134782611261565b600a820191506113578284611287565b915061136282610d8a565b600182019150819050949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6113aa82610b51565b91506113b583610b51565b92508282039050818111156113cd576113cc611373565b5b92915050565b5f6113dd82610b51565b91506113e883610b51565b92508282026113f681610b51565b9150828204841483151761140d5761140c611373565b5b5092915050565b5f8160011c9050919050565b5f5f8291508390505b60018511156114695780860481111561144557611444611373565b5b60018516156114545780820291505b808102905061146285611414565b9450611429565b94509492505050565b5f82611481576001905061153c565b8161148e575f905061153c565b81600181146114a457600281146114ae576114dd565b600191505061153c565b60ff8411156114c0576114bf611373565b5b8360020a9150848211156114d7576114d6611373565b5b5061153c565b5060208310610133831016604e8410600b84101617156115125782820a90508381111561150d5761150c611373565b5b61153c565b61151f8484846001611420565b9250905081840481111561153657611535611373565b5b81810290505b9392505050565b5f61154d82610b51565b915061155883610b51565b92506115857fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611472565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6115c482610b51565b91506115cf83610b51565b9250826115df576115de61158d565b5b828204905092915050565b5f60ff82169050919050565b5f611600826115ea565b915061160b836115ea565b92508261161b5761161a61158d565b5b828204905092915050565b5f611630826115ea565b915061163b836115ea565b9250828202611649816115ea565b915080821461165b5761165a611373565b5b5092915050565b5f61166c826115ea565b9150611677836115ea565b9250828203905060ff8111156116905761168f611373565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6116cd82610b51565b91506116d883610b51565b92508282019050808211156116f0576116ef611373565b5b92915050565b5f611700826115ea565b915061170b836115ea565b9250828201905060ff81111561172457611723611373565b5b9291505056fea26469706673582212205c6b56ab9336f867ee733421e44d6ceba88696cc6882f092d991931e05e579fc64736f6c634300081c0033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankCaller) API(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "API")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankSession) API() (common.Address, error) {
	return _Bank.Contract.API(&_Bank.CallOpts)
}

// API is a free data retrieval call binding the contract method 0x7d7b0099.
//
// Solidity: function API() view returns(address)
func (_Bank *BankCallerSession) API() (common.Address, error) {
	return _Bank.Contract.API(&_Bank.CallOpts)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCaller) AccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "AccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// AccountBalance is a free data retrieval call binding the contract method 0xe63f341f.
//
// Solidity: function AccountBalance(address account) view returns(uint256)
func (_Bank *BankCallerSession) AccountBalance(account common.Address) (*big.Int, error) {
	return _Bank.Contract.AccountBalance(&_Bank.CallOpts, account)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_Bank *BankCallerSession) Balance() (*big.Int, error) {
	return _Bank.Contract.Balance(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb4a99a4e.
//
// Solidity: function Owner() view returns(address)
func (_Bank *BankCallerSession) Owner() (common.Address, error) {
	return _Bank.Contract.Owner(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "Version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0xbb62860d.
//
// Solidity: function Version() view returns(string)
func (_Bank *BankCallerSession) Version() (string, error) {
	return _Bank.Contract.Version(&_Bank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Bank *BankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankTransactor) SetContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "SetContract", contractAddr)
}

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankSession) SetContract(contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SetContract(&_Bank.TransactOpts, contractAddr)
}

// SetContract is a paid mutator transaction binding the contract method 0xd2aadb3c.
//
// Solidity: function SetContract(address contractAddr) returns()
func (_Bank *BankTransactorSession) SetContract(contractAddr common.Address) (*types.Transaction, error) {
	return _Bank.Contract.SetContract(&_Bank.TransactOpts, contractAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "Withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x57ea89b6.
//
// Solidity: function Withdraw() payable returns()
func (_Bank *BankTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts)
}

// BankEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the Bank contract.
type BankEventLogIterator struct {
	Event *BankEventLog // Event containing the contract specifics and raw log

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
func (it *BankEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankEventLog)
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
		it.Event = new(BankEventLog)
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
func (it *BankEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankEventLog represents a EventLog event raised by the Bank contract.
type BankEventLog struct {
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) FilterEventLog(opts *bind.FilterOpts) (*BankEventLogIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return &BankEventLogIterator{contract: _Bank.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0xd3c51ea1865a5f43e30629abcc5e5f1f5a8a28d7cd45aface7cb4bb5c4a1a18a.
//
// Solidity: event EventLog(string value)
func (_Bank *BankFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *BankEventLog) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "EventLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankEventLog)
				if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
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
func (_Bank *BankFilterer) ParseEventLog(log types.Log) (*BankEventLog, error) {
	event := new(BankEventLog)
	if err := _Bank.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
