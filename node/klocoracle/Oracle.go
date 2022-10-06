// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package klocoracle

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OracleABI is the input ABI used to generate the binding from.
const OracleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_adapterId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"adapterId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"NewOracleRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adapterId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"fulfillOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"string\",\"name\":\"adapterID\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"}],\"name\":\"newOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"adapterId\",\"type\":\"string\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OracleBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OracleBinRuntime = ``

// OracleBin is the compiled bytecode used for deploying new contracts.
var OracleBin = "0x60806040523480156200001157600080fd5b50604051620017c4380380620017c4833981810160405281019062000037919062000289565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600190816200008891906200053a565b50505062000621565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000d282620000a5565b9050919050565b620000e481620000c5565b8114620000f057600080fd5b50565b6000815190506200010481620000d9565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200015f8262000114565b810181811067ffffffffffffffff8211171562000181576200018062000125565b5b80604052505050565b60006200019662000091565b9050620001a4828262000154565b919050565b600067ffffffffffffffff821115620001c757620001c662000125565b5b620001d28262000114565b9050602081019050919050565b60005b83811015620001ff578082015181840152602081019050620001e2565b60008484015250505050565b6000620002226200021c84620001a9565b6200018a565b9050828152602081018484840111156200024157620002406200010f565b5b6200024e848285620001df565b509392505050565b600082601f8301126200026e576200026d6200010a565b5b8151620002808482602086016200020b565b91505092915050565b60008060408385031215620002a357620002a26200009b565b5b6000620002b385828601620000f3565b925050602083015167ffffffffffffffff811115620002d757620002d6620000a0565b5b620002e58582860162000256565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200034257607f821691505b602082108103620003585762000357620002fa565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003c27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000383565b620003ce868362000383565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200041b620004156200040f84620003e6565b620003f0565b620003e6565b9050919050565b6000819050919050565b6200043783620003fa565b6200044f620004468262000422565b84845462000390565b825550505050565b600090565b6200046662000457565b620004738184846200042c565b505050565b5b818110156200049b576200048f6000826200045c565b60018101905062000479565b5050565b601f821115620004ea57620004b4816200035e565b620004bf8462000373565b81016020851015620004cf578190505b620004e7620004de8562000373565b83018262000478565b50505b505050565b600082821c905092915050565b60006200050f60001984600802620004ef565b1980831691505092915050565b60006200052a8383620004fc565b9150826002028217905092915050565b6200054582620002ef565b67ffffffffffffffff81111562000561576200056062000125565b5b6200056d825462000329565b6200057a8282856200049f565b600060209050601f831160018114620005b257600084156200059d578287015190505b620005a985826200051c565b86555062000619565b601f198416620005c2866200035e565b60005b82811015620005ec57848901518255600182019150602085019450602081019050620005c5565b868310156200060c578489015162000608601f891682620004fc565b8355505b6001600288020188555050505b505050505050565b61119380620006316000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632aa8481f1461005c5780635cd3275c1461007a5780635fb86b01146100aa57806381d12c58146100c8578063eb3377d5146100fd575b600080fd5b61006461012d565b6040516100719190610845565b60405180910390f35b610094600480360381019061008f9190610a3e565b610151565b6040516100a19190610ac8565b60405180910390f35b6100b261034b565b6040516100bf9190610b62565b60405180910390f35b6100e260048036038101906100dd9190610bba565b6103d9565b6040516100f496959493929190610c1e565b60405180910390f35b61011760048036038101906101129190610cb2565b6104fa565b6040516101249190610ac8565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600280549050905060006040518060c0016040528083815260200160008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001868152602001877bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020016000801916815250905060028190806001815401808255809150506001900390600052602060002090600502016000909190919091506000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201908161028e9190610efe565b5060608201518160030160006101000a81548163ffffffff021916908360e01c021790555060808201518160030160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600401555050817f394611a31616be6c6ac9b70c798c776a5d7533f25ad93f61d72db1968d3ec206866040516103369190610b62565b60405180910390a26001925050509392505050565b6001805461035890610d21565b80601f016020809104026020016040519081016040528092919081815260200182805461038490610d21565b80156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b505050505081565b600281815481106103e957600080fd5b90600052602060002090600502016000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600201805461043890610d21565b80601f016020809104026020016040519081016040528092919081815260200182805461046490610d21565b80156104b15780601f10610486576101008083540402835291602001916104b1565b820191906000526020600020905b81548152906001019060200180831161049457829003601f168201915b5050505050908060030160009054906101000a900460e01b908060030160049054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905086565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461058b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058290611042565b60405180910390fd5b6000600284815481106105a1576105a0611062565b5b90600052602060002090600502016040518060c0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461062a90610d21565b80601f016020809104026020016040519081016040528092919081815260200182805461065690610d21565b80156106a35780601f10610678576101008083540402835291602001916106a3565b820191906000526020600020905b81548152906001019060200180831161068657829003601f168201915b505050505081526020016003820160009054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020016003820160049054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152505090506000816080015173ffffffffffffffffffffffffffffffffffffffff168260600151856040516020016107979291906110d3565b6040516020818303038152906040526040516107b39190611146565b6000604051808303816000865af19150503d80600081146107f0576040519150601f19603f3d011682016040523d82523d6000602084013e6107f5565b606091505b50509050809250505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061082f82610804565b9050919050565b61083f81610824565b82525050565b600060208201905061085a6000830184610836565b92915050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6108a981610874565b81146108b457600080fd5b50565b6000813590506108c6816108a0565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61091f826108d6565b810181811067ffffffffffffffff8211171561093e5761093d6108e7565b5b80604052505050565b6000610951610860565b905061095d8282610916565b919050565b600067ffffffffffffffff82111561097d5761097c6108e7565b5b610986826108d6565b9050602081019050919050565b82818337600083830152505050565b60006109b56109b084610962565b610947565b9050828152602081018484840111156109d1576109d06108d1565b5b6109dc848285610993565b509392505050565b600082601f8301126109f9576109f86108cc565b5b8135610a098482602086016109a2565b91505092915050565b610a1b81610824565b8114610a2657600080fd5b50565b600081359050610a3881610a12565b92915050565b600080600060608486031215610a5757610a5661086a565b5b6000610a65868287016108b7565b935050602084013567ffffffffffffffff811115610a8657610a8561086f565b5b610a92868287016109e4565b9250506040610aa386828701610a29565b9150509250925092565b60008115159050919050565b610ac281610aad565b82525050565b6000602082019050610add6000830184610ab9565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b1d578082015181840152602081019050610b02565b60008484015250505050565b6000610b3482610ae3565b610b3e8185610aee565b9350610b4e818560208601610aff565b610b57816108d6565b840191505092915050565b60006020820190508181036000830152610b7c8184610b29565b905092915050565b6000819050919050565b610b9781610b84565b8114610ba257600080fd5b50565b600081359050610bb481610b8e565b92915050565b600060208284031215610bd057610bcf61086a565b5b6000610bde84828501610ba5565b91505092915050565b610bf081610b84565b82525050565b610bff81610874565b82525050565b6000819050919050565b610c1881610c05565b82525050565b600060c082019050610c336000830189610be7565b610c406020830188610836565b8181036040830152610c528187610b29565b9050610c616060830186610bf6565b610c6e6080830185610836565b610c7b60a0830184610c0f565b979650505050505050565b610c8f81610c05565b8114610c9a57600080fd5b50565b600081359050610cac81610c86565b92915050565b60008060408385031215610cc957610cc861086a565b5b6000610cd785828601610ba5565b9250506020610ce885828601610c9d565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610d3957607f821691505b602082108103610d4c57610d4b610cf2565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610db47fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d77565b610dbe8683610d77565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610dfb610df6610df184610b84565b610dd6565b610b84565b9050919050565b6000819050919050565b610e1583610de0565b610e29610e2182610e02565b848454610d84565b825550505050565b600090565b610e3e610e31565b610e49818484610e0c565b505050565b5b81811015610e6d57610e62600082610e36565b600181019050610e4f565b5050565b601f821115610eb257610e8381610d52565b610e8c84610d67565b81016020851015610e9b578190505b610eaf610ea785610d67565b830182610e4e565b50505b505050565b600082821c905092915050565b6000610ed560001984600802610eb7565b1980831691505092915050565b6000610eee8383610ec4565b9150826002028217905092915050565b610f0782610ae3565b67ffffffffffffffff811115610f2057610f1f6108e7565b5b610f2a8254610d21565b610f35828285610e71565b600060209050601f831160018114610f685760008415610f56578287015190505b610f608582610ee2565b865550610fc8565b601f198416610f7686610d52565b60005b82811015610f9e57848901518255600182019150602085019450602081019050610f79565b86831015610fbb5784890151610fb7601f891682610ec4565b8355505b6001600288020188555050505b505050505050565b7f4f7261636c653a205065726d697373696f6e206e656564656420666f72206e6f60008201527f6465000000000000000000000000000000000000000000000000000000000000602082015250565b600061102c602283610aee565b915061103782610fd0565b604082019050919050565b6000602082019050818103600083015261105b8161101f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000819050919050565b6110ac6110a782610874565b611091565b82525050565b6000819050919050565b6110cd6110c882610c05565b6110b2565b82525050565b60006110df828561109b565b6004820191506110ef82846110bc565b6020820191508190509392505050565b600081519050919050565b600081905092915050565b6000611120826110ff565b61112a818561110a565b935061113a818560208601610aff565b80840191505092915050565b60006111528284611115565b91508190509291505056fea2646970667358221220065105dc6a13c31d154ea5f88d6924b70659d94477c4f6a561e1e4a0393e501864736f6c63430008100033"

