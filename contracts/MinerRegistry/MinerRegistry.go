// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MinerRegistry

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

)

	//定义POSMINER访问客户端
var(
		PosConn *ethclient.Client
)	

// MinerRegistryABI is the input ABI used to generate the binding from.
const MinerRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"certifyInformation\",\"type\":\"uint256\"}],\"name\":\"MinerReg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"MinerRegistryTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"Miner\",\"type\":\"address\"},{\"name\":\"CertInfo\",\"type\":\"uint256\"}],\"name\":\"CertiyInfoSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ReceiveFoundation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"RegMac\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"bool\"},{\"name\":\"AppAddr\",\"type\":\"string\"}],\"name\":\"RegMachineSet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"transferManagement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"RegMachine\",\"outputs\":[{\"name\":\"status\",\"type\":\"bool\"},{\"name\":\"AppAddr\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// MinerRegistryBin is the compiled bytecode used for deploying new contracts.
const MinerRegistryBin = `{
	"linkReferences": {},
	"object": "6060604052341561000f57600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506107e98061005e6000396000f30060606040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063118818c4146100a05780634012a32e146100c35780635eb11b121461011057806378357e5314610152578063a22eef56146101a7578063b7ea4cca146101d0578063e4edf85214610257578063f57a659214610290575b34600160008282540192505081905550005b34156100ab57600080fd5b6100c1600480803590602001909190505061036a565b005b34156100ce57600080fd5b6100fa600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506103f9565b6040518082815260200191505060405180910390f35b341561011b57600080fd5b610150600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035906020019091905050610411565b005b341561015d57600080fd5b610165610529565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156101b257600080fd5b6101ba61054e565b6040518082815260200191505060405180910390f35b34156101db57600080fd5b610255600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091908035151590602001909190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610554565b005b341561026257600080fd5b61028e600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061064a565b005b341561029b57600080fd5b6102c7600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506106e8565b60405180831515151581526020018060200182810382528381815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561035a5780601f1061032f5761010080835404028352916020019161035a565b820191906000526020600020905b81548152906001019060200180831161033d57829003601f168201915b5050935050505060405180910390f35b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156103f65742600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b50565b60046020528060005260406000206000915090505481565b60011515600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900460ff16151514151561047357600080fd5b80600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550662386f26fc100008273ffffffffffffffffffffffffffffffffffffffff16311015610525578173ffffffffffffffffffffffffffffffffffffffff166108fc662386f26fc100009081150290604051600060405180830381858888f19350505050151561052457600080fd5b5b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105af57600080fd5b6040805190810160405280831515815260200182815250600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190610641929190610718565b50905050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106a557600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60026020528060005260406000206000915090508060000160009054906101000a900460ff169080600101905082565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061075957805160ff1916838001178555610787565b82800160010185558215610787579182015b8281111561078657825182559160200191906001019061076b565b5b5090506107949190610798565b5090565b6107ba91905b808211156107b657600081600090555060010161079e565b5090565b905600a165627a7a72305820064d0c0795d6f370564f0269209ad3f4c758215607a1971a1f6c81584e2238300029",
	"opcodes": "PUSH1 0x60 PUSH1 0x40 MSTORE CALLVALUE ISZERO PUSH2 0xF JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST CALLER PUSH1 0x0 DUP1 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP PUSH2 0x7E9 DUP1 PUSH2 0x5E PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN STOP PUSH1 0x60 PUSH1 0x40 MSTORE PUSH1 0x4 CALLDATASIZE LT PUSH2 0x8E JUMPI PUSH1 0x0 CALLDATALOAD PUSH29 0x100000000000000000000000000000000000000000000000000000000 SWAP1 DIV PUSH4 0xFFFFFFFF AND DUP1 PUSH4 0x118818C4 EQ PUSH2 0xA0 JUMPI DUP1 PUSH4 0x4012A32E EQ PUSH2 0xC3 JUMPI DUP1 PUSH4 0x5EB11B12 EQ PUSH2 0x110 JUMPI DUP1 PUSH4 0x78357E53 EQ PUSH2 0x152 JUMPI DUP1 PUSH4 0xA22EEF56 EQ PUSH2 0x1A7 JUMPI DUP1 PUSH4 0xB7EA4CCA EQ PUSH2 0x1D0 JUMPI DUP1 PUSH4 0xE4EDF852 EQ PUSH2 0x257 JUMPI DUP1 PUSH4 0xF57A6592 EQ PUSH2 0x290 JUMPI JUMPDEST CALLVALUE PUSH1 0x1 PUSH1 0x0 DUP3 DUP3 SLOAD ADD SWAP3 POP POP DUP2 SWAP1 SSTORE POP STOP JUMPDEST CALLVALUE ISZERO PUSH2 0xAB JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0xC1 PUSH1 0x4 DUP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH2 0x36A JUMP JUMPDEST STOP JUMPDEST CALLVALUE ISZERO PUSH2 0xCE JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0xFA PUSH1 0x4 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH2 0x3F9 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE ISZERO PUSH2 0x11B JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x150 PUSH1 0x4 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH2 0x411 JUMP JUMPDEST STOP JUMPDEST CALLVALUE ISZERO PUSH2 0x15D JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x165 PUSH2 0x529 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE ISZERO PUSH2 0x1B2 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x1BA PUSH2 0x54E JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP3 DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST CALLVALUE ISZERO PUSH2 0x1DB JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x255 PUSH1 0x4 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 DUP1 CALLDATALOAD ISZERO ISZERO SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP3 ADD DUP1 CALLDATALOAD SWAP1 PUSH1 0x20 ADD SWAP1 DUP1 DUP1 PUSH1 0x1F ADD PUSH1 0x20 DUP1 SWAP2 DIV MUL PUSH1 0x20 ADD PUSH1 0x40 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 SWAP4 SWAP3 SWAP2 SWAP1 DUP2 DUP2 MSTORE PUSH1 0x20 ADD DUP4 DUP4 DUP1 DUP3 DUP5 CALLDATACOPY DUP3 ADD SWAP2 POP POP POP POP POP POP SWAP2 SWAP1 POP POP PUSH2 0x554 JUMP JUMPDEST STOP JUMPDEST CALLVALUE ISZERO PUSH2 0x262 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x28E PUSH1 0x4 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH2 0x64A JUMP JUMPDEST STOP JUMPDEST CALLVALUE ISZERO PUSH2 0x29B JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x2C7 PUSH1 0x4 DUP1 DUP1 CALLDATALOAD PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND SWAP1 PUSH1 0x20 ADD SWAP1 SWAP2 SWAP1 POP POP PUSH2 0x6E8 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 DUP4 ISZERO ISZERO ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD DUP1 PUSH1 0x20 ADD DUP3 DUP2 SUB DUP3 MSTORE DUP4 DUP2 DUP2 SLOAD PUSH1 0x1 DUP2 PUSH1 0x1 AND ISZERO PUSH2 0x100 MUL SUB AND PUSH1 0x2 SWAP1 DIV DUP2 MSTORE PUSH1 0x20 ADD SWAP2 POP DUP1 SLOAD PUSH1 0x1 DUP2 PUSH1 0x1 AND ISZERO PUSH2 0x100 MUL SUB AND PUSH1 0x2 SWAP1 DIV DUP1 ISZERO PUSH2 0x35A JUMPI DUP1 PUSH1 0x1F LT PUSH2 0x32F JUMPI PUSH2 0x100 DUP1 DUP4 SLOAD DIV MUL DUP4 MSTORE SWAP2 PUSH1 0x20 ADD SWAP2 PUSH2 0x35A JUMP JUMPDEST DUP3 ADD SWAP2 SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 JUMPDEST DUP2 SLOAD DUP2 MSTORE SWAP1 PUSH1 0x1 ADD SWAP1 PUSH1 0x20 ADD DUP1 DUP4 GT PUSH2 0x33D JUMPI DUP3 SWAP1 SUB PUSH1 0x1F AND DUP3 ADD SWAP2 JUMPDEST POP POP SWAP4 POP POP POP POP PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST DUP1 PUSH1 0x3 PUSH1 0x0 CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 SLOAD EQ ISZERO PUSH2 0x3F6 JUMPI TIMESTAMP PUSH1 0x4 PUSH1 0x0 CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 DUP2 SWAP1 SSTORE POP JUMPDEST POP JUMP JUMPDEST PUSH1 0x4 PUSH1 0x20 MSTORE DUP1 PUSH1 0x0 MSTORE PUSH1 0x40 PUSH1 0x0 KECCAK256 PUSH1 0x0 SWAP2 POP SWAP1 POP SLOAD DUP2 JUMP JUMPDEST PUSH1 0x1 ISZERO ISZERO PUSH1 0x2 PUSH1 0x0 CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 ADD PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND ISZERO ISZERO EQ ISZERO ISZERO PUSH2 0x473 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP1 PUSH1 0x3 PUSH1 0x0 DUP5 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 DUP2 SWAP1 SSTORE POP PUSH7 0x2386F26FC10000 DUP3 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND BALANCE LT ISZERO PUSH2 0x525 JUMPI DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH2 0x8FC PUSH7 0x2386F26FC10000 SWAP1 DUP2 ISZERO MUL SWAP1 PUSH1 0x40 MLOAD PUSH1 0x0 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP6 DUP9 DUP9 CALL SWAP4 POP POP POP POP ISZERO ISZERO PUSH2 0x524 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST JUMPDEST POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 JUMP JUMPDEST PUSH1 0x1 SLOAD DUP2 JUMP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x5AF JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x40 DUP1 MLOAD SWAP1 DUP2 ADD PUSH1 0x40 MSTORE DUP1 DUP4 ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD DUP3 DUP2 MSTORE POP PUSH1 0x2 PUSH1 0x0 DUP6 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 DUP2 MSTORE PUSH1 0x20 ADD PUSH1 0x0 KECCAK256 PUSH1 0x0 DUP3 ADD MLOAD DUP2 PUSH1 0x0 ADD PUSH1 0x0 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH1 0xFF MUL NOT AND SWAP1 DUP4 ISZERO ISZERO MUL OR SWAP1 SSTORE POP PUSH1 0x20 DUP3 ADD MLOAD DUP2 PUSH1 0x1 ADD SWAP1 DUP1 MLOAD SWAP1 PUSH1 0x20 ADD SWAP1 PUSH2 0x641 SWAP3 SWAP2 SWAP1 PUSH2 0x718 JUMP JUMPDEST POP SWAP1 POP POP POP POP POP JUMP JUMPDEST PUSH1 0x0 DUP1 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND CALLER PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND EQ ISZERO ISZERO PUSH2 0x6A5 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST DUP1 PUSH1 0x0 DUP1 PUSH2 0x100 EXP DUP2 SLOAD DUP2 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF MUL NOT AND SWAP1 DUP4 PUSH20 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF AND MUL OR SWAP1 SSTORE POP POP JUMP JUMPDEST PUSH1 0x2 PUSH1 0x20 MSTORE DUP1 PUSH1 0x0 MSTORE PUSH1 0x40 PUSH1 0x0 KECCAK256 PUSH1 0x0 SWAP2 POP SWAP1 POP DUP1 PUSH1 0x0 ADD PUSH1 0x0 SWAP1 SLOAD SWAP1 PUSH2 0x100 EXP SWAP1 DIV PUSH1 0xFF AND SWAP1 DUP1 PUSH1 0x1 ADD SWAP1 POP DUP3 JUMP JUMPDEST DUP3 DUP1 SLOAD PUSH1 0x1 DUP2 PUSH1 0x1 AND ISZERO PUSH2 0x100 MUL SUB AND PUSH1 0x2 SWAP1 DIV SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 PUSH1 0x1F ADD PUSH1 0x20 SWAP1 DIV DUP2 ADD SWAP3 DUP3 PUSH1 0x1F LT PUSH2 0x759 JUMPI DUP1 MLOAD PUSH1 0xFF NOT AND DUP4 DUP1 ADD OR DUP6 SSTORE PUSH2 0x787 JUMP JUMPDEST DUP3 DUP1 ADD PUSH1 0x1 ADD DUP6 SSTORE DUP3 ISZERO PUSH2 0x787 JUMPI SWAP2 DUP3 ADD JUMPDEST DUP3 DUP2 GT ISZERO PUSH2 0x786 JUMPI DUP3 MLOAD DUP3 SSTORE SWAP2 PUSH1 0x20 ADD SWAP2 SWAP1 PUSH1 0x1 ADD SWAP1 PUSH2 0x76B JUMP JUMPDEST JUMPDEST POP SWAP1 POP PUSH2 0x794 SWAP2 SWAP1 PUSH2 0x798 JUMP JUMPDEST POP SWAP1 JUMP JUMPDEST PUSH2 0x7BA SWAP2 SWAP1 JUMPDEST DUP1 DUP3 GT ISZERO PUSH2 0x7B6 JUMPI PUSH1 0x0 DUP2 PUSH1 0x0 SWAP1 SSTORE POP PUSH1 0x1 ADD PUSH2 0x79E JUMP JUMPDEST POP SWAP1 JUMP JUMPDEST SWAP1 JUMP STOP LOG1 PUSH6 0x627A7A723058 KECCAK256 MOD 0x4d 0xc SMOD SWAP6 0xd6 RETURN PUSH17 0x564F0269209AD3F4C758215607A1971A1F PUSH13 0x81584E22383000290000000000 ",
	"sourceMap": "28:1868:0:-;;;788:74;;;;;;;;843:10;835:7;;:18;;;;;;;;;;;;;;;;;;28:1868;;;;;;"
}`

