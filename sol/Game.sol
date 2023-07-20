// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// The address the precompile will be mapped to.
address constant GAME_ADDRESS = address(0x80);

interface Game {
    // Returns the game version.
    function gameVersion() external pure returns (string memory);

    // Adds a player to the Game.
    // function joinGame(address playerAddress, uint8 roomId) external;
    
    // Registers player script address.
    // The script address will implement an interface that will be called by the game.
    // function registerScript(address playerAddress, address scriptAddress) external;

    // Executes game tick.
    // It runs each player's script and updates the game state.
    // function tick() external returns (uint256);
}
