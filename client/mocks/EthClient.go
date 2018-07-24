// Code generated by mockery v1.0.0
package mocks

import big "math/big"

import common "github.com/ethereum/go-ethereum/common"
import context "context"
import ethereum "github.com/ethereum/go-ethereum"
import mock "github.com/stretchr/testify/mock"
import model "github.com/getamis/eth-indexer/model"
import types "github.com/ethereum/go-ethereum/core/types"

// EthClient is an autogenerated mock type for the EthClient type
type EthClient struct {
	mock.Mock
}

// BalanceAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *EthClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, account, blockNumber)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) *big.Int); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BalanceOf provides a mock function with given fields: _a0, _a1, _a2
func (_m *EthClient) BalanceOf(_a0 context.Context, _a1 *big.Int, _a2 map[common.Address]map[common.Address]struct{}) (map[common.Address]map[common.Address]*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 map[common.Address]map[common.Address]*big.Int
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, map[common.Address]map[common.Address]struct{}) map[common.Address]map[common.Address]*big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[common.Address]map[common.Address]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, map[common.Address]map[common.Address]struct{}) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchBalanceAt provides a mock function with given fields: ctx, accounts, blockNumber
func (_m *EthClient) BatchBalanceAt(ctx context.Context, accounts []common.Address, blockNumber *big.Int) ([]*big.Int, error) {
	ret := _m.Called(ctx, accounts, blockNumber)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func(context.Context, []common.Address, *big.Int) []*big.Int); ok {
		r0 = rf(ctx, accounts, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []common.Address, *big.Int) error); ok {
		r1 = rf(ctx, accounts, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchCallContract provides a mock function with given fields: ctx, msgs, blockNumber
func (_m *EthClient) BatchCallContract(ctx context.Context, msgs []*ethereum.CallMsg, blockNumber *big.Int) ([][]byte, error) {
	ret := _m.Called(ctx, msgs, blockNumber)

	var r0 [][]byte
	if rf, ok := ret.Get(0).(func(context.Context, []*ethereum.CallMsg, *big.Int) [][]byte); ok {
		r0 = rf(ctx, msgs, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, msgs, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *EthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByNumber provides a mock function with given fields: ctx, number
func (_m *EthClient) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CallContract provides a mock function with given fields: ctx, msg, blockNumber
func (_m *EthClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, msg, blockNumber)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, msg, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, msg, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *EthClient) Close() {
	_m.Called()
}

// CodeAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *EthClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, account, blockNumber)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockReceipts provides a mock function with given fields: ctx, hash
func (_m *EthClient) GetBlockReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	ret := _m.Called(ctx, hash)

	var r0 types.Receipts
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) types.Receipts); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Receipts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetERC20 provides a mock function with given fields: ctx, addr
func (_m *EthClient) GetERC20(ctx context.Context, addr common.Address) (*model.ERC20, error) {
	ret := _m.Called(ctx, addr)

	var r0 *model.ERC20
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) *model.ERC20); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ERC20)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalDifficulty provides a mock function with given fields: ctx, hash
func (_m *EthClient) GetTotalDifficulty(ctx context.Context, hash common.Hash) (*big.Int, error) {
	ret := _m.Called(ctx, hash)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *big.Int); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransferLogs provides a mock function with given fields: ctx, hash
func (_m *EthClient) GetTransferLogs(ctx context.Context, hash common.Hash) ([]*types.TransferLog, error) {
	ret := _m.Called(ctx, hash)

	var r0 []*types.TransferLog
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) []*types.TransferLog); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.TransferLog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeNewHead provides a mock function with given fields: ctx, ch
func (_m *EthClient) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, ch)

	var r0 ethereum.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, chan<- *types.Header) ethereum.Subscription); ok {
		r0 = rf(ctx, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- *types.Header) error); ok {
		r1 = rf(ctx, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionByHash provides a mock function with given fields: ctx, hash
func (_m *EthClient) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	ret := _m.Called(ctx, hash)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Transaction); ok {
		r0 = rf(ctx, hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) bool); ok {
		r1 = rf(ctx, hash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, hash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TransactionReceipt provides a mock function with given fields: ctx, txHash
func (_m *EthClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UncleByBlockHashAndPosition provides a mock function with given fields: ctx, hash, position
func (_m *EthClient) UncleByBlockHashAndPosition(ctx context.Context, hash common.Hash, position uint) (*types.Header, error) {
	ret := _m.Called(ctx, hash, position)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint) *types.Header); ok {
		r0 = rf(ctx, hash, position)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint) error); ok {
		r1 = rf(ctx, hash, position)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnclesByBlockHash provides a mock function with given fields: ctx, blockHash
func (_m *EthClient) UnclesByBlockHash(ctx context.Context, blockHash common.Hash) ([]*types.Header, error) {
	ret := _m.Called(ctx, blockHash)

	var r0 []*types.Header
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) []*types.Header); ok {
		r0 = rf(ctx, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
