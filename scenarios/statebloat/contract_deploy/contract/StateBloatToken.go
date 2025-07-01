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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PADDING_DATA\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy10\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy11\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy12\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy13\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy14\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy15\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy16\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy17\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy18\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy19\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy20\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy21\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy22\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy23\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy24\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy25\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy26\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy27\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy28\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy29\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy30\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy31\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy32\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy33\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy34\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy35\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy36\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy37\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy38\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy39\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy4\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy40\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy41\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy42\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy43\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy44\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy45\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy46\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy47\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy48\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy49\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy5\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy50\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy51\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy52\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy53\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy54\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy55\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy56\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy57\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy58\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy59\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy6\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy7\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy8\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dummy9\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom10\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom11\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom12\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom13\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom14\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom15\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom16\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom17\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom18\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom19\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom3\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom4\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom5\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom6\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom7\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom8\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom9\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051615fe8380380615fe8833981810160405281019061003191906101f4565b6040518060400160405280601181526020017f537461746520426c6f617420546f6b656e0000000000000000000000000000008152505f90816100749190610453565b506040518060400160405280600381526020017f5342540000000000000000000000000000000000000000000000000000000000815250600190816100b99190610453565b50601260025f6101000a81548160ff021916908360ff160217905550806080818152505060025f9054906101000a900460ff16600a6100f8919061068a565b620f424061010691906106d4565b60038190555060035460045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6003546040516101af9190610724565b60405180910390a35061073d565b5f5ffd5b5f819050919050565b6101d3816101c1565b81146101dd575f5ffd5b50565b5f815190506101ee816101ca565b92915050565b5f60208284031215610209576102086101bd565b5b5f610216848285016101e0565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061029a57607f821691505b6020821081036102ad576102ac610256565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261030f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826102d4565b61031986836102d4565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61035461034f61034a846101c1565b610331565b6101c1565b9050919050565b5f819050919050565b61036d8361033a565b6103816103798261035b565b8484546102e0565b825550505050565b5f5f905090565b610398610389565b6103a3818484610364565b505050565b5b818110156103c6576103bb5f82610390565b6001810190506103a9565b5050565b601f82111561040b576103dc816102b3565b6103e5846102c5565b810160208510156103f4578190505b610408610400856102c5565b8301826103a8565b50505b505050565b5f82821c905092915050565b5f61042b5f1984600802610410565b1980831691505092915050565b5f610443838361041c565b9150826002028217905092915050565b61045c8261021f565b67ffffffffffffffff81111561047557610474610229565b5b61047f8254610283565b61048a8282856103ca565b5f60209050601f8311600181146104bb575f84156104a9578287015190505b6104b38582610438565b86555061051a565b601f1984166104c9866102b3565b5f5b828110156104f0578489015182556001820191506020850194506020810190506104cb565b8683101561050d5784890151610509601f89168261041c565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b60018511156105a4578086048111156105805761057f610522565b5b600185161561058f5780820291505b808102905061059d8561054f565b9450610564565b94509492505050565b5f826105bc5760019050610677565b816105c9575f9050610677565b81600181146105df57600281146105e957610618565b6001915050610677565b60ff8411156105fb576105fa610522565b5b8360020a91508482111561061257610611610522565b5b50610677565b5060208310610133831016604e8410600b841016171561064d5782820a90508381111561064857610647610522565b5b610677565b61065a848484600161055b565b9250905081840481111561067157610670610522565b5b81810290505b9392505050565b5f60ff82169050919050565b5f610694826101c1565b915061069f8361067e565b92506106cc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846105ad565b905092915050565b5f6106de826101c1565b91506106e9836101c1565b92508282026106f7816101c1565b9150828204841483151761070e5761070d610522565b5b5092915050565b61071e816101c1565b82525050565b5f6020820190506107375f830184610715565b92915050565b6080516158936107555f395f61475b01526158935ff3fe608060405234801561000f575f5ffd5b50600436106104ec575f3560e01c8063657b6ef71161028c578063aaa7af7011610165578063d3c1c838116100d7578063ee1682b611610090578063ee1682b614611064578063f26c779b14611082578063f8716f14146110a0578063faf35ced146110be578063fe7d5996146110ee578063ffbf04691461110c576104ec565b8063d3c1c83814610f8e578063d741946914610faa578063dc1d8a9b14610fc8578063dd62ed3e14610fe6578063e8c927b314611016578063eb4329c814611034576104ec565b8063bb9bfe0611610129578063bb9bfe0614610eb6578063bfa0b13314610ee6578063c2be97e314610f04578063c958d4bf14610f34578063cfd6686314610f52578063d101dcd014610f70576104ec565b8063aaa7af7014610dfc578063acc5aee914610e1a578063b1802b9a14610e4a578063b2bb360e14610e7a578063b66dd75014610e98576104ec565b80637dffdc32116101fe5780638f4a8406116101c25780638f4a840614610d3657806395d89b4114610d545780639c5dfe7314610d725780639df61a2514610d90578063a891d4d414610dae578063a9059cbb14610dcc576104ec565b80637dffdc3214610c8e5780637e54493714610cac5780637f34d94b14610cca5780638619d60714610ce85780638789ca6714610d06576104ec565b806374e73fd31161025057806374e73fd314610bc857806374f83d0214610be657806377c0209e14610c04578063792c7f3e14610c225780637c66673e14610c405780637c72ed0d14610c70576104ec565b8063657b6ef714610afc578063672151fe14610b2c5780636abceacd14610b4a5780636c12ed2814610b6857806370a0823114610b98576104ec565b80633125f37a116103c95780634a2e93c61161033b578063552a1b56116102f4578063552a1b5614610a3657806358b6a9bd14610a665780635af92c0514610a8457806361b970eb14610aa2578063639ec53a14610ac05780636547317414610ade576104ec565b80634a2e93c6146109705780634b3c7f5f1461098e5780634e1dbb82146109ac5780634f5e5557146109ca5780634f7bd75a146109e857806354c2792014610a18576104ec565b80633ea117ce1161038d5780633ea117ce146108aa5780634128a85d146108c8578063418b1816146108e657806342937dbd1461090457806343a6b92d1461092257806344050a2814610940576104ec565b80633125f37a146107f0578063313ce5671461080e57806339e0bd121461082c5780633a1319901461085c5780633b6be4591461088c576104ec565b80631a97f18e116104625780631fd298ec116104265780631fd298ec1461071857806321ecd7a314610736578063239af2a51461075457806323b872dd146107725780632545d8b7146107a2578063291c3bd7146107c0576104ec565b80631a97f18e1461065e5780631b17c65c1461067c5780631bbffe6f146106ac5780631d527cde146106dc5780631eaa7c52146106fa576104ec565b80631215a3ab116104b45780631215a3ab146105aa57806312901b42146105c857806313ebb5ec146105e657806316a3045b1461060457806318160ddd1461062257806319cf6a9114610640576104ec565b80630460faf6146104f057806306fdde03146105205780630717b1611461053e578063095ea7b31461055c5780630cb7a9e71461058c575b5f5ffd5b61050a600480360381019061050591906151a8565b61112a565b6040516105179190615212565b60405180910390f35b61052861140a565b604051610535919061529b565b60405180910390f35b610546611495565b60405161055391906152ca565b60405180910390f35b610576600480360381019061057191906152e3565b61149d565b6040516105839190615212565b60405180910390f35b61059461158a565b6040516105a191906152ca565b60405180910390f35b6105b2611592565b6040516105bf91906152ca565b60405180910390f35b6105d061159a565b6040516105dd91906152ca565b60405180910390f35b6105ee6115a2565b6040516105fb91906152ca565b60405180910390f35b61060c6115aa565b60405161061991906152ca565b60405180910390f35b61062a6115b2565b60405161063791906152ca565b60405180910390f35b6106486115b8565b60405161065591906152ca565b60405180910390f35b6106666115c0565b60405161067391906152ca565b60405180910390f35b610696600480360381019061069191906151a8565b6115c8565b6040516106a39190615212565b60405180910390f35b6106c660048036038101906106c191906151a8565b6118a8565b6040516106d39190615212565b60405180910390f35b6106e4611b88565b6040516106f191906152ca565b60405180910390f35b610702611b90565b60405161070f91906152ca565b60405180910390f35b610720611b98565b60405161072d91906152ca565b60405180910390f35b61073e611ba0565b60405161074b91906152ca565b60405180910390f35b61075c611ba8565b60405161076991906152ca565b60405180910390f35b61078c600480360381019061078791906151a8565b611bb0565b6040516107999190615212565b60405180910390f35b6107aa611e90565b6040516107b791906152ca565b60405180910390f35b6107da60048036038101906107d591906151a8565b611e98565b6040516107e79190615212565b60405180910390f35b6107f8612178565b60405161080591906152ca565b60405180910390f35b610816612180565b604051610823919061533c565b60405180910390f35b610846600480360381019061084191906151a8565b612192565b6040516108539190615212565b60405180910390f35b610876600480360381019061087191906151a8565b612472565b6040516108839190615212565b60405180910390f35b610894612752565b6040516108a191906152ca565b60405180910390f35b6108b261275a565b6040516108bf91906152ca565b60405180910390f35b6108d0612762565b6040516108dd91906152ca565b60405180910390f35b6108ee61276a565b6040516108fb91906152ca565b60405180910390f35b61090c612772565b60405161091991906152ca565b60405180910390f35b61092a61277a565b60405161093791906152ca565b60405180910390f35b61095a600480360381019061095591906151a8565b612782565b6040516109679190615212565b60405180910390f35b610978612a62565b60405161098591906152ca565b60405180910390f35b610996612a6a565b6040516109a391906152ca565b60405180910390f35b6109b4612a72565b6040516109c191906152ca565b60405180910390f35b6109d2612a7a565b6040516109df91906152ca565b60405180910390f35b610a0260048036038101906109fd91906151a8565b612a82565b604051610a0f9190615212565b60405180910390f35b610a20612d62565b604051610a2d91906152ca565b60405180910390f35b610a506004803603810190610a4b91906151a8565b612d6a565b604051610a5d9190615212565b60405180910390f35b610a6e61304a565b604051610a7b91906152ca565b60405180910390f35b610a8c613052565b604051610a9991906152ca565b60405180910390f35b610aaa61305a565b604051610ab791906152ca565b60405180910390f35b610ac8613062565b604051610ad591906152ca565b60405180910390f35b610ae661306a565b604051610af391906152ca565b60405180910390f35b610b166004803603810190610b1191906151a8565b613072565b604051610b239190615212565b60405180910390f35b610b34613352565b604051610b4191906152ca565b60405180910390f35b610b5261335a565b604051610b5f91906152ca565b60405180910390f35b610b826004803603810190610b7d91906151a8565b613362565b604051610b8f9190615212565b60405180910390f35b610bb26004803603810190610bad9190615355565b613642565b604051610bbf91906152ca565b60405180910390f35b610bd0613657565b604051610bdd91906152ca565b60405180910390f35b610bee61365f565b604051610bfb91906152ca565b60405180910390f35b610c0c613667565b604051610c1991906152ca565b60405180910390f35b610c2a61366f565b604051610c3791906152ca565b60405180910390f35b610c5a6004803603810190610c5591906151a8565b613677565b604051610c679190615212565b60405180910390f35b610c78613957565b604051610c8591906152ca565b60405180910390f35b610c9661395f565b604051610ca391906152ca565b60405180910390f35b610cb4613967565b604051610cc191906152ca565b60405180910390f35b610cd261396f565b604051610cdf91906152ca565b60405180910390f35b610cf0613977565b604051610cfd91906152ca565b60405180910390f35b610d206004803603810190610d1b91906151a8565b61397f565b604051610d2d9190615212565b60405180910390f35b610d3e613c5f565b604051610d4b91906152ca565b60405180910390f35b610d5c613c67565b604051610d69919061529b565b60405180910390f35b610d7a613cf3565b604051610d8791906152ca565b60405180910390f35b610d98613cfb565b604051610da591906152ca565b60405180910390f35b610db6613d03565b604051610dc391906152ca565b60405180910390f35b610de66004803603810190610de191906152e3565b613d0b565b604051610df39190615212565b60405180910390f35b610e04613ea1565b604051610e1191906152ca565b60405180910390f35b610e346004803603810190610e2f91906151a8565b613ea9565b604051610e419190615212565b60405180910390f35b610e646004803603810190610e5f91906151a8565b614189565b604051610e719190615212565b60405180910390f35b610e82614469565b604051610e8f91906152ca565b60405180910390f35b610ea0614471565b604051610ead91906152ca565b60405180910390f35b610ed06004803603810190610ecb91906151a8565b614479565b604051610edd9190615212565b60405180910390f35b610eee614759565b604051610efb91906152ca565b60405180910390f35b610f1e6004803603810190610f1991906151a8565b61477d565b604051610f2b9190615212565b60405180910390f35b610f3c614a5d565b604051610f4991906152ca565b60405180910390f35b610f5a614a65565b604051610f6791906152ca565b60405180910390f35b610f78614a6d565b604051610f8591906152ca565b60405180910390f35b610fa86004803603810190610fa391906153e1565b614a75565b005b610fb2614b97565b604051610fbf91906152ca565b60405180910390f35b610fd0614b9f565b604051610fdd91906152ca565b60405180910390f35b6110006004803603810190610ffb919061542c565b614ba7565b60405161100d91906152ca565b60405180910390f35b61101e614bc7565b60405161102b91906152ca565b60405180910390f35b61104e600480360381019061104991906151a8565b614bcf565b60405161105b9190615212565b60405180910390f35b61106c614df4565b604051611079919061529b565b60405180910390f35b61108a614e13565b60405161109791906152ca565b60405180910390f35b6110a8614e1b565b6040516110b591906152ca565b60405180910390f35b6110d860048036038101906110d391906151a8565b614e23565b6040516110e59190615212565b60405180910390f35b6110f6615103565b60405161110391906152ca565b60405180910390f35b61111461510b565b60405161112191906152ca565b60405180910390f35b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156111ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111a2906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611266576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125d9061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546112b29190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611305919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546113939190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113f791906152ca565b60405180910390a3600190509392505050565b5f8054611416906155fa565b80601f0160208091040260200160405190810160405280929190818152602001828054611442906155fa565b801561148d5780601f106114645761010080835404028352916020019161148d565b820191905f5260205f20905b81548152906001019060200180831161147057829003601f168201915b505050505081565b5f602f905090565b5f8160055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161157891906152ca565b60405180910390a36001905092915050565b5f601b905090565b5f6005905090565b5f601a905090565b5f601d905090565b5f6033905090565b60035481565b5f6008905090565b5f6003905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611649576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611640906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611704576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116fb9061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546117509190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546117a3919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546118319190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161189591906152ca565b60405180910390a3600190509392505050565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611929576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611920906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156119e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119db9061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611a309190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611a83919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611b119190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611b7591906152ca565b60405180910390a3600190509392505050565b5f6002905090565b5f6010905090565b5f6035905090565b5f602e905090565b5f602c905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611c31576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c28906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611cec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ce39061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611d389190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611d8b919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254611e199190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611e7d91906152ca565b60405180910390a3600190509392505050565b5f6001905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611f19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f1090615674565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015611fd4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fcb906156dc565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546120209190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612073919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546121019190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161216591906152ca565b60405180910390a3600190509392505050565b5f6034905090565b60025f9054906101000a900460ff1681565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015612213576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161220a906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156122ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122c59061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461231a9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461236d919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546123fb9190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161245f91906152ca565b60405180910390a3600190509392505050565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156124f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124ea906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156125ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125a59061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546125fa9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461264d919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546126db9190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161273f91906152ca565b60405180910390a3600190509392505050565b5f6004905090565b5f600c905090565b5f6006905090565b5f601c905090565b5f6032905090565b5f6023905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015612803576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127fa906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156128be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128b59061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461290a9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461295d919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546129eb9190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051612a4f91906152ca565b60405180910390a3600190509392505050565b5f6011905090565b5f6021905090565b5f601f905090565b5f6026905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015612b03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612afa906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015612bbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612bb59061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612c0a9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612c5d919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612ceb9190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051612d4f91906152ca565b60405180910390a3600190509392505050565b5f603a905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015612deb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612de2906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015612ea6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e9d9061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612ef29190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612f45919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254612fd39190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161303791906152ca565b60405180910390a3600190509392505050565b5f6009905090565b5f602a905090565b5f6014905090565b5f6025905090565b5f6013905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156130f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130ea90615674565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156131ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131a5906156dc565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546131fa9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461324d919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546132db9190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161333f91906152ca565b60405180910390a3600190509392505050565b5f600d905090565b5f6017905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156133e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133da906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561349e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134959061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546134ea9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461353d919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546135cb9190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161362f91906152ca565b60405180910390a3600190509392505050565b6004602052805f5260405f205f915090505481565b5f6016905090565b5f6039905090565b5f603b905090565b5f6036905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156136f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136ef906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156137b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137aa9061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546137ff9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254613852919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546138e09190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161394491906152ca565b60405180910390a3600190509392505050565b5f602b905090565b5f6037905090565b5f600a905090565b5f602d905090565b5f6030905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015613a00576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139f7906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015613abb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ab29061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254613b079190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254613b5a919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254613be89190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051613c4c91906152ca565b60405180910390a3600190509392505050565b5f600b905090565b60018054613c74906155fa565b80601f0160208091040260200160405190810160405280929190818152602001828054613ca0906155fa565b8015613ceb5780601f10613cc257610100808354040283529160200191613ceb565b820191905f5260205f20905b815481529060010190602001808311613cce57829003601f168201915b505050505081565b5f6015905090565b5f6029905090565b5f6038905090565b5f8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015613d8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d83906154b4565b60405180910390fd5b8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254613dd89190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254613e2b919061559a565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051613e8f91906152ca565b60405180910390a36001905092915050565b5f6028905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015613f2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613f21906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015613fe5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613fdc9061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546140319190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614084919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546141129190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161417691906152ca565b60405180910390a3600190509392505050565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561420a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614201906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156142c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016142bc9061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546143119190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614364919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546143f29190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161445691906152ca565b60405180910390a3600190509392505050565b5f6018905090565b5f600f905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156144fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016144f1906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156145b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016145ac9061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546146019190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614654919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546146e29190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161474691906152ca565b60405180910390a3600190509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156147fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016147f5906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156148b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016148b09061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546149059190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614958919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546149e69190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051614a4a91906152ca565b60405180910390a3600190509392505050565b5f6024905090565b5f601e905090565b5f6031905090565b5f600190505f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f5f90505b84849050811015614b4e5782820391508260045f878785818110614ae757614ae66156fa565b5b9050602002016020810190614afc9190615355565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055508080600101915050614ac0565b508060045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f208190555050505050565b5f6019905090565b5f6027905090565b6005602052815f5260405f20602052805f5260405f205f91509150505481565b5f6022905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015614c50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614c4790615674565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614c9c9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614cef919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614d7d9190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051614de191906152ca565b60405180910390a3600190509392505050565b6040518061016001604052806101368152602001615728610136913981565b5f6007905090565b5f6020905090565b5f8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015614ea4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614e9b906154b4565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015614f5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614f569061551c565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614fab9190615567565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f828254614ffe919061559a565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461508c9190615567565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516150f091906152ca565b60405180910390a3600190509392505050565b5f600e905090565b5f6012905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6151448261511b565b9050919050565b6151548161513a565b811461515e575f5ffd5b50565b5f8135905061516f8161514b565b92915050565b5f819050919050565b61518781615175565b8114615191575f5ffd5b50565b5f813590506151a28161517e565b92915050565b5f5f5f606084860312156151bf576151be615113565b5b5f6151cc86828701615161565b93505060206151dd86828701615161565b92505060406151ee86828701615194565b9150509250925092565b5f8115159050919050565b61520c816151f8565b82525050565b5f6020820190506152255f830184615203565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61526d8261522b565b6152778185615235565b9350615287818560208601615245565b61529081615253565b840191505092915050565b5f6020820190508181035f8301526152b38184615263565b905092915050565b6152c481615175565b82525050565b5f6020820190506152dd5f8301846152bb565b92915050565b5f5f604083850312156152f9576152f8615113565b5b5f61530685828601615161565b925050602061531785828601615194565b9150509250929050565b5f60ff82169050919050565b61533681615321565b82525050565b5f60208201905061534f5f83018461532d565b92915050565b5f6020828403121561536a57615369615113565b5b5f61537784828501615161565b91505092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126153a1576153a0615380565b5b8235905067ffffffffffffffff8111156153be576153bd615384565b5b6020830191508360208202830111156153da576153d9615388565b5b9250929050565b5f5f602083850312156153f7576153f6615113565b5b5f83013567ffffffffffffffff81111561541457615413615117565b5b6154208582860161538c565b92509250509250929050565b5f5f6040838503121561544257615441615113565b5b5f61544f85828601615161565b925050602061546085828601615161565b9150509250929050565b7f496e73756666696369656e742062616c616e63650000000000000000000000005f82015250565b5f61549e601483615235565b91506154a98261546a565b602082019050919050565b5f6020820190508181035f8301526154cb81615492565b9050919050565b7f496e73756666696369656e7420616c6c6f77616e6365000000000000000000005f82015250565b5f615506601683615235565b9150615511826154d2565b602082019050919050565b5f6020820190508181035f830152615533816154fa565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61557182615175565b915061557c83615175565b92508282039050818111156155945761559361553a565b5b92915050565b5f6155a482615175565b91506155af83615175565b92508282019050808211156155c7576155c661553a565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061561157607f821691505b602082108103615624576156236155cd565b5b50919050565b7f41000000000000000000000000000000000000000000000000000000000000005f82015250565b5f61565e600183615235565b91506156698261562a565b602082019050919050565b5f6020820190508181035f83015261568b81615652565b9050919050565b7f42000000000000000000000000000000000000000000000000000000000000005f82015250565b5f6156c6600183615235565b91506156d182615692565b602082019050919050565b5f6020820190508181035f8301526156f3816156ba565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfe41414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141414141a264697066735822122038fe1f89d282a80c0590d8913a838bb5df465ec740e0b92c9d3acd06a087a6ca64736f6c634300081e0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _salt *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, _salt)
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

