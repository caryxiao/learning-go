// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simple_storage

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

// SimpleStorageMetaData contains all meta data concerning the SimpleStorage contract.
var SimpleStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"NumberStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isContractWriteable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setReadonly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setWriteable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f60025f6101000a81548160ff0219169083151502179055503480156027575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167f342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a73560405160405180910390a361055e806100ee5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80632e64cec11461006457806331ba3990146100825780636057361d1461008c578063893d20e8146100a85780639ea05385146100c6578063a884255c146100e4575b5f5ffd5b61006c6100ee565b604051610079919061034b565b60405180910390f35b61008a6100f7565b005b6100a660048036038101906100a19190610392565b6101a0565b005b6100b061024d565b6040516100bd91906103fc565b60405180910390f35b6100ce610274565b6040516100db919061042f565b60405180910390f35b6100ec610289565b005b5f600154905090565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610185576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161017c906104a2565b60405180910390fd5b5f60025f6101000a81548160ff021916908315150217905550565b6001151560025f9054906101000a900460ff161515146101f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ec9061050a565b60405180910390fd5b806001819055503373ffffffffffffffffffffffffffffffffffffffff167fed65cf72675c57660e67b5e615ec8981e01e5deebce7713b9ee2a1c991e5eb2682604051610242919061034b565b60405180910390a250565b5f5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60025f9054906101000a900460ff16905090565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610317576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030e906104a2565b60405180910390fd5b600160025f6101000a81548160ff021916908315150217905550565b5f819050919050565b61034581610333565b82525050565b5f60208201905061035e5f83018461033c565b92915050565b5f5ffd5b61037181610333565b811461037b575f5ffd5b50565b5f8135905061038c81610368565b92915050565b5f602082840312156103a7576103a6610364565b5b5f6103b48482850161037e565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6103e6826103bd565b9050919050565b6103f6816103dc565b82525050565b5f60208201905061040f5f8301846103ed565b92915050565b5f8115159050919050565b61042981610415565b82525050565b5f6020820190506104425f830184610420565b92915050565b5f82825260208201905092915050565b7f43616c6c6572206973206e6f74206f776e6572000000000000000000000000005f82015250565b5f61048c601383610448565b915061049782610458565b602082019050919050565b5f6020820190508181035f8301526104b981610480565b9050919050565b7f436f6e7472616374206973206e6f7420777269746561626c65000000000000005f82015250565b5f6104f4601983610448565b91506104ff826104c0565b602082019050919050565b5f6020820190508181035f830152610521816104e8565b905091905056fea2646970667358221220ab99aa2c2d01f2aa0196b25ddad49c88cb29216fcb956bf81d76014e097b82e564736f6c634300081e0033",
}

// SimpleStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleStorageMetaData.ABI instead.
var SimpleStorageABI = SimpleStorageMetaData.ABI

// SimpleStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleStorageMetaData.Bin instead.
var SimpleStorageBin = SimpleStorageMetaData.Bin

