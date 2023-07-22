package main

import (
	"github.com/vasa-develop/automata-nova-core/engine/pcs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete"
	"github.com/ethereum/go-ethereum/concrete/precompiles"
)

func setup(engine concrete.ConcreteApp) {
	address_1 := common.HexToAddress("0x80")
	metadata_1 := precompiles.PrecompileMetadata{
		Name:        "Automata Nova Core",
		Version:     precompiles.NewVersion(0, 0, 0),
		Author:      "Automata Nova authors",
		Description: "Core logic for Automata Nova Game",
		Source:      "https://github.com/vasa-develop/automata-nova-core/blob/main/engine/pcs/game.go",
	}
	engine.AddPrecompile(address_1, pcs.GamePrecompile, metadata_1)

	address_2 := common.HexToAddress("0x81")
	metadata_2 := precompiles.PrecompileMetadata{
		Name:        "Automaton",
		Version:     precompiles.NewVersion(0, 0, 0),
		Author:      "Automata Nova authors",
		Description: "Core logic for Automaton",
		Source:      "https://github.com/vasa-develop/automata-nova-core/blob/main/engine/pcs/automaton.go",
	}
	engine.AddPrecompile(address_2, pcs.AutomatonPrecompile, metadata_2)
}

func main() {
	engine := concrete.ConcreteGeth
	setup(engine)
	engine.Run()
}
