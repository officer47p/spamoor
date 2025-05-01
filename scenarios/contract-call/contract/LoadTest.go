// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DataReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"name\":\"PrecompileResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operations\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"}],\"name\":\"SloadPerformed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"}],\"name\":\"executeModExpAttack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"processCallData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"testArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"}],\"name\":\"testArraySload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"}],\"name\":\"testCombinedSload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"testMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"}],\"name\":\"testMappingSload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"iterations\",\"type\":\"uint256\"}],\"name\":\"testSimpleSload\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50602a60008190555060005b60648110156200008e57600281620000369190620000ce565b60016000838152602001908152602001600020819055506003816200005c9190620000ce565b6002826064811062000073576200007262000119565b5b01819055508080620000859062000148565b9150506200001d565b5062000195565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620000db8262000095565b9150620000e88362000095565b9250828202620000f88162000095565b915082820484148315176200011257620001116200009f565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000620001558262000095565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036200018a57620001896200009f565b5b600182019050919050565b610ec080620001a56000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80635ea958c5116100715780635ea958c51461018c5780638af5de72146101a85780639b3874b1146101c6578063a1d9ea81146101f6578063a82252d714610226578063dbdf7fce14610242576100a9565b806323d3c142146100ae578063280a6a19146100cc57806332153f97146100fc578063463f5c301461012c5780635ea946ce1461015c575b600080fd5b6100b661024c565b6040516100c3919061096e565b60405180910390f35b6100e660048036038101906100e191906109c9565b610252565b6040516100f3919061096e565b60405180910390f35b610116600480360381019061011191906109c9565b6102ee565b604051610123919061096e565b60405180910390f35b610146600480360381019061014191906109c9565b6103a9565b604051610153919061096e565b60405180910390f35b610176600480360381019061017191906109c9565b6103c4565b604051610183919061096e565b60405180910390f35b6101a660048036038101906101a191906109c9565b61047d565b005b6101b06107e3565b6040516101bd919061096e565b60405180910390f35b6101e060048036038101906101db91906109c9565b6107e9565b6040516101ed919061096e565b60405180910390f35b610210600480360381019061020b91906109c9565b610801565b60405161021d919061096e565b60405180910390f35b610240600480360381019061023b9190610b3c565b610911565b005b61024a61094b565b005b60665481565b6000805a9050600080600090505b84811015610280576000549150808061027890610bb4565b915050610260565b5060005a8361028f9190610bfc565b905084606660008282546102a39190610c30565b925050819055507fc3c21cc9c847e02fad29d076e5ab095fbc80398901690cb580f2d97b412321dd85826040516102db929190610c64565b60405180910390a1819350505050919050565b6000805a9050600080600090505b8481101561033b5760026064826103139190610cbc565b6064811061032457610323610ced565b5b01549150808061033390610bb4565b9150506102fc565b5060005a8361034a9190610bfc565b9050846066600082825461035e9190610c30565b925050819055507fc3c21cc9c847e02fad29d076e5ab095fbc80398901690cb580f2d97b412321dd8582604051610396929190610c64565b60405180910390a1819350505050919050565b600281606481106103b957600080fd5b016000915090505481565b6000805a9050600080600090505b8481101561040f57600160006064836103eb9190610cbc565b8152602001908152602001600020549150808061040790610bb4565b9150506103d2565b5060005a8361041e9190610bfc565b905084606660008282546104329190610c30565b925050819055507fc3c21cc9c847e02fad29d076e5ab095fbc80398901690cb580f2d97b412321dd858260405161046a929190610c64565b60405180910390a1819350505050919050565b6000602067ffffffffffffffff81111561049a57610499610a11565b5b6040519080825280601f01601f1916602001820160405280156104cc5781602001600182028036833780820191505090505b5090506000604067ffffffffffffffff8111156104ec576104eb610a11565b5b6040519080825280601f01601f19166020018201604052801561051e5781602001600182028036833780820191505090505b5090506000602067ffffffffffffffff81111561053e5761053d610a11565b5b6040519080825280601f01601f1916602001820160405280156105705781602001600182028036833780820191505090505b50905060005b83518110156105d95760ff60f81b84828151811061059757610596610ced565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806105d190610bb4565b915050610576565b5060005b82518110156106405760ff60f81b8382815181106105fe576105fd610ced565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808061063890610bb4565b9150506105dd565b5060005b81518110156106a75760ff60f81b82828151811061066557610664610ced565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808061069f90610bb4565b915050610644565b50600081518351855160206106bc9190610c30565b6106c69190610c30565b6106d09190610c30565b905060008167ffffffffffffffff8111156106ee576106ed610a11565b5b6040519080825280601f01601f1916602001820160405280156107205781602001600182028036833780820191505090505b5090506020808201526040808201526020606082015260208501516080820152602084015160a0820152604084015160c0820152602083015160e082015260008060005b888110156107d8575a9250602060008660208701600060056207a120f191507f975cd3e5daaf32e0f68a95002705ced1ea299b11971472398e116310e4eef0aa5a846107b09190610bfc565b6040516107bd9190610d79565b60405180910390a180806107d090610bb4565b915050610764565b505050505050505050565b60005481565b60016020528060005260406000206000915090505481565b6000805a9050600080600090505b8481101561088b5760005491506001600060648361082d9190610cbc565b815260200190815260200160002054826108479190610c30565b915060026064826108589190610cbc565b6064811061086957610868610ced565b5b0154826108769190610c30565b9150808061088390610bb4565b91505061080f565b5060005a8361089a9190610bfc565b90506003856108a99190610da7565b606660008282546108ba9190610c30565b925050819055507fc3c21cc9c847e02fad29d076e5ab095fbc80398901690cb580f2d97b412321dd6003866108ef9190610da7565b826040516108fe929190610c64565b60405180910390a1819350505050919050565b7fc0629c5930c116f6e8c1eed6df81fc90be33cc9a751b2188394105dea097b024816040516109409190610e68565b60405180910390a150565b6000606681905550565b6000819050919050565b61096881610955565b82525050565b6000602082019050610983600083018461095f565b92915050565b6000604051905090565b600080fd5b600080fd5b6109a681610955565b81146109b157600080fd5b50565b6000813590506109c38161099d565b92915050565b6000602082840312156109df576109de610993565b5b60006109ed848285016109b4565b91505092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610a4982610a00565b810181811067ffffffffffffffff82111715610a6857610a67610a11565b5b80604052505050565b6000610a7b610989565b9050610a878282610a40565b919050565b600067ffffffffffffffff821115610aa757610aa6610a11565b5b610ab082610a00565b9050602081019050919050565b82818337600083830152505050565b6000610adf610ada84610a8c565b610a71565b905082815260208101848484011115610afb57610afa6109fb565b5b610b06848285610abd565b509392505050565b600082601f830112610b2357610b226109f6565b5b8135610b33848260208601610acc565b91505092915050565b600060208284031215610b5257610b51610993565b5b600082013567ffffffffffffffff811115610b7057610b6f610998565b5b610b7c84828501610b0e565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610bbf82610955565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610bf157610bf0610b85565b5b600182019050919050565b6000610c0782610955565b9150610c1283610955565b9250828203905081811115610c2a57610c29610b85565b5b92915050565b6000610c3b82610955565b9150610c4683610955565b9250828201905080821115610c5e57610c5d610b85565b5b92915050565b6000604082019050610c79600083018561095f565b610c86602083018461095f565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610cc782610955565b9150610cd283610955565b925082610ce257610ce1610c8d565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082825260208201905092915050565b7f4d4f444558500000000000000000000000000000000000000000000000000000600082015250565b6000610d63600683610d1c565b9150610d6e82610d2d565b602082019050919050565b60006040820190508181036000830152610d9281610d56565b9050610da1602083018461095f565b92915050565b6000610db282610955565b9150610dbd83610955565b9250828202610dcb81610955565b91508282048414831517610de257610de1610b85565b5b5092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610e23578082015181840152602081019050610e08565b60008484015250505050565b6000610e3a82610de9565b610e448185610df4565b9350610e54818560208601610e05565b610e5d81610a00565b840191505092915050565b60006020820190508181036000830152610e828184610e2f565b90509291505056fea264697066735822122073b1a1249cf6d8b24f776b24386bd89554bd1559be9c71202661bd09ada86d0064736f6c63430008130033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// OperationCount is a free data retrieval call binding the contract method 0x23d3c142.
//
// Solidity: function operationCount() view returns(uint256)
func (_Contract *ContractCaller) OperationCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "operationCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperationCount is a free data retrieval call binding the contract method 0x23d3c142.
//
// Solidity: function operationCount() view returns(uint256)
func (_Contract *ContractSession) OperationCount() (*big.Int, error) {
	return _Contract.Contract.OperationCount(&_Contract.CallOpts)
}

