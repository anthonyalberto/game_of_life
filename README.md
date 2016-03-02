# Game of Life
Go implementation of Conway's Game of Life

## Installation

```sh
  go get github.com/anthonyalberto/gameoflife
```

## Usage

Provided `$GOPATH/bin` is in your `$PATH` :

```sh
  gameoflife [options]
```

Supported options :
- `-board`: `regular` or `toroidal`. Toroidal wraps the board around. Default: `toroidal`
- `-time`: how many ms to pause between generations. Default: `50`
- `-pattern`: pattern name. Look into the `patterns` folder for all supported values. Default: `glidergun`
- `-width`: the width of the board. Adapt to your terminal. Default: `150`, suggested > `80` to support all patterns.
- `-height`: the height of the board. Adapt to your terminal. Default: `50`, suggested > `35` to support all patterns.Â

### Example

Play a game that:
- has the `glidergun` pattern
- has a toroidal board that is 140x54
- pauses between generations for 40ms

```sh
  gameoflife -board=toroidal -time=40 -pattern=glidergun -width=140 -height=54
```

Look at that gun being destroyed by its own bullets!