// PADDINGDATA is a free data retrieval call binding the contract method 0xee1682b6.
//
// Solidity: function PADDING_DATA() view returns(string)
func (_Contract *ContractCaller) PADDINGDATA(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "PADDING_DATA")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PADDINGDATA is a free data retrieval call binding the contract method 0xee1682b6.
//
// Solidity: function PADDING_DATA() view returns(string)
func (_Contract *ContractSession) PADDINGDATA() (string, error) {
	return _Contract.Contract.PADDINGDATA(&_Contract.CallOpts)
}

// PADDINGDATA is a free data retrieval call binding the contract method 0xee1682b6.
//
// Solidity: function PADDING_DATA() view returns(string)
func (_Contract *ContractCallerSession) PADDINGDATA() (string, error) {
	return _Contract.Contract.PADDINGDATA(&_Contract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contract *ContractCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contract *ContractSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contract *ContractCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contract *ContractSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contract *ContractCallerSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

// Dummy1 is a free data retrieval call binding the contract method 0x2545d8b7.
//
// Solidity: function dummy1() pure returns(uint256)
func (_Contract *ContractCaller) Dummy1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy1 is a free data retrieval call binding the contract method 0x2545d8b7.
//
// Solidity: function dummy1() pure returns(uint256)
func (_Contract *ContractSession) Dummy1() (*big.Int, error) {
	return _Contract.Contract.Dummy1(&_Contract.CallOpts)
}

// Dummy1 is a free data retrieval call binding the contract method 0x2545d8b7.
//
// Solidity: function dummy1() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy1() (*big.Int, error) {
	return _Contract.Contract.Dummy1(&_Contract.CallOpts)
}

// Dummy10 is a free data retrieval call binding the contract method 0x7e544937.
//
// Solidity: function dummy10() pure returns(uint256)
func (_Contract *ContractCaller) Dummy10(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy10")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy10 is a free data retrieval call binding the contract method 0x7e544937.
//
// Solidity: function dummy10() pure returns(uint256)
func (_Contract *ContractSession) Dummy10() (*big.Int, error) {
	return _Contract.Contract.Dummy10(&_Contract.CallOpts)
}

// Dummy10 is a free data retrieval call binding the contract method 0x7e544937.
//
// Solidity: function dummy10() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy10() (*big.Int, error) {
	return _Contract.Contract.Dummy10(&_Contract.CallOpts)
}

// Dummy11 is a free data retrieval call binding the contract method 0x8f4a8406.
//
// Solidity: function dummy11() pure returns(uint256)
func (_Contract *ContractCaller) Dummy11(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy11")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy11 is a free data retrieval call binding the contract method 0x8f4a8406.
//
// Solidity: function dummy11() pure returns(uint256)
func (_Contract *ContractSession) Dummy11() (*big.Int, error) {
	return _Contract.Contract.Dummy11(&_Contract.CallOpts)
}

// Dummy11 is a free data retrieval call binding the contract method 0x8f4a8406.
//
// Solidity: function dummy11() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy11() (*big.Int, error) {
	return _Contract.Contract.Dummy11(&_Contract.CallOpts)
}

// Dummy12 is a free data retrieval call binding the contract method 0x3ea117ce.
//
// Solidity: function dummy12() pure returns(uint256)
func (_Contract *ContractCaller) Dummy12(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy12")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy12 is a free data retrieval call binding the contract method 0x3ea117ce.
//
// Solidity: function dummy12() pure returns(uint256)
func (_Contract *ContractSession) Dummy12() (*big.Int, error) {
	return _Contract.Contract.Dummy12(&_Contract.CallOpts)
}

// Dummy12 is a free data retrieval call binding the contract method 0x3ea117ce.
//
// Solidity: function dummy12() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy12() (*big.Int, error) {
	return _Contract.Contract.Dummy12(&_Contract.CallOpts)
}

// Dummy13 is a free data retrieval call binding the contract method 0x672151fe.
//
// Solidity: function dummy13() pure returns(uint256)
func (_Contract *ContractCaller) Dummy13(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy13")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy13 is a free data retrieval call binding the contract method 0x672151fe.
//
// Solidity: function dummy13() pure returns(uint256)
func (_Contract *ContractSession) Dummy13() (*big.Int, error) {
	return _Contract.Contract.Dummy13(&_Contract.CallOpts)
}

// Dummy13 is a free data retrieval call binding the contract method 0x672151fe.
//
// Solidity: function dummy13() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy13() (*big.Int, error) {
	return _Contract.Contract.Dummy13(&_Contract.CallOpts)
}

// Dummy14 is a free data retrieval call binding the contract method 0xfe7d5996.
//
// Solidity: function dummy14() pure returns(uint256)
func (_Contract *ContractCaller) Dummy14(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy14")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy14 is a free data retrieval call binding the contract method 0xfe7d5996.
//
// Solidity: function dummy14() pure returns(uint256)
func (_Contract *ContractSession) Dummy14() (*big.Int, error) {
	return _Contract.Contract.Dummy14(&_Contract.CallOpts)
}

// Dummy14 is a free data retrieval call binding the contract method 0xfe7d5996.
//
// Solidity: function dummy14() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy14() (*big.Int, error) {
	return _Contract.Contract.Dummy14(&_Contract.CallOpts)
}

// Dummy15 is a free data retrieval call binding the contract method 0xb66dd750.
//
// Solidity: function dummy15() pure returns(uint256)
func (_Contract *ContractCaller) Dummy15(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy15")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy15 is a free data retrieval call binding the contract method 0xb66dd750.
//
// Solidity: function dummy15() pure returns(uint256)
func (_Contract *ContractSession) Dummy15() (*big.Int, error) {
	return _Contract.Contract.Dummy15(&_Contract.CallOpts)
}

// Dummy15 is a free data retrieval call binding the contract method 0xb66dd750.
//
// Solidity: function dummy15() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy15() (*big.Int, error) {
	return _Contract.Contract.Dummy15(&_Contract.CallOpts)
}

// Dummy16 is a free data retrieval call binding the contract method 0x1eaa7c52.
//
// Solidity: function dummy16() pure returns(uint256)
func (_Contract *ContractCaller) Dummy16(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy16")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy16 is a free data retrieval call binding the contract method 0x1eaa7c52.
//
// Solidity: function dummy16() pure returns(uint256)
func (_Contract *ContractSession) Dummy16() (*big.Int, error) {
	return _Contract.Contract.Dummy16(&_Contract.CallOpts)
}

// Dummy16 is a free data retrieval call binding the contract method 0x1eaa7c52.
//
// Solidity: function dummy16() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy16() (*big.Int, error) {
	return _Contract.Contract.Dummy16(&_Contract.CallOpts)
}

// Dummy17 is a free data retrieval call binding the contract method 0x4a2e93c6.
//
// Solidity: function dummy17() pure returns(uint256)
func (_Contract *ContractCaller) Dummy17(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy17")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy17 is a free data retrieval call binding the contract method 0x4a2e93c6.
//
// Solidity: function dummy17() pure returns(uint256)
func (_Contract *ContractSession) Dummy17() (*big.Int, error) {
	return _Contract.Contract.Dummy17(&_Contract.CallOpts)
}

// Dummy17 is a free data retrieval call binding the contract method 0x4a2e93c6.
//
// Solidity: function dummy17() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy17() (*big.Int, error) {
	return _Contract.Contract.Dummy17(&_Contract.CallOpts)
}

// Dummy18 is a free data retrieval call binding the contract method 0xffbf0469.
//
// Solidity: function dummy18() pure returns(uint256)
func (_Contract *ContractCaller) Dummy18(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy18")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy18 is a free data retrieval call binding the contract method 0xffbf0469.
//
// Solidity: function dummy18() pure returns(uint256)
func (_Contract *ContractSession) Dummy18() (*big.Int, error) {
	return _Contract.Contract.Dummy18(&_Contract.CallOpts)
}

// Dummy18 is a free data retrieval call binding the contract method 0xffbf0469.
//
// Solidity: function dummy18() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy18() (*big.Int, error) {
	return _Contract.Contract.Dummy18(&_Contract.CallOpts)
}

// Dummy19 is a free data retrieval call binding the contract method 0x65473174.
//
// Solidity: function dummy19() pure returns(uint256)
func (_Contract *ContractCaller) Dummy19(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy19")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy19 is a free data retrieval call binding the contract method 0x65473174.
//
// Solidity: function dummy19() pure returns(uint256)
func (_Contract *ContractSession) Dummy19() (*big.Int, error) {
	return _Contract.Contract.Dummy19(&_Contract.CallOpts)
}

// Dummy19 is a free data retrieval call binding the contract method 0x65473174.
//
// Solidity: function dummy19() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy19() (*big.Int, error) {
	return _Contract.Contract.Dummy19(&_Contract.CallOpts)
}

// Dummy2 is a free data retrieval call binding the contract method 0x1d527cde.
//
// Solidity: function dummy2() pure returns(uint256)
func (_Contract *ContractCaller) Dummy2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy2 is a free data retrieval call binding the contract method 0x1d527cde.
//
// Solidity: function dummy2() pure returns(uint256)
func (_Contract *ContractSession) Dummy2() (*big.Int, error) {
	return _Contract.Contract.Dummy2(&_Contract.CallOpts)
}

// Dummy2 is a free data retrieval call binding the contract method 0x1d527cde.
//
// Solidity: function dummy2() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy2() (*big.Int, error) {
	return _Contract.Contract.Dummy2(&_Contract.CallOpts)
}

