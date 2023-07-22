// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Structure } from "./Structure.sol";

// @dev A structure that can store huge amount of resource units. Only one structure per room is allowed.
interface StructureStorage is Structure {
    // @dev Gets the player address that owns the structure.
    // @param structureId The id of the structure.
    // @return The player address that owns the structure.
    function getOwner(bytes32 structureId) external view returns (address);

    // @dev Gets the store id that contains cargo of the structure.
    // @param structureId The id of the structure.
    // @return The store id that contains cargo of the structure.
    function getStoreId(bytes32 structureId) external view returns (bytes32);
}