// OperationCount is a free data retrieval call binding the contract method 0x23d3c142.
//
// Solidity: function operationCount() view returns(uint256)
func (_Contract *ContractCallerSession) OperationCount() (*big.Int, error) {
	return _Contract.Contract.OperationCount(&_Contract.CallOpts)
}

// TestArray is a free data retrieval call binding the contract method 0x463f5c30.
//
// Solidity: function testArray(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) TestArray(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "testArray", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestArray is a free data retrieval call binding the contract method 0x463f5c30.
//
// Solidity: function testArray(uint256 ) view returns(uint256)
func (_Contract *ContractSession) TestArray(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.TestArray(&_Contract.CallOpts, arg0)
}

// TestArray is a free data retrieval call binding the contract method 0x463f5c30.
//
// Solidity: function testArray(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) TestArray(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.TestArray(&_Contract.CallOpts, arg0)
}

// TestMapping is a free data retrieval call binding the contract method 0x9b3874b1.
//
// Solidity: function testMapping(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) TestMapping(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "testMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestMapping is a free data retrieval call binding the contract method 0x9b3874b1.
//
// Solidity: function testMapping(uint256 ) view returns(uint256)
func (_Contract *ContractSession) TestMapping(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.TestMapping(&_Contract.CallOpts, arg0)
}

