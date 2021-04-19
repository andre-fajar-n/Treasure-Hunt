# Treasure-Hunt
### Description
Given layout grid like this
```
# # # # # # # #
# . . . . . . #
# . # # # . . #
# . . . # . # #
# X # . . . . #
# # # # # # # #
```
**\#** represents an obstacle.</br>
**.** represents a clear path.</br>
**X** represents the player’s starting position.

A treasure is hidden within one of the clear path points. From the starting position, the user must navigate in a specific order:
- Up/North A step(s), then
- Right/East B step(s), then
- Down/South C step(s).

The program show list of probable coordinate points where the treasure might be located, navigation to reach the treasure, and the grid with the probable treasure locations marked with **$** symbol.

Example:
```
Up/North 1 step(s)
Right/East 2 step(s)
Down/South 1 step(s)
# # # # # # # #
# . . . . . . #
# . # # # . . #
# . . . # . # #
# X # $ . . . #
# # # # # # # #
```

### Setup package
```
go mod vendor
```

### How to run program 
```
go run main.go
```