// Dummy20 is a free data retrieval call binding the contract method 0x61b970eb.
//
// Solidity: function dummy20() pure returns(uint256)
func (_Contract *ContractCaller) Dummy20(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy20")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy20 is a free data retrieval call binding the contract method 0x61b970eb.
//
// Solidity: function dummy20() pure returns(uint256)
func (_Contract *ContractSession) Dummy20() (*big.Int, error) {
	return _Contract.Contract.Dummy20(&_Contract.CallOpts)
}

// Dummy20 is a free data retrieval call binding the contract method 0x61b970eb.
//
// Solidity: function dummy20() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy20() (*big.Int, error) {
	return _Contract.Contract.Dummy20(&_Contract.CallOpts)
}

// Dummy21 is a free data retrieval call binding the contract method 0x9c5dfe73.
//
// Solidity: function dummy21() pure returns(uint256)
func (_Contract *ContractCaller) Dummy21(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy21")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy21 is a free data retrieval call binding the contract method 0x9c5dfe73.
//
// Solidity: function dummy21() pure returns(uint256)
func (_Contract *ContractSession) Dummy21() (*big.Int, error) {
	return _Contract.Contract.Dummy21(&_Contract.CallOpts)
}

// Dummy21 is a free data retrieval call binding the contract method 0x9c5dfe73.
//
// Solidity: function dummy21() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy21() (*big.Int, error) {
	return _Contract.Contract.Dummy21(&_Contract.CallOpts)
}