// TestMapping is a free data retrieval call binding the contract method 0x9b3874b1.
//
// Solidity: function testMapping(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) TestMapping(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.TestMapping(&_Contract.CallOpts, arg0)
}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_Contract *ContractCaller) TestValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "testValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_Contract *ContractSession) TestValue() (*big.Int, error) {
	return _Contract.Contract.TestValue(&_Contract.CallOpts)
}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_Contract *ContractCallerSession) TestValue() (*big.Int, error) {
	return _Contract.Contract.TestValue(&_Contract.CallOpts)
}

// ExecuteModExpAttack is a paid mutator transaction binding the contract method 0x5ea958c5.
//
// Solidity: function executeModExpAttack(uint256 iterations) returns()
func (_Contract *ContractTransactor) ExecuteModExpAttack(opts *bind.TransactOpts, iterations *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "executeModExpAttack", iterations)
}

// ExecuteModExpAttack is a paid mutator transaction binding the contract method 0x5ea958c5.
//
// Solidity: function executeModExpAttack(uint256 iterations) returns()
func (_Contract *ContractSession) ExecuteModExpAttack(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteModExpAttack(&_Contract.TransactOpts, iterations)
}

// ExecuteModExpAttack is a paid mutator transaction binding the contract method 0x5ea958c5.
//
// Solidity: function executeModExpAttack(uint256 iterations) returns()
func (_Contract *ContractTransactorSession) ExecuteModExpAttack(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ExecuteModExpAttack(&_Contract.TransactOpts, iterations)
}