// DeploySimpleStorage deploys a new Ethereum contract, binding an instance of SimpleStorage to it.
func DeploySimpleStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimpleStorage, error) {
	parsed, err := SimpleStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// SimpleStorage is an auto generated Go binding around an Ethereum contract.
type SimpleStorage struct {
	SimpleStorageCaller     // Read-only binding to the contract
	SimpleStorageTransactor // Write-only binding to the contract
	SimpleStorageFilterer   // Log filterer for contract events
}

// SimpleStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleStorageSession struct {
	Contract     *SimpleStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleStorageCallerSession struct {
	Contract *SimpleStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleStorageTransactorSession struct {
	Contract     *SimpleStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleStorageRaw struct {
	Contract *SimpleStorage // Generic contract binding to access the raw methods on
}

// SimpleStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleStorageCallerRaw struct {
	Contract *SimpleStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleStorageTransactorRaw struct {
	Contract *SimpleStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleStorage creates a new instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorage(address common.Address, backend bind.ContractBackend) (*SimpleStorage, error) {
	contract, err := bindSimpleStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleStorage{SimpleStorageCaller: SimpleStorageCaller{contract: contract}, SimpleStorageTransactor: SimpleStorageTransactor{contract: contract}, SimpleStorageFilterer: SimpleStorageFilterer{contract: contract}}, nil
}

// NewSimpleStorageCaller creates a new read-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageCaller(address common.Address, caller bind.ContractCaller) (*SimpleStorageCaller, error) {
	contract, err := bindSimpleStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageCaller{contract: contract}, nil
}

// NewSimpleStorageTransactor creates a new write-only instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleStorageTransactor, error) {
	contract, err := bindSimpleStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageTransactor{contract: contract}, nil
}

// NewSimpleStorageFilterer creates a new log filterer instance of SimpleStorage, bound to a specific deployed contract.
func NewSimpleStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleStorageFilterer, error) {
	contract, err := bindSimpleStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageFilterer{contract: contract}, nil
}

// bindSimpleStorage binds a generic wrapper to an already deployed contract.
func bindSimpleStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.SimpleStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.SimpleStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleStorage *SimpleStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleStorage *SimpleStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleStorage.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SimpleStorage *SimpleStorageCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SimpleStorage *SimpleStorageSession) GetOwner() (common.Address, error) {
	return _SimpleStorage.Contract.GetOwner(&_SimpleStorage.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_SimpleStorage *SimpleStorageCallerSession) GetOwner() (common.Address, error) {
	return _SimpleStorage.Contract.GetOwner(&_SimpleStorage.CallOpts)
}

// IsContractWriteable is a free data retrieval call binding the contract method 0x9ea05385.
//
// Solidity: function isContractWriteable() view returns(bool)
func (_SimpleStorage *SimpleStorageCaller) IsContractWriteable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "isContractWriteable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractWriteable is a free data retrieval call binding the contract method 0x9ea05385.
//
// Solidity: function isContractWriteable() view returns(bool)
func (_SimpleStorage *SimpleStorageSession) IsContractWriteable() (bool, error) {
	return _SimpleStorage.Contract.IsContractWriteable(&_SimpleStorage.CallOpts)
}

// IsContractWriteable is a free data retrieval call binding the contract method 0x9ea05385.
//
// Solidity: function isContractWriteable() view returns(bool)
func (_SimpleStorage *SimpleStorageCallerSession) IsContractWriteable() (bool, error) {
	return _SimpleStorage.Contract.IsContractWriteable(&_SimpleStorage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_SimpleStorage *SimpleStorageCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleStorage.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_SimpleStorage *SimpleStorageSession) Retrieve() (*big.Int, error) {
	return _SimpleStorage.Contract.Retrieve(&_SimpleStorage.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_SimpleStorage *SimpleStorageCallerSession) Retrieve() (*big.Int, error) {
	return _SimpleStorage.Contract.Retrieve(&_SimpleStorage.CallOpts)
}

// SetReadonly is a paid mutator transaction binding the contract method 0x31ba3990.
//
// Solidity: function setReadonly() returns()
func (_SimpleStorage *SimpleStorageTransactor) SetReadonly(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "setReadonly")
}

// SetReadonly is a paid mutator transaction binding the contract method 0x31ba3990.
//
// Solidity: function setReadonly() returns()
func (_SimpleStorage *SimpleStorageSession) SetReadonly() (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetReadonly(&_SimpleStorage.TransactOpts)
}

// SetReadonly is a paid mutator transaction binding the contract method 0x31ba3990.
//
// Solidity: function setReadonly() returns()
func (_SimpleStorage *SimpleStorageTransactorSession) SetReadonly() (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetReadonly(&_SimpleStorage.TransactOpts)
}

// SetWriteable is a paid mutator transaction binding the contract method 0xa884255c.
//
// Solidity: function setWriteable() returns()
func (_SimpleStorage *SimpleStorageTransactor) SetWriteable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "setWriteable")
}

// SetWriteable is a paid mutator transaction binding the contract method 0xa884255c.
//
// Solidity: function setWriteable() returns()
func (_SimpleStorage *SimpleStorageSession) SetWriteable() (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetWriteable(&_SimpleStorage.TransactOpts)
}