// Dummy22 is a free data retrieval call binding the contract method 0x74e73fd3.
//
// Solidity: function dummy22() pure returns(uint256)
func (_Contract *ContractCaller) Dummy22(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy22")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy22 is a free data retrieval call binding the contract method 0x74e73fd3.
//
// Solidity: function dummy22() pure returns(uint256)
func (_Contract *ContractSession) Dummy22() (*big.Int, error) {
	return _Contract.Contract.Dummy22(&_Contract.CallOpts)
}

// Dummy22 is a free data retrieval call binding the contract method 0x74e73fd3.
//
// Solidity: function dummy22() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy22() (*big.Int, error) {
	return _Contract.Contract.Dummy22(&_Contract.CallOpts)
}

// Dummy23 is a free data retrieval call binding the contract method 0x6abceacd.
//
// Solidity: function dummy23() pure returns(uint256)
func (_Contract *ContractCaller) Dummy23(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy23")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy23 is a free data retrieval call binding the contract method 0x6abceacd.
//
// Solidity: function dummy23() pure returns(uint256)
func (_Contract *ContractSession) Dummy23() (*big.Int, error) {
	return _Contract.Contract.Dummy23(&_Contract.CallOpts)
}

// Dummy23 is a free data retrieval call binding the contract method 0x6abceacd.
//
// Solidity: function dummy23() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy23() (*big.Int, error) {
	return _Contract.Contract.Dummy23(&_Contract.CallOpts)
}

// Dummy24 is a free data retrieval call binding the contract method 0xb2bb360e.
//
// Solidity: function dummy24() pure returns(uint256)
func (_Contract *ContractCaller) Dummy24(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy24")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy24 is a free data retrieval call binding the contract method 0xb2bb360e.
//
// Solidity: function dummy24() pure returns(uint256)
func (_Contract *ContractSession) Dummy24() (*big.Int, error) {
	return _Contract.Contract.Dummy24(&_Contract.CallOpts)
}

// Dummy24 is a free data retrieval call binding the contract method 0xb2bb360e.
//
// Solidity: function dummy24() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy24() (*big.Int, error) {
	return _Contract.Contract.Dummy24(&_Contract.CallOpts)
}

// Dummy25 is a free data retrieval call binding the contract method 0xd7419469.
//
// Solidity: function dummy25() pure returns(uint256)
func (_Contract *ContractCaller) Dummy25(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy25")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy25 is a free data retrieval call binding the contract method 0xd7419469.
//
// Solidity: function dummy25() pure returns(uint256)
func (_Contract *ContractSession) Dummy25() (*big.Int, error) {
	return _Contract.Contract.Dummy25(&_Contract.CallOpts)
}

// Dummy25 is a free data retrieval call binding the contract method 0xd7419469.
//
// Solidity: function dummy25() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy25() (*big.Int, error) {
	return _Contract.Contract.Dummy25(&_Contract.CallOpts)
}

// Dummy26 is a free data retrieval call binding the contract method 0x12901b42.
//
// Solidity: function dummy26() pure returns(uint256)
func (_Contract *ContractCaller) Dummy26(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy26")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy26 is a free data retrieval call binding the contract method 0x12901b42.
//
// Solidity: function dummy26() pure returns(uint256)
func (_Contract *ContractSession) Dummy26() (*big.Int, error) {
	return _Contract.Contract.Dummy26(&_Contract.CallOpts)
}

// Dummy26 is a free data retrieval call binding the contract method 0x12901b42.
//
// Solidity: function dummy26() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy26() (*big.Int, error) {
	return _Contract.Contract.Dummy26(&_Contract.CallOpts)
}

// Dummy27 is a free data retrieval call binding the contract method 0x0cb7a9e7.
//
// Solidity: function dummy27() pure returns(uint256)
func (_Contract *ContractCaller) Dummy27(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy27")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy27 is a free data retrieval call binding the contract method 0x0cb7a9e7.
//
// Solidity: function dummy27() pure returns(uint256)
func (_Contract *ContractSession) Dummy27() (*big.Int, error) {
	return _Contract.Contract.Dummy27(&_Contract.CallOpts)
}

// Dummy27 is a free data retrieval call binding the contract method 0x0cb7a9e7.
//
// Solidity: function dummy27() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy27() (*big.Int, error) {
	return _Contract.Contract.Dummy27(&_Contract.CallOpts)
}

