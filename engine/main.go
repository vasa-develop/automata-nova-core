package main

import (
	"github.com/vasa-develop/automata-nova-core/engine/pcs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/concrete"
	"github.com/ethereum/go-ethereum/concrete/precompiles"
)

func setup(engine concrete.ConcreteApp) {
	address := common.HexToAddress("0x80")
	metadata := precompiles.PrecompileMetadata{
		Name:        "Automata Nova Core",
		Version:     precompiles.NewVersion(0, 1, 0),
		Author:      "Automata Nova authors",
		Description: "Core logic for Automata Nova Game",
		Source:      "https://github.com/vasa-develop/automata-nova-core/blob/main/engine/pcs/game.go",
	}
	engine.AddPrecompile(address, pcs.GamePrecompile, metadata)
}

func main() {
	engine := concrete.ConcreteGeth
	setup(engine)
	engine.Run()
}