// ProcessCallData is a paid mutator transaction binding the contract method 0xa82252d7.
//
// Solidity: function processCallData(bytes data) returns()
func (_Contract *ContractTransactor) ProcessCallData(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "processCallData", data)
}

// ProcessCallData is a paid mutator transaction binding the contract method 0xa82252d7.
//
// Solidity: function processCallData(bytes data) returns()
func (_Contract *ContractSession) ProcessCallData(data []byte) (*types.Transaction, error) {
	return _Contract.Contract.ProcessCallData(&_Contract.TransactOpts, data)
}

// ProcessCallData is a paid mutator transaction binding the contract method 0xa82252d7.
//
// Solidity: function processCallData(bytes data) returns()
func (_Contract *ContractTransactorSession) ProcessCallData(data []byte) (*types.Transaction, error) {
	return _Contract.Contract.ProcessCallData(&_Contract.TransactOpts, data)
}

// ResetCounter is a paid mutator transaction binding the contract method 0xdbdf7fce.
//
// Solidity: function resetCounter() returns()
func (_Contract *ContractTransactor) ResetCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "resetCounter")
}

// ResetCounter is a paid mutator transaction binding the contract method 0xdbdf7fce.
//
// Solidity: function resetCounter() returns()
func (_Contract *ContractSession) ResetCounter() (*types.Transaction, error) {
	return _Contract.Contract.ResetCounter(&_Contract.TransactOpts)
}

// ResetCounter is a paid mutator transaction binding the contract method 0xdbdf7fce.
//
// Solidity: function resetCounter() returns()
func (_Contract *ContractTransactorSession) ResetCounter() (*types.Transaction, error) {
	return _Contract.Contract.ResetCounter(&_Contract.TransactOpts)
}

// TestArraySload is a paid mutator transaction binding the contract method 0x32153f97.
//
// Solidity: function testArraySload(uint256 iterations) returns(uint256)
func (_Contract *ContractTransactor) TestArraySload(opts *bind.TransactOpts, iterations *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "testArraySload", iterations)
}

// TestArraySload is a paid mutator transaction binding the contract method 0x32153f97.
//
// Solidity: function testArraySload(uint256 iterations) returns(uint256)
func (_Contract *ContractSession) TestArraySload(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TestArraySload(&_Contract.TransactOpts, iterations)
}

// TestArraySload is a paid mutator transaction binding the contract method 0x32153f97.
//
// Solidity: function testArraySload(uint256 iterations) returns(uint256)
func (_Contract *ContractTransactorSession) TestArraySload(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TestArraySload(&_Contract.TransactOpts, iterations)
}

// TestCombinedSload is a paid mutator transaction binding the contract method 0xa1d9ea81.
//
// Solidity: function testCombinedSload(uint256 iterations) returns(uint256)
func (_Contract *ContractTransactor) TestCombinedSload(opts *bind.TransactOpts, iterations *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "testCombinedSload", iterations)
}

// TestCombinedSload is a paid mutator transaction binding the contract method 0xa1d9ea81.
//
// Solidity: function testCombinedSload(uint256 iterations) returns(uint256)
func (_Contract *ContractSession) TestCombinedSload(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TestCombinedSload(&_Contract.TransactOpts, iterations)
}

// TestCombinedSload is a paid mutator transaction binding the contract method 0xa1d9ea81.
//
// Solidity: function testCombinedSload(uint256 iterations) returns(uint256)
func (_Contract *ContractTransactorSession) TestCombinedSload(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TestCombinedSload(&_Contract.TransactOpts, iterations)
}

// TestMappingSload is a paid mutator transaction binding the contract method 0x5ea946ce.
//
// Solidity: function testMappingSload(uint256 iterations) returns(uint256)
func (_Contract *ContractTransactor) TestMappingSload(opts *bind.TransactOpts, iterations *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "testMappingSload", iterations)
}