// Dummy28 is a free data retrieval call binding the contract method 0x418b1816.
//
// Solidity: function dummy28() pure returns(uint256)
func (_Contract *ContractCaller) Dummy28(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy28")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy28 is a free data retrieval call binding the contract method 0x418b1816.
//
// Solidity: function dummy28() pure returns(uint256)
func (_Contract *ContractSession) Dummy28() (*big.Int, error) {
	return _Contract.Contract.Dummy28(&_Contract.CallOpts)
}

// Dummy28 is a free data retrieval call binding the contract method 0x418b1816.
//
// Solidity: function dummy28() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy28() (*big.Int, error) {
	return _Contract.Contract.Dummy28(&_Contract.CallOpts)
}

// Dummy29 is a free data retrieval call binding the contract method 0x13ebb5ec.
//
// Solidity: function dummy29() pure returns(uint256)
func (_Contract *ContractCaller) Dummy29(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy29")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy29 is a free data retrieval call binding the contract method 0x13ebb5ec.
//
// Solidity: function dummy29() pure returns(uint256)
func (_Contract *ContractSession) Dummy29() (*big.Int, error) {
	return _Contract.Contract.Dummy29(&_Contract.CallOpts)
}

// Dummy29 is a free data retrieval call binding the contract method 0x13ebb5ec.
//
// Solidity: function dummy29() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy29() (*big.Int, error) {
	return _Contract.Contract.Dummy29(&_Contract.CallOpts)
}

// Dummy3 is a free data retrieval call binding the contract method 0x1a97f18e.
//
// Solidity: function dummy3() pure returns(uint256)
func (_Contract *ContractCaller) Dummy3(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy3")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy3 is a free data retrieval call binding the contract method 0x1a97f18e.
//
// Solidity: function dummy3() pure returns(uint256)
func (_Contract *ContractSession) Dummy3() (*big.Int, error) {
	return _Contract.Contract.Dummy3(&_Contract.CallOpts)
}

// Dummy3 is a free data retrieval call binding the contract method 0x1a97f18e.
//
// Solidity: function dummy3() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy3() (*big.Int, error) {
	return _Contract.Contract.Dummy3(&_Contract.CallOpts)
}

// Dummy30 is a free data retrieval call binding the contract method 0xcfd66863.
//
// Solidity: function dummy30() pure returns(uint256)
func (_Contract *ContractCaller) Dummy30(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy30")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy30 is a free data retrieval call binding the contract method 0xcfd66863.
//
// Solidity: function dummy30() pure returns(uint256)
func (_Contract *ContractSession) Dummy30() (*big.Int, error) {
	return _Contract.Contract.Dummy30(&_Contract.CallOpts)
}

// Dummy30 is a free data retrieval call binding the contract method 0xcfd66863.
//
// Solidity: function dummy30() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy30() (*big.Int, error) {
	return _Contract.Contract.Dummy30(&_Contract.CallOpts)
}

// Dummy31 is a free data retrieval call binding the contract method 0x4e1dbb82.
//
// Solidity: function dummy31() pure returns(uint256)
func (_Contract *ContractCaller) Dummy31(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy31")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy31 is a free data retrieval call binding the contract method 0x4e1dbb82.
//
// Solidity: function dummy31() pure returns(uint256)
func (_Contract *ContractSession) Dummy31() (*big.Int, error) {
	return _Contract.Contract.Dummy31(&_Contract.CallOpts)
}

// Dummy31 is a free data retrieval call binding the contract method 0x4e1dbb82.
//
// Solidity: function dummy31() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy31() (*big.Int, error) {
	return _Contract.Contract.Dummy31(&_Contract.CallOpts)
}

// Dummy32 is a free data retrieval call binding the contract method 0xf8716f14.
//
// Solidity: function dummy32() pure returns(uint256)
func (_Contract *ContractCaller) Dummy32(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy32")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy32 is a free data retrieval call binding the contract method 0xf8716f14.
//
// Solidity: function dummy32() pure returns(uint256)
func (_Contract *ContractSession) Dummy32() (*big.Int, error) {
	return _Contract.Contract.Dummy32(&_Contract.CallOpts)
}

// Dummy32 is a free data retrieval call binding the contract method 0xf8716f14.
//
// Solidity: function dummy32() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy32() (*big.Int, error) {
	return _Contract.Contract.Dummy32(&_Contract.CallOpts)
}

// Dummy33 is a free data retrieval call binding the contract method 0x4b3c7f5f.
//
// Solidity: function dummy33() pure returns(uint256)
func (_Contract *ContractCaller) Dummy33(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy33")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy33 is a free data retrieval call binding the contract method 0x4b3c7f5f.
//
// Solidity: function dummy33() pure returns(uint256)
func (_Contract *ContractSession) Dummy33() (*big.Int, error) {
	return _Contract.Contract.Dummy33(&_Contract.CallOpts)
}

// Dummy33 is a free data retrieval call binding the contract method 0x4b3c7f5f.
//
// Solidity: function dummy33() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy33() (*big.Int, error) {
	return _Contract.Contract.Dummy33(&_Contract.CallOpts)
}

// Dummy34 is a free data retrieval call binding the contract method 0xe8c927b3.
//
// Solidity: function dummy34() pure returns(uint256)
func (_Contract *ContractCaller) Dummy34(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy34")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy34 is a free data retrieval call binding the contract method 0xe8c927b3.
//
// Solidity: function dummy34() pure returns(uint256)
func (_Contract *ContractSession) Dummy34() (*big.Int, error) {
	return _Contract.Contract.Dummy34(&_Contract.CallOpts)
}

// Dummy34 is a free data retrieval call binding the contract method 0xe8c927b3.
//
// Solidity: function dummy34() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy34() (*big.Int, error) {
	return _Contract.Contract.Dummy34(&_Contract.CallOpts)
}

// Dummy35 is a free data retrieval call binding the contract method 0x43a6b92d.
//
// Solidity: function dummy35() pure returns(uint256)
func (_Contract *ContractCaller) Dummy35(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy35")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy35 is a free data retrieval call binding the contract method 0x43a6b92d.
//
// Solidity: function dummy35() pure returns(uint256)
func (_Contract *ContractSession) Dummy35() (*big.Int, error) {
	return _Contract.Contract.Dummy35(&_Contract.CallOpts)
}

// Dummy35 is a free data retrieval call binding the contract method 0x43a6b92d.
//
// Solidity: function dummy35() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy35() (*big.Int, error) {
	return _Contract.Contract.Dummy35(&_Contract.CallOpts)
}

// Dummy36 is a free data retrieval call binding the contract method 0xc958d4bf.
//
// Solidity: function dummy36() pure returns(uint256)
func (_Contract *ContractCaller) Dummy36(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy36")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy36 is a free data retrieval call binding the contract method 0xc958d4bf.
//
// Solidity: function dummy36() pure returns(uint256)
func (_Contract *ContractSession) Dummy36() (*big.Int, error) {
	return _Contract.Contract.Dummy36(&_Contract.CallOpts)
}

// Dummy36 is a free data retrieval call binding the contract method 0xc958d4bf.
//
// Solidity: function dummy36() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy36() (*big.Int, error) {
	return _Contract.Contract.Dummy36(&_Contract.CallOpts)
}

// Dummy37 is a free data retrieval call binding the contract method 0x639ec53a.
//
// Solidity: function dummy37() pure returns(uint256)
func (_Contract *ContractCaller) Dummy37(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy37")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy37 is a free data retrieval call binding the contract method 0x639ec53a.
//
// Solidity: function dummy37() pure returns(uint256)
func (_Contract *ContractSession) Dummy37() (*big.Int, error) {
	return _Contract.Contract.Dummy37(&_Contract.CallOpts)
}

// Dummy37 is a free data retrieval call binding the contract method 0x639ec53a.
//
// Solidity: function dummy37() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy37() (*big.Int, error) {
	return _Contract.Contract.Dummy37(&_Contract.CallOpts)
}

// Dummy38 is a free data retrieval call binding the contract method 0x4f5e5557.
//
// Solidity: function dummy38() pure returns(uint256)
func (_Contract *ContractCaller) Dummy38(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy38")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy38 is a free data retrieval call binding the contract method 0x4f5e5557.
//
// Solidity: function dummy38() pure returns(uint256)
func (_Contract *ContractSession) Dummy38() (*big.Int, error) {
	return _Contract.Contract.Dummy38(&_Contract.CallOpts)
}

// Dummy38 is a free data retrieval call binding the contract method 0x4f5e5557.
//
// Solidity: function dummy38() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy38() (*big.Int, error) {
	return _Contract.Contract.Dummy38(&_Contract.CallOpts)
}

// Dummy39 is a free data retrieval call binding the contract method 0xdc1d8a9b.
//
// Solidity: function dummy39() pure returns(uint256)
func (_Contract *ContractCaller) Dummy39(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy39")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy39 is a free data retrieval call binding the contract method 0xdc1d8a9b.
//
// Solidity: function dummy39() pure returns(uint256)
func (_Contract *ContractSession) Dummy39() (*big.Int, error) {
	return _Contract.Contract.Dummy39(&_Contract.CallOpts)
}

// Dummy39 is a free data retrieval call binding the contract method 0xdc1d8a9b.
//
// Solidity: function dummy39() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy39() (*big.Int, error) {
	return _Contract.Contract.Dummy39(&_Contract.CallOpts)
}

