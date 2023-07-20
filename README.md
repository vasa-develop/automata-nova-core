# automata-nova-core

Defines the core logic of automata nova using Concrete framework.

## Directory structure

```bash
├── engine
│   ├── main.go                 # Chain setup
│   ├── main_test.go            # Minimal setup test
│   ├── pcs
│   │   ├── abi
│   │   │   └── Game.json   # Precompile ABI created by `yarn build:abi`
│   │   ├── game.go         # Example precompile
│   │   └── game_test.go    # Precompile test
└── sol
    └── Game.sol            # Precompile interface in Solidity
```

## Running a devnet

Run a normal Optimism Bedrock devnet replacing op-geth for `github.com/therealbytes/concrete-template-geth:latest` in `ops-bedrock/Dockerfile.l2`.

Alternatively, clone and run an already modified version of Bedrock:

```bash
# Clone repo
git clone -b concrete-template https://github.com/therealbytes/optimism.git
cd optimism
# Start devnet
make devnet-up
# Stop and clean devnet
make devnet-down && make devnet-clean
```