// DeployMinerRegistry deploys a new Ethereum contract, binding an instance of MinerRegistry to it.
func DeployMinerRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MinerRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(MinerRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MinerRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MinerRegistry{MinerRegistryCaller: MinerRegistryCaller{contract: contract}, MinerRegistryTransactor: MinerRegistryTransactor{contract: contract}}, nil
}

// MinerRegistry is an auto generated Go binding around an Ethereum contract.
type MinerRegistry struct {
	MinerRegistryCaller     // Read-only binding to the contract
	MinerRegistryTransactor // Write-only binding to the contract
}

// MinerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinerRegistrySession struct {
	Contract     *MinerRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinerRegistryCallerSession struct {
	Contract *MinerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MinerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinerRegistryTransactorSession struct {
	Contract     *MinerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MinerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinerRegistryRaw struct {
	Contract *MinerRegistry // Generic contract binding to access the raw methods on
}

// MinerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinerRegistryCallerRaw struct {
	Contract *MinerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MinerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinerRegistryTransactorRaw struct {
	Contract *MinerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinerRegistry creates a new instance of MinerRegistry, bound to a specific deployed contract.
func NewMinerRegistry(address common.Address, backend bind.ContractBackend) (*MinerRegistry, error) {
	contract, err := bindMinerRegistry(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinerRegistry{MinerRegistryCaller: MinerRegistryCaller{contract: contract}, MinerRegistryTransactor: MinerRegistryTransactor{contract: contract}}, nil
}

// NewMinerRegistryCaller creates a new read-only instance of MinerRegistry, bound to a specific deployed contract.
func NewMinerRegistryCaller(address common.Address, caller bind.ContractCaller) (*MinerRegistryCaller, error) {
	contract, err := bindMinerRegistry(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MinerRegistryCaller{contract: contract}, nil
}

// NewMinerRegistryTransactor creates a new write-only instance of MinerRegistry, bound to a specific deployed contract.
func NewMinerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MinerRegistryTransactor, error) {
	contract, err := bindMinerRegistry(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MinerRegistryTransactor{contract: contract}, nil
}

// bindMinerRegistry binds a generic wrapper to an already deployed contract.
func bindMinerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinerRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerRegistry *MinerRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinerRegistry.Contract.MinerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerRegistry *MinerRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerRegistry.Contract.MinerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerRegistry *MinerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerRegistry.Contract.MinerRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerRegistry *MinerRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerRegistry *MinerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerRegistry *MinerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerRegistry.Contract.contract.Transact(opts, method, params...)
}

// Manager is a free data retrieval call binding the contract method 0x78357e53.
//
// Solidity: function Manager() constant returns(address)
func (_MinerRegistry *MinerRegistryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MinerRegistry.contract.Call(opts, out, "Manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x78357e53.
//
// Solidity: function Manager() constant returns(address)
func (_MinerRegistry *MinerRegistrySession) Manager() (common.Address, error) {
	return _MinerRegistry.Contract.Manager(&_MinerRegistry.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x78357e53.
//
// Solidity: function Manager() constant returns(address)
func (_MinerRegistry *MinerRegistryCallerSession) Manager() (common.Address, error) {
	return _MinerRegistry.Contract.Manager(&_MinerRegistry.CallOpts)
}

// MinerRegistryTime is a free data retrieval call binding the contract method 0x4012a32e.
//
// Solidity: function MinerRegistryTime( address) constant returns(uint256)
func (_MinerRegistry *MinerRegistryCaller) MinerRegistryTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MinerRegistry.contract.Call(opts, out, "MinerRegistryTime", arg0)
	return *ret0, err
}

// MinerRegistryTime is a free data retrieval call binding the contract method 0x4012a32e.
//
// Solidity: function MinerRegistryTime( address) constant returns(uint256)
func (_MinerRegistry *MinerRegistrySession) MinerRegistryTime(arg0 common.Address) (*big.Int, error) {
	return _MinerRegistry.Contract.MinerRegistryTime(&_MinerRegistry.CallOpts, arg0)
}

// MinerRegistryTime is a free data retrieval call binding the contract method 0x4012a32e.
//
// Solidity: function MinerRegistryTime( address) constant returns(uint256)
func (_MinerRegistry *MinerRegistryCallerSession) MinerRegistryTime(arg0 common.Address) (*big.Int, error) {
	return _MinerRegistry.Contract.MinerRegistryTime(&_MinerRegistry.CallOpts, arg0)
}

// ReceiveFoundation is a free data retrieval call binding the contract method 0xa22eef56.
//
// Solidity: function ReceiveFoundation() constant returns(uint256)
func (_MinerRegistry *MinerRegistryCaller) ReceiveFoundation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MinerRegistry.contract.Call(opts, out, "ReceiveFoundation")
	return *ret0, err
}

// ReceiveFoundation is a free data retrieval call binding the contract method 0xa22eef56.
//
// Solidity: function ReceiveFoundation() constant returns(uint256)
func (_MinerRegistry *MinerRegistrySession) ReceiveFoundation() (*big.Int, error) {
	return _MinerRegistry.Contract.ReceiveFoundation(&_MinerRegistry.CallOpts)
}

// ReceiveFoundation is a free data retrieval call binding the contract method 0xa22eef56.
//
// Solidity: function ReceiveFoundation() constant returns(uint256)
func (_MinerRegistry *MinerRegistryCallerSession) ReceiveFoundation() (*big.Int, error) {
	return _MinerRegistry.Contract.ReceiveFoundation(&_MinerRegistry.CallOpts)
}

// RegMachine is a free data retrieval call binding the contract method 0xf57a6592.
//
// Solidity: function RegMachine( address) constant returns(status bool, AppAddr string)
func (_MinerRegistry *MinerRegistryCaller) RegMachine(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status  bool
	AppAddr string
}, error) {
	ret := new(struct {
		Status  bool
		AppAddr string
	})
	out := ret
	err := _MinerRegistry.contract.Call(opts, out, "RegMachine", arg0)
	return *ret, err
}

// RegMachine is a free data retrieval call binding the contract method 0xf57a6592.
//
// Solidity: function RegMachine( address) constant returns(status bool, AppAddr string)
func (_MinerRegistry *MinerRegistrySession) RegMachine(arg0 common.Address) (struct {
	Status  bool
	AppAddr string
}, error) {
	return _MinerRegistry.Contract.RegMachine(&_MinerRegistry.CallOpts, arg0)
}

// RegMachine is a free data retrieval call binding the contract method 0xf57a6592.
//
// Solidity: function RegMachine( address) constant returns(status bool, AppAddr string)
func (_MinerRegistry *MinerRegistryCallerSession) RegMachine(arg0 common.Address) (struct {
	Status  bool
	AppAddr string
}, error) {
	return _MinerRegistry.Contract.RegMachine(&_MinerRegistry.CallOpts, arg0)
}

// CertiyInfoSet is a paid mutator transaction binding the contract method 0x5eb11b12.
//
// Solidity: function CertiyInfoSet(Miner address, CertInfo uint256) returns()
func (_MinerRegistry *MinerRegistryTransactor) CertiyInfoSet(opts *bind.TransactOpts, Miner common.Address, CertInfo *big.Int) (*types.Transaction, error) {
	return _MinerRegistry.contract.Transact(opts, "CertiyInfoSet", Miner, CertInfo)
}

// CertiyInfoSet is a paid mutator transaction binding the contract method 0x5eb11b12.
//
// Solidity: function CertiyInfoSet(Miner address, CertInfo uint256) returns()
func (_MinerRegistry *MinerRegistrySession) CertiyInfoSet(Miner common.Address, CertInfo *big.Int) (*types.Transaction, error) {
	return _MinerRegistry.Contract.CertiyInfoSet(&_MinerRegistry.TransactOpts, Miner, CertInfo)
}

// CertiyInfoSet is a paid mutator transaction binding the contract method 0x5eb11b12.
//
// Solidity: function CertiyInfoSet(Miner address, CertInfo uint256) returns()
func (_MinerRegistry *MinerRegistryTransactorSession) CertiyInfoSet(Miner common.Address, CertInfo *big.Int) (*types.Transaction, error) {
	return _MinerRegistry.Contract.CertiyInfoSet(&_MinerRegistry.TransactOpts, Miner, CertInfo)
}

// MinerReg is a paid mutator transaction binding the contract method 0x118818c4.
//
// Solidity: function MinerReg(certifyInformation uint256) returns()
func (_MinerRegistry *MinerRegistryTransactor) MinerReg(opts *bind.TransactOpts, certifyInformation *big.Int) (*types.Transaction, error) {
	return _MinerRegistry.contract.Transact(opts, "MinerReg", certifyInformation)
}

// MinerReg is a paid mutator transaction binding the contract method 0x118818c4.
//
// Solidity: function MinerReg(certifyInformation uint256) returns()
func (_MinerRegistry *MinerRegistrySession) MinerReg(certifyInformation *big.Int) (*types.Transaction, error) {
	return _MinerRegistry.Contract.MinerReg(&_MinerRegistry.TransactOpts, certifyInformation)
}

// MinerReg is a paid mutator transaction binding the contract method 0x118818c4.
//
// Solidity: function MinerReg(certifyInformation uint256) returns()
func (_MinerRegistry *MinerRegistryTransactorSession) MinerReg(certifyInformation *big.Int) (*types.Transaction, error) {
	return _MinerRegistry.Contract.MinerReg(&_MinerRegistry.TransactOpts, certifyInformation)
}

// RegMachineSet is a paid mutator transaction binding the contract method 0xb7ea4cca.
//
// Solidity: function RegMachineSet(RegMac address, status bool, AppAddr string) returns()
func (_MinerRegistry *MinerRegistryTransactor) RegMachineSet(opts *bind.TransactOpts, RegMac common.Address, status bool, AppAddr string) (*types.Transaction, error) {
	return _MinerRegistry.contract.Transact(opts, "RegMachineSet", RegMac, status, AppAddr)
}

// RegMachineSet is a paid mutator transaction binding the contract method 0xb7ea4cca.
//
// Solidity: function RegMachineSet(RegMac address, status bool, AppAddr string) returns()
func (_MinerRegistry *MinerRegistrySession) RegMachineSet(RegMac common.Address, status bool, AppAddr string) (*types.Transaction, error) {
	return _MinerRegistry.Contract.RegMachineSet(&_MinerRegistry.TransactOpts, RegMac, status, AppAddr)
}

// RegMachineSet is a paid mutator transaction binding the contract method 0xb7ea4cca.
//
// Solidity: function RegMachineSet(RegMac address, status bool, AppAddr string) returns()
func (_MinerRegistry *MinerRegistryTransactorSession) RegMachineSet(RegMac common.Address, status bool, AppAddr string) (*types.Transaction, error) {
	return _MinerRegistry.Contract.RegMachineSet(&_MinerRegistry.TransactOpts, RegMac, status, AppAddr)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(newManager address) returns()
func (_MinerRegistry *MinerRegistryTransactor) TransferManagement(opts *bind.TransactOpts, newManager common.Address) (*types.Transaction, error) {
	return _MinerRegistry.contract.Transact(opts, "transferManagement", newManager)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(newManager address) returns()
func (_MinerRegistry *MinerRegistrySession) TransferManagement(newManager common.Address) (*types.Transaction, error) {
	return _MinerRegistry.Contract.TransferManagement(&_MinerRegistry.TransactOpts, newManager)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(newManager address) returns()
func (_MinerRegistry *MinerRegistryTransactorSession) TransferManagement(newManager common.Address) (*types.Transaction, error) {
	return _MinerRegistry.Contract.TransferManagement(&_MinerRegistry.TransactOpts, newManager)
}