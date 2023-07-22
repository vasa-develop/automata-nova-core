## Basics of the core Game logic 

Generally, ticks run in an infinite loop of player's `main` module. It is important to understand that this loop is turn- and multiplayer-based: the next tick (next `Game.getTime()` value) begins only after the full execution of all `main` modules of all players.


We will now analyze the tick execution mechanism with conditional dividing it into beginning, middle, and end stages.

In the beginning of the tick, there is a certain game situation: different game objects with some property values. Take note that any changes in these properties, appearing of new objects, and dismantling of old ones will happen only at the start of the next tick.

In the middle part of the tick, the main module is executed (along with the modules required from it). The main operates the unchangeable game condition in the beginning of the tick. For example, by executing automaton.move() and then (in the same tick) automaton.attack() the attack still runs from the old coordinates, since properties automaton.pos.x / automaton.pos.y will change only in the next tick.

In the end of the tick, the commands specified in the main accumulate in order to change the game situation by the beginning of the next tick instantaneously and independently from each other. If any conflicts arise – for example, multiple automatons want to move to the same coordinates, or you have scheduled contradictory orders – these conflicts are solved according to predefined priorities. Another example: a mutual attack does not result in a conflict, and automatons can die at the same time.