// TestMappingSload is a paid mutator transaction binding the contract method 0x5ea946ce.
//
// Solidity: function testMappingSload(uint256 iterations) returns(uint256)
func (_Contract *ContractSession) TestMappingSload(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TestMappingSload(&_Contract.TransactOpts, iterations)
}

// TestMappingSload is a paid mutator transaction binding the contract method 0x5ea946ce.
//
// Solidity: function testMappingSload(uint256 iterations) returns(uint256)
func (_Contract *ContractTransactorSession) TestMappingSload(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TestMappingSload(&_Contract.TransactOpts, iterations)
}

// TestSimpleSload is a paid mutator transaction binding the contract method 0x280a6a19.
//
// Solidity: function testSimpleSload(uint256 iterations) returns(uint256)
func (_Contract *ContractTransactor) TestSimpleSload(opts *bind.TransactOpts, iterations *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "testSimpleSload", iterations)
}

// TestSimpleSload is a paid mutator transaction binding the contract method 0x280a6a19.
//
// Solidity: function testSimpleSload(uint256 iterations) returns(uint256)
func (_Contract *ContractSession) TestSimpleSload(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TestSimpleSload(&_Contract.TransactOpts, iterations)
}

// TestSimpleSload is a paid mutator transaction binding the contract method 0x280a6a19.
//
// Solidity: function testSimpleSload(uint256 iterations) returns(uint256)
func (_Contract *ContractTransactorSession) TestSimpleSload(iterations *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TestSimpleSload(&_Contract.TransactOpts, iterations)
}

