package pcs

import (
	_ "embed"
	"encoding/json"

	// "math/big"

	// "math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	// "github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete/api"
	"github.com/ethereum/go-ethereum/concrete/lib"
	// "github.com/ethereum/go-ethereum/crypto"
)

//go:embed abi/Automaton.json
var automatonAbiJson []byte

var (
	AutomatonABI        abi.ABI
	AutomatonPrecompile api.Precompile
)

var (
	// This is the storage slot for the mapping players to their owned rooms
	// PlayerMapKey = crypto.Keccak256Hash([]byte("player.to.room.map"))

	// This is the storage slot for the tracking the game time
	// GameTimeReferenceKey = crypto.Keccak256Hash([]byte("game.time.reference"))
)

func init() {
	// Get the ABI from the JSON file
	var jsonAbi struct {
		ABI abi.ABI `json:"abi"`
	}
	err := json.Unmarshal(automatonAbiJson, &jsonAbi)
	if err != nil {
		panic(err)
	}
	AutomatonABI = jsonAbi.ABI
	// Create the precompile
	AutomatonPrecompile = lib.NewPrecompileWithABI(AutomatonABI, map[string]lib.MethodPrecompile{
		"getMessage": &getMessageMethod{},
	})
}

type getMessageMethod struct {
	lib.BlankMethodPrecompile
}

func (p *getMessageMethod) RequiredGas(input []byte) uint64 {
	// Simple fixed gas cost, note this is oversimplified and does not actually reflect
	// the underlying resource usage of the operation.
	return 100
}

func (p *getMessageMethod) Run(concrete api.API, input []byte) ([]byte, error) {
	return p.CallRunWithArgs(func(concrete api.API, args []interface{}) ([]interface{}, error) {
		// idBN := args[0].(*big.Int)
		// idHash := common.BigToHash(idBN)
		// counters := concrete.Persistent().NewMap(CounterMapKey)
		// value := counters.Get(idHash).Big()
		// return []interface{}{value}, nil

		return []interface{}{GAME_VERSION}, nil
	}, concrete, input)
}

/* type incrementMethod struct {
	lib.BlankMethodPrecompile
}

func (p *incrementMethod) RequiredGas(input []byte) uint64 {
	// Simple fixed gas cost, note this is oversimplified and does not actually reflect
	// the underlying resource usage of the operation.
	return 10_000
}

func (p *incrementMethod) Run(concrete api.API, input []byte) ([]byte, error) {
	return p.CallRunWithArgs(func(concrete api.API, args []interface{}) ([]interface{}, error) {
		// Get the counter ID
		idBN := args[0].(*big.Int)
		// Convert it to a Hash type -- a 32-byte array
		// Note BigToHash is a type conversion, not a cryptographic hash!
		idHash := common.BigToHash(idBN)
		// Get the map of counters
		counters := concrete.Persistent().NewMap(CounterMapKey)
		// Get the current value of the counter as a BigInt
		value := counters.Get(idHash).Big()
		// Increment the counter
		value = value.Add(value, big.NewInt(1))
		// Store the new value
		counters.Set(idHash, common.BigToHash(value))
		// Return the new value
		return []interface{}{value}, nil
	}, concrete, input)
}

// Same as above, but without the increment
type getMethod struct {
	lib.BlankMethodPrecompile
}

func (p *getMethod) RequiredGas(input []byte) uint64 {
	return 2_000
}

func (p *getMethod) Run(concrete api.API, input []byte) ([]byte, error) {
	return p.CallRunWithArgs(func(concrete api.API, args []interface{}) ([]interface{}, error) {
		idBN := args[0].(*big.Int)
		idHash := common.BigToHash(idBN)
		counters := concrete.Persistent().NewMap(CounterMapKey)
		value := counters.Get(idHash).Big()
		return []interface{}{value}, nil
	}, concrete, input)
} */