// DeployOracle deploys a new Klaytn contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeAddress common.Address, _adapterId string) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleBin), backend, _nodeAddress, _adapterId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around a Klaytn contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around a Klaytn contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around a Klaytn contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around a Klaytn contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(string)
func (_Oracle *OracleCaller) AdapterId(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "adapterId")
	return *ret0, err
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(string)
func (_Oracle *OracleSession) AdapterId() (string, error) {
	return _Oracle.Contract.AdapterId(&_Oracle.CallOpts)
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(string)
func (_Oracle *OracleCallerSession) AdapterId() (string, error) {
	return _Oracle.Contract.AdapterId(&_Oracle.CallOpts)
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_Oracle *OracleCaller) NodeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Oracle.contract.Call(opts, out, "nodeAddress")
	return *ret0, err
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_Oracle *OracleSession) NodeAddress() (common.Address, error) {
	return _Oracle.Contract.NodeAddress(&_Oracle.CallOpts)
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_Oracle *OracleCallerSession) NodeAddress() (common.Address, error) {
	return _Oracle.Contract.NodeAddress(&_Oracle.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestId, address nodeAddress, string adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data)
func (_Oracle *OracleCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RequestId          *big.Int
	NodeAddress        common.Address
	AdapterId          string
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
}, error) {
	ret := new(struct {
		RequestId          *big.Int
		NodeAddress        common.Address
		AdapterId          string
		CallbackFunctionId [4]byte
		CallBackContract   common.Address
		Data               [32]byte
	})
	out := ret
	err := _Oracle.contract.Call(opts, out, "requests", arg0)
	return *ret, err
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestId, address nodeAddress, string adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data)
func (_Oracle *OracleSession) Requests(arg0 *big.Int) (struct {
	RequestId          *big.Int
	NodeAddress        common.Address
	AdapterId          string
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
}, error) {
	return _Oracle.Contract.Requests(&_Oracle.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(uint256 requestId, address nodeAddress, string adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data)
func (_Oracle *OracleCallerSession) Requests(arg0 *big.Int) (struct {
	RequestId          *big.Int
	NodeAddress        common.Address
	AdapterId          string
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
}, error) {
	return _Oracle.Contract.Requests(&_Oracle.CallOpts, arg0)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0xeb3377d5.
//
// Solidity: function fulfillOracleRequest(uint256 requestId, bytes32 data) returns(bool)
func (_Oracle *OracleTransactor) FulfillOracleRequest(opts *bind.TransactOpts, requestId *big.Int, data [32]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "fulfillOracleRequest", requestId, data)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0xeb3377d5.
//
// Solidity: function fulfillOracleRequest(uint256 requestId, bytes32 data) returns(bool)
func (_Oracle *OracleSession) FulfillOracleRequest(requestId *big.Int, data [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.FulfillOracleRequest(&_Oracle.TransactOpts, requestId, data)
}

// FulfillOracleRequest is a paid mutator transaction binding the contract method 0xeb3377d5.
//
// Solidity: function fulfillOracleRequest(uint256 requestId, bytes32 data) returns(bool)
func (_Oracle *OracleTransactorSession) FulfillOracleRequest(requestId *big.Int, data [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.FulfillOracleRequest(&_Oracle.TransactOpts, requestId, data)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0x5cd3275c.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, string adapterID, address callBackContract) returns(bool)
func (_Oracle *OracleTransactor) NewOracleRequest(opts *bind.TransactOpts, callbackFunctionId [4]byte, adapterID string, callBackContract common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "newOracleRequest", callbackFunctionId, adapterID, callBackContract)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0x5cd3275c.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, string adapterID, address callBackContract) returns(bool)
func (_Oracle *OracleSession) NewOracleRequest(callbackFunctionId [4]byte, adapterID string, callBackContract common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.NewOracleRequest(&_Oracle.TransactOpts, callbackFunctionId, adapterID, callBackContract)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0x5cd3275c.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, string adapterID, address callBackContract) returns(bool)
func (_Oracle *OracleTransactorSession) NewOracleRequest(callbackFunctionId [4]byte, adapterID string, callBackContract common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.NewOracleRequest(&_Oracle.TransactOpts, callbackFunctionId, adapterID, callBackContract)
}

// OracleNewOracleRequestIterator is returned from FilterNewOracleRequest and is used to iterate over the raw logs and unpacked data for NewOracleRequest events raised by the Oracle contract.
type OracleNewOracleRequestIterator struct {
	Event *OracleNewOracleRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleNewOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleNewOracleRequest)
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
		it.Event = new(OracleNewOracleRequest)
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
func (it *OracleNewOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleNewOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleNewOracleRequest represents a NewOracleRequest event raised by the Oracle contract.
type OracleNewOracleRequest struct {
	AdapterId string
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewOracleRequest is a free log retrieval operation binding the contract event 0x394611a31616be6c6ac9b70c798c776a5d7533f25ad93f61d72db1968d3ec206.
//
// Solidity: event NewOracleRequest(string adapterId, uint256 indexed requestId)
func (_Oracle *OracleFilterer) FilterNewOracleRequest(opts *bind.FilterOpts, requestId []*big.Int) (*OracleNewOracleRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "NewOracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleNewOracleRequestIterator{contract: _Oracle.contract, event: "NewOracleRequest", logs: logs, sub: sub}, nil
}

// WatchNewOracleRequest is a free log subscription operation binding the contract event 0x394611a31616be6c6ac9b70c798c776a5d7533f25ad93f61d72db1968d3ec206.
//
// Solidity: event NewOracleRequest(string adapterId, uint256 indexed requestId)
func (_Oracle *OracleFilterer) WatchNewOracleRequest(opts *bind.WatchOpts, sink chan<- *OracleNewOracleRequest, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "NewOracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleNewOracleRequest)
				if err := _Oracle.contract.UnpackLog(event, "NewOracleRequest", log); err != nil {
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

// ParseNewOracleRequest is a log parse operation binding the contract event 0x394611a31616be6c6ac9b70c798c776a5d7533f25ad93f61d72db1968d3ec206.
//
// Solidity: event NewOracleRequest(string adapterId, uint256 indexed requestId)
func (_Oracle *OracleFilterer) ParseNewOracleRequest(log types.Log) (*OracleNewOracleRequest, error) {
	event := new(OracleNewOracleRequest)
	if err := _Oracle.contract.UnpackLog(event, "NewOracleRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}