// ContractDataReceivedIterator is returned from FilterDataReceived and is used to iterate over the raw logs and unpacked data for DataReceived events raised by the Contract contract.
type ContractDataReceivedIterator struct {
	Event *ContractDataReceived // Event containing the contract specifics and raw log

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
func (it *ContractDataReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDataReceived)
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
		it.Event = new(ContractDataReceived)
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
func (it *ContractDataReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDataReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDataReceived represents a DataReceived event raised by the Contract contract.
type ContractDataReceived struct {
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDataReceived is a free log retrieval operation binding the contract event 0xc0629c5930c116f6e8c1eed6df81fc90be33cc9a751b2188394105dea097b024.
//
// Solidity: event DataReceived(bytes data)
func (_Contract *ContractFilterer) FilterDataReceived(opts *bind.FilterOpts) (*ContractDataReceivedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DataReceived")
	if err != nil {
		return nil, err
	}
	return &ContractDataReceivedIterator{contract: _Contract.contract, event: "DataReceived", logs: logs, sub: sub}, nil
}

// WatchDataReceived is a free log subscription operation binding the contract event 0xc0629c5930c116f6e8c1eed6df81fc90be33cc9a751b2188394105dea097b024.
//
// Solidity: event DataReceived(bytes data)
func (_Contract *ContractFilterer) WatchDataReceived(opts *bind.WatchOpts, sink chan<- *ContractDataReceived) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DataReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDataReceived)
				if err := _Contract.contract.UnpackLog(event, "DataReceived", log); err != nil {
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

// ParseDataReceived is a log parse operation binding the contract event 0xc0629c5930c116f6e8c1eed6df81fc90be33cc9a751b2188394105dea097b024.
//
// Solidity: event DataReceived(bytes data)
func (_Contract *ContractFilterer) ParseDataReceived(log types.Log) (*ContractDataReceived, error) {
	event := new(ContractDataReceived)
	if err := _Contract.contract.UnpackLog(event, "DataReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPrecompileResultIterator is returned from FilterPrecompileResult and is used to iterate over the raw logs and unpacked data for PrecompileResult events raised by the Contract contract.
type ContractPrecompileResultIterator struct {
	Event *ContractPrecompileResult // Event containing the contract specifics and raw log

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
func (it *ContractPrecompileResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPrecompileResult)
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
		it.Event = new(ContractPrecompileResult)
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
func (it *ContractPrecompileResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPrecompileResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPrecompileResult represents a PrecompileResult event raised by the Contract contract.
type ContractPrecompileResult struct {
	Name    string
	GasUsed *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPrecompileResult is a free log retrieval operation binding the contract event 0x975cd3e5daaf32e0f68a95002705ced1ea299b11971472398e116310e4eef0aa.
//
// Solidity: event PrecompileResult(string name, uint256 gasUsed)
func (_Contract *ContractFilterer) FilterPrecompileResult(opts *bind.FilterOpts) (*ContractPrecompileResultIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PrecompileResult")
	if err != nil {
		return nil, err
	}
	return &ContractPrecompileResultIterator{contract: _Contract.contract, event: "PrecompileResult", logs: logs, sub: sub}, nil
}

// WatchPrecompileResult is a free log subscription operation binding the contract event 0x975cd3e5daaf32e0f68a95002705ced1ea299b11971472398e116310e4eef0aa.
//
// Solidity: event PrecompileResult(string name, uint256 gasUsed)
func (_Contract *ContractFilterer) WatchPrecompileResult(opts *bind.WatchOpts, sink chan<- *ContractPrecompileResult) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PrecompileResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPrecompileResult)
				if err := _Contract.contract.UnpackLog(event, "PrecompileResult", log); err != nil {
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

// ParsePrecompileResult is a log parse operation binding the contract event 0x975cd3e5daaf32e0f68a95002705ced1ea299b11971472398e116310e4eef0aa.
//
// Solidity: event PrecompileResult(string name, uint256 gasUsed)
func (_Contract *ContractFilterer) ParsePrecompileResult(log types.Log) (*ContractPrecompileResult, error) {
	event := new(ContractPrecompileResult)
	if err := _Contract.contract.UnpackLog(event, "PrecompileResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSloadPerformedIterator is returned from FilterSloadPerformed and is used to iterate over the raw logs and unpacked data for SloadPerformed events raised by the Contract contract.
type ContractSloadPerformedIterator struct {
	Event *ContractSloadPerformed // Event containing the contract specifics and raw log

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
func (it *ContractSloadPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSloadPerformed)
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
		it.Event = new(ContractSloadPerformed)
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
func (it *ContractSloadPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSloadPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSloadPerformed represents a SloadPerformed event raised by the Contract contract.
type ContractSloadPerformed struct {
	Operations *big.Int
	GasUsed    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSloadPerformed is a free log retrieval operation binding the contract event 0xc3c21cc9c847e02fad29d076e5ab095fbc80398901690cb580f2d97b412321dd.
//
// Solidity: event SloadPerformed(uint256 operations, uint256 gasUsed)
func (_Contract *ContractFilterer) FilterSloadPerformed(opts *bind.FilterOpts) (*ContractSloadPerformedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SloadPerformed")
	if err != nil {
		return nil, err
	}
	return &ContractSloadPerformedIterator{contract: _Contract.contract, event: "SloadPerformed", logs: logs, sub: sub}, nil
}

// WatchSloadPerformed is a free log subscription operation binding the contract event 0xc3c21cc9c847e02fad29d076e5ab095fbc80398901690cb580f2d97b412321dd.
//
// Solidity: event SloadPerformed(uint256 operations, uint256 gasUsed)
func (_Contract *ContractFilterer) WatchSloadPerformed(opts *bind.WatchOpts, sink chan<- *ContractSloadPerformed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SloadPerformed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSloadPerformed)
				if err := _Contract.contract.UnpackLog(event, "SloadPerformed", log); err != nil {
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

// ParseSloadPerformed is a log parse operation binding the contract event 0xc3c21cc9c847e02fad29d076e5ab095fbc80398901690cb580f2d97b412321dd.
//
// Solidity: event SloadPerformed(uint256 operations, uint256 gasUsed)
func (_Contract *ContractFilterer) ParseSloadPerformed(log types.Log) (*ContractSloadPerformed, error) {
	event := new(ContractSloadPerformed)
	if err := _Contract.contract.UnpackLog(event, "SloadPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