// Dummy4 is a free data retrieval call binding the contract method 0x3b6be459.
//
// Solidity: function dummy4() pure returns(uint256)
func (_Contract *ContractCaller) Dummy4(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy4")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy4 is a free data retrieval call binding the contract method 0x3b6be459.
//
// Solidity: function dummy4() pure returns(uint256)
func (_Contract *ContractSession) Dummy4() (*big.Int, error) {
	return _Contract.Contract.Dummy4(&_Contract.CallOpts)
}

// Dummy4 is a free data retrieval call binding the contract method 0x3b6be459.
//
// Solidity: function dummy4() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy4() (*big.Int, error) {
	return _Contract.Contract.Dummy4(&_Contract.CallOpts)
}

// Dummy40 is a free data retrieval call binding the contract method 0xaaa7af70.
//
// Solidity: function dummy40() pure returns(uint256)
func (_Contract *ContractCaller) Dummy40(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy40")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy40 is a free data retrieval call binding the contract method 0xaaa7af70.
//
// Solidity: function dummy40() pure returns(uint256)
func (_Contract *ContractSession) Dummy40() (*big.Int, error) {
	return _Contract.Contract.Dummy40(&_Contract.CallOpts)
}

// Dummy40 is a free data retrieval call binding the contract method 0xaaa7af70.
//
// Solidity: function dummy40() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy40() (*big.Int, error) {
	return _Contract.Contract.Dummy40(&_Contract.CallOpts)
}

// Dummy41 is a free data retrieval call binding the contract method 0x9df61a25.
//
// Solidity: function dummy41() pure returns(uint256)
func (_Contract *ContractCaller) Dummy41(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy41")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy41 is a free data retrieval call binding the contract method 0x9df61a25.
//
// Solidity: function dummy41() pure returns(uint256)
func (_Contract *ContractSession) Dummy41() (*big.Int, error) {
	return _Contract.Contract.Dummy41(&_Contract.CallOpts)
}

// Dummy41 is a free data retrieval call binding the contract method 0x9df61a25.
//
// Solidity: function dummy41() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy41() (*big.Int, error) {
	return _Contract.Contract.Dummy41(&_Contract.CallOpts)
}

// Dummy42 is a free data retrieval call binding the contract method 0x5af92c05.
//
// Solidity: function dummy42() pure returns(uint256)
func (_Contract *ContractCaller) Dummy42(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy42")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy42 is a free data retrieval call binding the contract method 0x5af92c05.
//
// Solidity: function dummy42() pure returns(uint256)
func (_Contract *ContractSession) Dummy42() (*big.Int, error) {
	return _Contract.Contract.Dummy42(&_Contract.CallOpts)
}

// Dummy42 is a free data retrieval call binding the contract method 0x5af92c05.
//
// Solidity: function dummy42() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy42() (*big.Int, error) {
	return _Contract.Contract.Dummy42(&_Contract.CallOpts)
}

// Dummy43 is a free data retrieval call binding the contract method 0x7c72ed0d.
//
// Solidity: function dummy43() pure returns(uint256)
func (_Contract *ContractCaller) Dummy43(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy43")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy43 is a free data retrieval call binding the contract method 0x7c72ed0d.
//
// Solidity: function dummy43() pure returns(uint256)
func (_Contract *ContractSession) Dummy43() (*big.Int, error) {
	return _Contract.Contract.Dummy43(&_Contract.CallOpts)
}

// Dummy43 is a free data retrieval call binding the contract method 0x7c72ed0d.
//
// Solidity: function dummy43() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy43() (*big.Int, error) {
	return _Contract.Contract.Dummy43(&_Contract.CallOpts)
}

// Dummy44 is a free data retrieval call binding the contract method 0x239af2a5.
//
// Solidity: function dummy44() pure returns(uint256)
func (_Contract *ContractCaller) Dummy44(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy44")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy44 is a free data retrieval call binding the contract method 0x239af2a5.
//
// Solidity: function dummy44() pure returns(uint256)
func (_Contract *ContractSession) Dummy44() (*big.Int, error) {
	return _Contract.Contract.Dummy44(&_Contract.CallOpts)
}

// Dummy44 is a free data retrieval call binding the contract method 0x239af2a5.
//
// Solidity: function dummy44() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy44() (*big.Int, error) {
	return _Contract.Contract.Dummy44(&_Contract.CallOpts)
}

// Dummy45 is a free data retrieval call binding the contract method 0x7f34d94b.
//
// Solidity: function dummy45() pure returns(uint256)
func (_Contract *ContractCaller) Dummy45(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy45")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy45 is a free data retrieval call binding the contract method 0x7f34d94b.
//
// Solidity: function dummy45() pure returns(uint256)
func (_Contract *ContractSession) Dummy45() (*big.Int, error) {
	return _Contract.Contract.Dummy45(&_Contract.CallOpts)
}

// Dummy45 is a free data retrieval call binding the contract method 0x7f34d94b.
//
// Solidity: function dummy45() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy45() (*big.Int, error) {
	return _Contract.Contract.Dummy45(&_Contract.CallOpts)
}

// Dummy46 is a free data retrieval call binding the contract method 0x21ecd7a3.
//
// Solidity: function dummy46() pure returns(uint256)
func (_Contract *ContractCaller) Dummy46(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy46")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy46 is a free data retrieval call binding the contract method 0x21ecd7a3.
//
// Solidity: function dummy46() pure returns(uint256)
func (_Contract *ContractSession) Dummy46() (*big.Int, error) {
	return _Contract.Contract.Dummy46(&_Contract.CallOpts)
}

// Dummy46 is a free data retrieval call binding the contract method 0x21ecd7a3.
//
// Solidity: function dummy46() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy46() (*big.Int, error) {
	return _Contract.Contract.Dummy46(&_Contract.CallOpts)
}

// Dummy47 is a free data retrieval call binding the contract method 0x0717b161.
//
// Solidity: function dummy47() pure returns(uint256)
func (_Contract *ContractCaller) Dummy47(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy47")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy47 is a free data retrieval call binding the contract method 0x0717b161.
//
// Solidity: function dummy47() pure returns(uint256)
func (_Contract *ContractSession) Dummy47() (*big.Int, error) {
	return _Contract.Contract.Dummy47(&_Contract.CallOpts)
}

// Dummy47 is a free data retrieval call binding the contract method 0x0717b161.
//
// Solidity: function dummy47() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy47() (*big.Int, error) {
	return _Contract.Contract.Dummy47(&_Contract.CallOpts)
}

// Dummy48 is a free data retrieval call binding the contract method 0x8619d607.
//
// Solidity: function dummy48() pure returns(uint256)
func (_Contract *ContractCaller) Dummy48(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy48")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy48 is a free data retrieval call binding the contract method 0x8619d607.
//
// Solidity: function dummy48() pure returns(uint256)
func (_Contract *ContractSession) Dummy48() (*big.Int, error) {
	return _Contract.Contract.Dummy48(&_Contract.CallOpts)
}

// Dummy48 is a free data retrieval call binding the contract method 0x8619d607.
//
// Solidity: function dummy48() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy48() (*big.Int, error) {
	return _Contract.Contract.Dummy48(&_Contract.CallOpts)
}

// Dummy49 is a free data retrieval call binding the contract method 0xd101dcd0.
//
// Solidity: function dummy49() pure returns(uint256)
func (_Contract *ContractCaller) Dummy49(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy49")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy49 is a free data retrieval call binding the contract method 0xd101dcd0.
//
// Solidity: function dummy49() pure returns(uint256)
func (_Contract *ContractSession) Dummy49() (*big.Int, error) {
	return _Contract.Contract.Dummy49(&_Contract.CallOpts)
}

// Dummy49 is a free data retrieval call binding the contract method 0xd101dcd0.
//
// Solidity: function dummy49() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy49() (*big.Int, error) {
	return _Contract.Contract.Dummy49(&_Contract.CallOpts)
}

// Dummy5 is a free data retrieval call binding the contract method 0x1215a3ab.
//
// Solidity: function dummy5() pure returns(uint256)
func (_Contract *ContractCaller) Dummy5(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy5")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy5 is a free data retrieval call binding the contract method 0x1215a3ab.
//
// Solidity: function dummy5() pure returns(uint256)
func (_Contract *ContractSession) Dummy5() (*big.Int, error) {
	return _Contract.Contract.Dummy5(&_Contract.CallOpts)
}

// Dummy5 is a free data retrieval call binding the contract method 0x1215a3ab.
//
// Solidity: function dummy5() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy5() (*big.Int, error) {
	return _Contract.Contract.Dummy5(&_Contract.CallOpts)
}

// Dummy50 is a free data retrieval call binding the contract method 0x42937dbd.
//
// Solidity: function dummy50() pure returns(uint256)
func (_Contract *ContractCaller) Dummy50(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy50")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy50 is a free data retrieval call binding the contract method 0x42937dbd.
//
// Solidity: function dummy50() pure returns(uint256)
func (_Contract *ContractSession) Dummy50() (*big.Int, error) {
	return _Contract.Contract.Dummy50(&_Contract.CallOpts)
}

// Dummy50 is a free data retrieval call binding the contract method 0x42937dbd.
//
// Solidity: function dummy50() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy50() (*big.Int, error) {
	return _Contract.Contract.Dummy50(&_Contract.CallOpts)
}