// SetWriteable is a paid mutator transaction binding the contract method 0xa884255c.
//
// Solidity: function setWriteable() returns()
func (_SimpleStorage *SimpleStorageTransactorSession) SetWriteable() (*types.Transaction, error) {
	return _SimpleStorage.Contract.SetWriteable(&_SimpleStorage.TransactOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_SimpleStorage *SimpleStorageTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_SimpleStorage *SimpleStorageSession) Store(num *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.Contract.Store(&_SimpleStorage.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_SimpleStorage *SimpleStorageTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _SimpleStorage.Contract.Store(&_SimpleStorage.TransactOpts, num)
}

// SimpleStorageNumberStoredIterator is returned from FilterNumberStored and is used to iterate over the raw logs and unpacked data for NumberStored events raised by the SimpleStorage contract.
type SimpleStorageNumberStoredIterator struct {
	Event *SimpleStorageNumberStored // Event containing the contract specifics and raw log

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
func (it *SimpleStorageNumberStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleStorageNumberStored)
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
		it.Event = new(SimpleStorageNumberStored)
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
func (it *SimpleStorageNumberStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleStorageNumberStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleStorageNumberStored represents a NumberStored event raised by the SimpleStorage contract.
type SimpleStorageNumberStored struct {
	Sender   common.Address
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNumberStored is a free log retrieval operation binding the contract event 0xed65cf72675c57660e67b5e615ec8981e01e5deebce7713b9ee2a1c991e5eb26.
//
// Solidity: event NumberStored(address indexed sender, uint256 newValue)
func (_SimpleStorage *SimpleStorageFilterer) FilterNumberStored(opts *bind.FilterOpts, sender []common.Address) (*SimpleStorageNumberStoredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleStorage.contract.FilterLogs(opts, "NumberStored", senderRule)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageNumberStoredIterator{contract: _SimpleStorage.contract, event: "NumberStored", logs: logs, sub: sub}, nil
}

// WatchNumberStored is a free log subscription operation binding the contract event 0xed65cf72675c57660e67b5e615ec8981e01e5deebce7713b9ee2a1c991e5eb26.
//
// Solidity: event NumberStored(address indexed sender, uint256 newValue)
func (_SimpleStorage *SimpleStorageFilterer) WatchNumberStored(opts *bind.WatchOpts, sink chan<- *SimpleStorageNumberStored, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SimpleStorage.contract.WatchLogs(opts, "NumberStored", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleStorageNumberStored)
				if err := _SimpleStorage.contract.UnpackLog(event, "NumberStored", log); err != nil {
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

// ParseNumberStored is a log parse operation binding the contract event 0xed65cf72675c57660e67b5e615ec8981e01e5deebce7713b9ee2a1c991e5eb26.
//
// Solidity: event NumberStored(address indexed sender, uint256 newValue)
func (_SimpleStorage *SimpleStorageFilterer) ParseNumberStored(log types.Log) (*SimpleStorageNumberStored, error) {
	event := new(SimpleStorageNumberStored)
	if err := _SimpleStorage.contract.UnpackLog(event, "NumberStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleStorageOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the SimpleStorage contract.
type SimpleStorageOwnerSetIterator struct {
	Event *SimpleStorageOwnerSet // Event containing the contract specifics and raw log

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
func (it *SimpleStorageOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleStorageOwnerSet)
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
		it.Event = new(SimpleStorageOwnerSet)
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
func (it *SimpleStorageOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleStorageOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleStorageOwnerSet represents a OwnerSet event raised by the SimpleStorage contract.
type SimpleStorageOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_SimpleStorage *SimpleStorageFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*SimpleStorageOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SimpleStorage.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleStorageOwnerSetIterator{contract: _SimpleStorage.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_SimpleStorage *SimpleStorageFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *SimpleStorageOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SimpleStorage.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleStorageOwnerSet)
				if err := _SimpleStorage.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// ParseOwnerSet is a log parse operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_SimpleStorage *SimpleStorageFilterer) ParseOwnerSet(log types.Log) (*SimpleStorageOwnerSet, error) {
	event := new(SimpleStorageOwnerSet)
	if err := _SimpleStorage.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