// Dummy51 is a free data retrieval call binding the contract method 0x16a3045b.
//
// Solidity: function dummy51() pure returns(uint256)
func (_Contract *ContractCaller) Dummy51(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy51")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy51 is a free data retrieval call binding the contract method 0x16a3045b.
//
// Solidity: function dummy51() pure returns(uint256)
func (_Contract *ContractSession) Dummy51() (*big.Int, error) {
	return _Contract.Contract.Dummy51(&_Contract.CallOpts)
}

// Dummy51 is a free data retrieval call binding the contract method 0x16a3045b.
//
// Solidity: function dummy51() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy51() (*big.Int, error) {
	return _Contract.Contract.Dummy51(&_Contract.CallOpts)
}

// Dummy52 is a free data retrieval call binding the contract method 0x3125f37a.
//
// Solidity: function dummy52() pure returns(uint256)
func (_Contract *ContractCaller) Dummy52(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy52")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy52 is a free data retrieval call binding the contract method 0x3125f37a.
//
// Solidity: function dummy52() pure returns(uint256)
func (_Contract *ContractSession) Dummy52() (*big.Int, error) {
	return _Contract.Contract.Dummy52(&_Contract.CallOpts)
}

// Dummy52 is a free data retrieval call binding the contract method 0x3125f37a.
//
// Solidity: function dummy52() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy52() (*big.Int, error) {
	return _Contract.Contract.Dummy52(&_Contract.CallOpts)
}

// Dummy53 is a free data retrieval call binding the contract method 0x1fd298ec.
//
// Solidity: function dummy53() pure returns(uint256)
func (_Contract *ContractCaller) Dummy53(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy53")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy53 is a free data retrieval call binding the contract method 0x1fd298ec.
//
// Solidity: function dummy53() pure returns(uint256)
func (_Contract *ContractSession) Dummy53() (*big.Int, error) {
	return _Contract.Contract.Dummy53(&_Contract.CallOpts)
}

// Dummy53 is a free data retrieval call binding the contract method 0x1fd298ec.
//
// Solidity: function dummy53() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy53() (*big.Int, error) {
	return _Contract.Contract.Dummy53(&_Contract.CallOpts)
}

// Dummy54 is a free data retrieval call binding the contract method 0x792c7f3e.
//
// Solidity: function dummy54() pure returns(uint256)
func (_Contract *ContractCaller) Dummy54(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy54")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy54 is a free data retrieval call binding the contract method 0x792c7f3e.
//
// Solidity: function dummy54() pure returns(uint256)
func (_Contract *ContractSession) Dummy54() (*big.Int, error) {
	return _Contract.Contract.Dummy54(&_Contract.CallOpts)
}

// Dummy54 is a free data retrieval call binding the contract method 0x792c7f3e.
//
// Solidity: function dummy54() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy54() (*big.Int, error) {
	return _Contract.Contract.Dummy54(&_Contract.CallOpts)
}

// Dummy55 is a free data retrieval call binding the contract method 0x7dffdc32.
//
// Solidity: function dummy55() pure returns(uint256)
func (_Contract *ContractCaller) Dummy55(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy55")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy55 is a free data retrieval call binding the contract method 0x7dffdc32.
//
// Solidity: function dummy55() pure returns(uint256)
func (_Contract *ContractSession) Dummy55() (*big.Int, error) {
	return _Contract.Contract.Dummy55(&_Contract.CallOpts)
}

// Dummy55 is a free data retrieval call binding the contract method 0x7dffdc32.
//
// Solidity: function dummy55() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy55() (*big.Int, error) {
	return _Contract.Contract.Dummy55(&_Contract.CallOpts)
}

// Dummy56 is a free data retrieval call binding the contract method 0xa891d4d4.
//
// Solidity: function dummy56() pure returns(uint256)
func (_Contract *ContractCaller) Dummy56(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy56")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy56 is a free data retrieval call binding the contract method 0xa891d4d4.
//
// Solidity: function dummy56() pure returns(uint256)
func (_Contract *ContractSession) Dummy56() (*big.Int, error) {
	return _Contract.Contract.Dummy56(&_Contract.CallOpts)
}

// Dummy56 is a free data retrieval call binding the contract method 0xa891d4d4.
//
// Solidity: function dummy56() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy56() (*big.Int, error) {
	return _Contract.Contract.Dummy56(&_Contract.CallOpts)
}

// Dummy57 is a free data retrieval call binding the contract method 0x74f83d02.
//
// Solidity: function dummy57() pure returns(uint256)
func (_Contract *ContractCaller) Dummy57(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy57")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy57 is a free data retrieval call binding the contract method 0x74f83d02.
//
// Solidity: function dummy57() pure returns(uint256)
func (_Contract *ContractSession) Dummy57() (*big.Int, error) {
	return _Contract.Contract.Dummy57(&_Contract.CallOpts)
}

// Dummy57 is a free data retrieval call binding the contract method 0x74f83d02.
//
// Solidity: function dummy57() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy57() (*big.Int, error) {
	return _Contract.Contract.Dummy57(&_Contract.CallOpts)
}

// Dummy58 is a free data retrieval call binding the contract method 0x54c27920.
//
// Solidity: function dummy58() pure returns(uint256)
func (_Contract *ContractCaller) Dummy58(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy58")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy58 is a free data retrieval call binding the contract method 0x54c27920.
//
// Solidity: function dummy58() pure returns(uint256)
func (_Contract *ContractSession) Dummy58() (*big.Int, error) {
	return _Contract.Contract.Dummy58(&_Contract.CallOpts)
}

// Dummy58 is a free data retrieval call binding the contract method 0x54c27920.
//
// Solidity: function dummy58() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy58() (*big.Int, error) {
	return _Contract.Contract.Dummy58(&_Contract.CallOpts)
}

// Dummy59 is a free data retrieval call binding the contract method 0x77c0209e.
//
// Solidity: function dummy59() pure returns(uint256)
func (_Contract *ContractCaller) Dummy59(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy59")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy59 is a free data retrieval call binding the contract method 0x77c0209e.
//
// Solidity: function dummy59() pure returns(uint256)
func (_Contract *ContractSession) Dummy59() (*big.Int, error) {
	return _Contract.Contract.Dummy59(&_Contract.CallOpts)
}

// Dummy59 is a free data retrieval call binding the contract method 0x77c0209e.
//
// Solidity: function dummy59() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy59() (*big.Int, error) {
	return _Contract.Contract.Dummy59(&_Contract.CallOpts)
}

// Dummy6 is a free data retrieval call binding the contract method 0x4128a85d.
//
// Solidity: function dummy6() pure returns(uint256)
func (_Contract *ContractCaller) Dummy6(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy6")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy6 is a free data retrieval call binding the contract method 0x4128a85d.
//
// Solidity: function dummy6() pure returns(uint256)
func (_Contract *ContractSession) Dummy6() (*big.Int, error) {
	return _Contract.Contract.Dummy6(&_Contract.CallOpts)
}

// Dummy6 is a free data retrieval call binding the contract method 0x4128a85d.
//
// Solidity: function dummy6() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy6() (*big.Int, error) {
	return _Contract.Contract.Dummy6(&_Contract.CallOpts)
}

// Dummy7 is a free data retrieval call binding the contract method 0xf26c779b.
//
// Solidity: function dummy7() pure returns(uint256)
func (_Contract *ContractCaller) Dummy7(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy7")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy7 is a free data retrieval call binding the contract method 0xf26c779b.
//
// Solidity: function dummy7() pure returns(uint256)
func (_Contract *ContractSession) Dummy7() (*big.Int, error) {
	return _Contract.Contract.Dummy7(&_Contract.CallOpts)
}

// Dummy7 is a free data retrieval call binding the contract method 0xf26c779b.
//
// Solidity: function dummy7() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy7() (*big.Int, error) {
	return _Contract.Contract.Dummy7(&_Contract.CallOpts)
}

// Dummy8 is a free data retrieval call binding the contract method 0x19cf6a91.
//
// Solidity: function dummy8() pure returns(uint256)
func (_Contract *ContractCaller) Dummy8(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy8")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy8 is a free data retrieval call binding the contract method 0x19cf6a91.
//
// Solidity: function dummy8() pure returns(uint256)
func (_Contract *ContractSession) Dummy8() (*big.Int, error) {
	return _Contract.Contract.Dummy8(&_Contract.CallOpts)
}

// Dummy8 is a free data retrieval call binding the contract method 0x19cf6a91.
//
// Solidity: function dummy8() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy8() (*big.Int, error) {
	return _Contract.Contract.Dummy8(&_Contract.CallOpts)
}

// Dummy9 is a free data retrieval call binding the contract method 0x58b6a9bd.
//
// Solidity: function dummy9() pure returns(uint256)
func (_Contract *ContractCaller) Dummy9(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dummy9")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dummy9 is a free data retrieval call binding the contract method 0x58b6a9bd.
//
// Solidity: function dummy9() pure returns(uint256)
func (_Contract *ContractSession) Dummy9() (*big.Int, error) {
	return _Contract.Contract.Dummy9(&_Contract.CallOpts)
}

// Dummy9 is a free data retrieval call binding the contract method 0x58b6a9bd.
//
// Solidity: function dummy9() pure returns(uint256)
func (_Contract *ContractCallerSession) Dummy9() (*big.Int, error) {
	return _Contract.Contract.Dummy9(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(uint256)
func (_Contract *ContractCaller) Salt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(uint256)
func (_Contract *ContractSession) Salt() (*big.Int, error) {
	return _Contract.Contract.Salt(&_Contract.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(uint256)
func (_Contract *ContractCallerSession) Salt() (*big.Int, error) {
	return _Contract.Contract.Salt(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contract *ContractSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, spender, value)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0xd3c1c838.
//
// Solidity: function batchTransfer(address[] recipients) returns()
func (_Contract *ContractTransactor) BatchTransfer(opts *bind.TransactOpts, recipients []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "batchTransfer", recipients)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0xd3c1c838.
//
// Solidity: function batchTransfer(address[] recipients) returns()
func (_Contract *ContractSession) BatchTransfer(recipients []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.BatchTransfer(&_Contract.TransactOpts, recipients)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0xd3c1c838.
//
// Solidity: function batchTransfer(address[] recipients) returns()
func (_Contract *ContractTransactorSession) BatchTransfer(recipients []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.BatchTransfer(&_Contract.TransactOpts, recipients)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contract *ContractSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0xbb9bfe06.
//
// Solidity: function transferFrom1(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom1(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom1", from, to, value)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0xbb9bfe06.
//
// Solidity: function transferFrom1(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom1(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom1(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom1 is a paid mutator transaction binding the contract method 0xbb9bfe06.
//
// Solidity: function transferFrom1(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom1(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom1(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom10 is a paid mutator transaction binding the contract method 0xb1802b9a.
//
// Solidity: function transferFrom10(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom10(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom10", from, to, value)
}

// TransferFrom10 is a paid mutator transaction binding the contract method 0xb1802b9a.
//
// Solidity: function transferFrom10(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom10(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom10(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom10 is a paid mutator transaction binding the contract method 0xb1802b9a.
//
// Solidity: function transferFrom10(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom10(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom10(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom11 is a paid mutator transaction binding the contract method 0xc2be97e3.
//
// Solidity: function transferFrom11(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom11(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom11", from, to, value)
}

// TransferFrom11 is a paid mutator transaction binding the contract method 0xc2be97e3.
//
// Solidity: function transferFrom11(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom11(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom11(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom11 is a paid mutator transaction binding the contract method 0xc2be97e3.
//
// Solidity: function transferFrom11(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom11(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom11(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom12 is a paid mutator transaction binding the contract method 0x44050a28.
//
// Solidity: function transferFrom12(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom12(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom12", from, to, value)
}

// TransferFrom12 is a paid mutator transaction binding the contract method 0x44050a28.
//
// Solidity: function transferFrom12(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom12(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom12(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom12 is a paid mutator transaction binding the contract method 0x44050a28.
//
// Solidity: function transferFrom12(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom12(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom12(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom13 is a paid mutator transaction binding the contract method 0xacc5aee9.
//
// Solidity: function transferFrom13(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom13(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom13", from, to, value)
}

// TransferFrom13 is a paid mutator transaction binding the contract method 0xacc5aee9.
//
// Solidity: function transferFrom13(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom13(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom13(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom13 is a paid mutator transaction binding the contract method 0xacc5aee9.
//
// Solidity: function transferFrom13(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom13(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom13(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom14 is a paid mutator transaction binding the contract method 0x1bbffe6f.
//
// Solidity: function transferFrom14(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom14(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom14", from, to, value)
}

// TransferFrom14 is a paid mutator transaction binding the contract method 0x1bbffe6f.
//
// Solidity: function transferFrom14(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom14(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom14(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom14 is a paid mutator transaction binding the contract method 0x1bbffe6f.
//
// Solidity: function transferFrom14(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom14(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom14(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom15 is a paid mutator transaction binding the contract method 0x8789ca67.
//
// Solidity: function transferFrom15(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom15(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom15", from, to, value)
}

// TransferFrom15 is a paid mutator transaction binding the contract method 0x8789ca67.
//
// Solidity: function transferFrom15(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom15(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom15(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom15 is a paid mutator transaction binding the contract method 0x8789ca67.
//
// Solidity: function transferFrom15(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom15(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom15(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom16 is a paid mutator transaction binding the contract method 0x39e0bd12.
//
// Solidity: function transferFrom16(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom16(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom16", from, to, value)
}

// TransferFrom16 is a paid mutator transaction binding the contract method 0x39e0bd12.
//
// Solidity: function transferFrom16(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom16(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom16(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom16 is a paid mutator transaction binding the contract method 0x39e0bd12.
//
// Solidity: function transferFrom16(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom16(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom16(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom17 is a paid mutator transaction binding the contract method 0x291c3bd7.
//
// Solidity: function transferFrom17(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom17(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom17", from, to, value)
}

// TransferFrom17 is a paid mutator transaction binding the contract method 0x291c3bd7.
//
// Solidity: function transferFrom17(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom17(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom17(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom17 is a paid mutator transaction binding the contract method 0x291c3bd7.
//
// Solidity: function transferFrom17(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom17(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom17(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom18 is a paid mutator transaction binding the contract method 0x657b6ef7.
//
// Solidity: function transferFrom18(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom18(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom18", from, to, value)
}

// TransferFrom18 is a paid mutator transaction binding the contract method 0x657b6ef7.
//
// Solidity: function transferFrom18(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom18(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom18(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom18 is a paid mutator transaction binding the contract method 0x657b6ef7.
//
// Solidity: function transferFrom18(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom18(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom18(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom19 is a paid mutator transaction binding the contract method 0xeb4329c8.
//
// Solidity: function transferFrom19(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom19(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom19", from, to, value)
}

// TransferFrom19 is a paid mutator transaction binding the contract method 0xeb4329c8.
//
// Solidity: function transferFrom19(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom19(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom19(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom19 is a paid mutator transaction binding the contract method 0xeb4329c8.
//
// Solidity: function transferFrom19(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom19(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom19(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x6c12ed28.
//
// Solidity: function transferFrom2(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom2(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom2", from, to, value)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x6c12ed28.
//
// Solidity: function transferFrom2(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom2(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom2(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom2 is a paid mutator transaction binding the contract method 0x6c12ed28.
//
// Solidity: function transferFrom2(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom2(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom2(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom3 is a paid mutator transaction binding the contract method 0x1b17c65c.
//
// Solidity: function transferFrom3(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom3(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom3", from, to, value)
}

// TransferFrom3 is a paid mutator transaction binding the contract method 0x1b17c65c.
//
// Solidity: function transferFrom3(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom3(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom3(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom3 is a paid mutator transaction binding the contract method 0x1b17c65c.
//
// Solidity: function transferFrom3(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom3(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom3(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom4 is a paid mutator transaction binding the contract method 0x3a131990.
//
// Solidity: function transferFrom4(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom4(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom4", from, to, value)
}

// TransferFrom4 is a paid mutator transaction binding the contract method 0x3a131990.
//
// Solidity: function transferFrom4(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom4(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom4(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom4 is a paid mutator transaction binding the contract method 0x3a131990.
//
// Solidity: function transferFrom4(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom4(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom4(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom5 is a paid mutator transaction binding the contract method 0x0460faf6.
//
// Solidity: function transferFrom5(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom5(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom5", from, to, value)
}

// TransferFrom5 is a paid mutator transaction binding the contract method 0x0460faf6.
//
// Solidity: function transferFrom5(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom5(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom5(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom5 is a paid mutator transaction binding the contract method 0x0460faf6.
//
// Solidity: function transferFrom5(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom5(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom5(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom6 is a paid mutator transaction binding the contract method 0x7c66673e.
//
// Solidity: function transferFrom6(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom6(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom6", from, to, value)
}

// TransferFrom6 is a paid mutator transaction binding the contract method 0x7c66673e.
//
// Solidity: function transferFrom6(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom6(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom6(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom6 is a paid mutator transaction binding the contract method 0x7c66673e.
//
// Solidity: function transferFrom6(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom6(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom6(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom7 is a paid mutator transaction binding the contract method 0xfaf35ced.
//
// Solidity: function transferFrom7(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom7(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom7", from, to, value)
}

// TransferFrom7 is a paid mutator transaction binding the contract method 0xfaf35ced.
//
// Solidity: function transferFrom7(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom7(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom7(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom7 is a paid mutator transaction binding the contract method 0xfaf35ced.
//
// Solidity: function transferFrom7(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom7(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom7(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom8 is a paid mutator transaction binding the contract method 0x552a1b56.
//
// Solidity: function transferFrom8(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom8(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom8", from, to, value)
}

// TransferFrom8 is a paid mutator transaction binding the contract method 0x552a1b56.
//
// Solidity: function transferFrom8(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom8(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom8(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom8 is a paid mutator transaction binding the contract method 0x552a1b56.
//
// Solidity: function transferFrom8(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom8(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom8(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom9 is a paid mutator transaction binding the contract method 0x4f7bd75a.
//
// Solidity: function transferFrom9(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactor) TransferFrom9(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom9", from, to, value)
}

// TransferFrom9 is a paid mutator transaction binding the contract method 0x4f7bd75a.
//
// Solidity: function transferFrom9(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractSession) TransferFrom9(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom9(&_Contract.TransactOpts, from, to, value)
}

// TransferFrom9 is a paid mutator transaction binding the contract method 0x4f7bd75a.
//
// Solidity: function transferFrom9(address from, address to, uint256 value) returns(bool)
func (_Contract *ContractTransactorSession) TransferFrom9(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom9(&_Contract.TransactOpts, from, to, value)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
