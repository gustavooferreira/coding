package solutions

import (
	"fmt"
)

type Solver struct {
	debug        bool
	inputContent []string
	result       int

	// puzzle specific fields
	puzzle        [][]rune
	startPosition Position
	tiles         [][]TileState
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) SetDebug(enable bool) {
	s.debug = enable
}

// LoadArrayOfLines loads the array in the internal representation of the input content.
func (s *Solver) LoadArrayOfLines(lines []string) {
	for _, line := range lines {
		s.LoadLine(line)
	}
}

// LoadLine loads the input line into the internal representation of the input content.
func (s *Solver) LoadLine(line string) {
	s.inputContent = append(s.inputContent, line)
}

// ComputeResult computes the result and stores it in the result field.
func (s *Solver) ComputeResult() {
	// create puzzle representation
	for _, line := range s.inputContent {
		if len(line) == 0 {
			continue
		}

		puzzleLine := make([]rune, 0, len(line))
		tilesLine := make([]TileState, 0, len(line))

		for _, char := range line {
			puzzleLine = append(puzzleLine, char)
			tilesLine = append(tilesLine, TileState{
				TileContainerType: TileContainerType_Cell,
			})
		}
		s.puzzle = append(s.puzzle, puzzleLine)
		s.tiles = append(s.tiles, tilesLine)
	}

	// find starting position
	s.startPosition = s.FindStartingPosition()

	if s.debug {
		fmt.Printf("Starting position found at x: %d, y: %d\n", s.startPosition.x, s.startPosition.y)
	}

	// Find candidates around
	tileCandidates, startTileType := s.FindTileCandidates(s.startPosition)

	if s.debug {
		fmt.Printf("Candidates: %+v\n", tileCandidates)
	}

	// setup map with the walls part of the path only

	currentPosition := tileCandidates[0].Position
	nextDirection := tileCandidates[0].Out
	for {
		currentTileType := TileType(s.puzzle[currentPosition.y][currentPosition.x])

		tile := s.tiles[currentPosition.y][currentPosition.x]
		tile.TileContainerType = TileContainerType_Wall
		tile.WallType = currentTileType
		s.tiles[currentPosition.y][currentPosition.x] = tile

		if currentTileType == TileTypeStart {
			break
		}

		currentPosition = GetNextPosition(currentPosition, nextDirection)

		newTile := TileType(s.puzzle[currentPosition.y][currentPosition.x])
		nextDirection = FindNextDirection(TileType(newTile), nextDirection)
	}

	// Replace S with correct pipe type
	s.tiles[s.startPosition.y][s.startPosition.x].WallType = startTileType

	if s.debug {
		fmt.Println("Walls:")
		for _, line := range s.tiles {
			for _, tile := range line {
				if tile.TileContainerType == TileContainerType_Wall {
					fmt.Printf("%s", string(tile.WallType))
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	// Start with initial state of all cells marked as being inside
	for _, line := range s.tiles {
		for _, tile := range line {
			if tile.TileContainerType == TileContainerType_Cell {
				tile.TileContainerLocation = TileContainerLocation_Inside
			}
		}
	}

	// do stuff here

	if s.debug {
		fmt.Println("Tiles State:")
		for _, line := range s.tiles {
			for _, tile := range line {
				if tile.TileContainerType == TileContainerType_Wall {
					fmt.Printf("%s", string(tile.WallType))
				} else {
					if tile.TileContainerLocation == TileContainerLocation_Inside {
						fmt.Printf("I")
					} else if tile.TileContainerLocation == TileContainerLocation_Outside {
						fmt.Printf("O")
					}
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	// Count how many tiles are of type cell and are inside the loop
	result := 0
	for _, line := range s.tiles {
		for _, tile := range line {
			if tile.TileContainerType == TileContainerType_Cell && tile.TileContainerLocation == TileContainerLocation_Inside {
				result++
			}
		}
	}

	s.result = result
}

// Result returns the current result stored in the Solver.
func (s *Solver) Result() int {
	// Not exporting the variable makes sure the user only gets read-only access to the underlying field.
	return s.result
}

// FindStartingPosition finds where the character S is, i.e., where is the starting position.
func (s *Solver) FindStartingPosition() Position {
	for j, line := range s.puzzle {
		for i, char := range line {
			if char == 'S' {
				return Position{
					x: i,
					y: j,
				}
			}
		}
	}

	return Position{}
}

func (s *Solver) FindTileCandidates(startPosition Position) (posDirections []PositionDirection, startTileType TileType) {
	var candidates []PositionDirection
	var connections []Direction

	if startPosition.y-1 >= 0 {
		tileType := TileType(s.puzzle[startPosition.y-1][startPosition.x])
		tileInfo := PipeInfoMap[tileType]
		if tileInfo.Terminal1 == DirectionSouth || tileInfo.Terminal2 == DirectionSouth {
			outDirection := FindNextDirection(tileType, DirectionNorth)
			candidates = append(candidates, PositionDirection{
				Position: Position{
					x: startPosition.x,
					y: startPosition.y - 1,
				},
				Out: outDirection,
			})
			connections = append(connections, DirectionNorth)
		}
	}

	if startPosition.y+1 < len(s.puzzle) {
		tileType := TileType(s.puzzle[startPosition.y+1][startPosition.x])
		tileInfo := PipeInfoMap[tileType]
		if tileInfo.Terminal1 == DirectionNorth || tileInfo.Terminal2 == DirectionNorth {
			outDirection := FindNextDirection(tileType, DirectionSouth)
			candidates = append(candidates, PositionDirection{
				Position: Position{
					x: startPosition.x,
					y: startPosition.y + 1,
				},
				Out: outDirection,
			})
			connections = append(connections, DirectionSouth)
		}
	}

	if startPosition.x-1 >= 0 {
		tileType := TileType(s.puzzle[startPosition.y][startPosition.x-1])
		tileInfo := PipeInfoMap[tileType]
		if tileInfo.Terminal1 == DirectionEast || tileInfo.Terminal2 == DirectionEast {
			outDirection := FindNextDirection(tileType, DirectionWest)
			candidates = append(candidates, PositionDirection{
				Position: Position{
					x: startPosition.x - 1,
					y: startPosition.y,
				},
				Out: outDirection,
			})
			connections = append(connections, DirectionWest)
		}
	}

	if startPosition.x+1 < len(s.puzzle[startPosition.y]) {
		tileType := TileType(s.puzzle[startPosition.y][startPosition.x+1])
		tileInfo := PipeInfoMap[tileType]
		if tileInfo.Terminal1 == DirectionWest || tileInfo.Terminal2 == DirectionWest {
			outDirection := FindNextDirection(tileType, DirectionEast)
			candidates = append(candidates, PositionDirection{
				Position: Position{
					x: startPosition.x + 1,
					y: startPosition.y,
				},
				Out: outDirection,
			})
			connections = append(connections, DirectionEast)
		}
	}

	// work out start tile type
	// this assumes only two candidates
	for k, v := range PipeInfoMap {
		if (v.Terminal1 == connections[0] && v.Terminal2 == connections[1]) ||
			(v.Terminal1 == connections[1] && v.Terminal2 == connections[0]) {
			startTileType = k
		}
	}

	return candidates, startTileType
}

// PositionDirection holds a position and the out direction for the next tile.
type PositionDirection struct {
	Position Position
	Out      Direction
}

type Position struct {
	x int
	y int
}

type TileType rune

const (
	TileTypeVertical          TileType = '|'
	TileTypeHorizontal        TileType = '-'
	TileTypeBottomLeftCorner  TileType = 'L'
	TileTypeBottomRightCorner TileType = 'J'
	TileTypeTopLeftCorner     TileType = 'F'
	TileTypeTopRightCorner    TileType = '7'
	TileTypeGround            TileType = '.'
	TileTypeStart             TileType = 'S'
)

type Direction string

const (
	DirectionNorth Direction = "north"
	DirectionSouth Direction = "south"
	DirectionEast  Direction = "east"
	DirectionWest  Direction = "west"
)

type Pipe struct {
	Terminal1 Direction
	Terminal2 Direction
}

var PipeInfoMap = map[TileType]Pipe{
	TileTypeVertical: {
		Terminal1: DirectionNorth,
		Terminal2: DirectionSouth,
	},
	TileTypeHorizontal: {
		Terminal2: DirectionWest,
		Terminal1: DirectionEast,
	},
	TileTypeBottomLeftCorner: {
		Terminal1: DirectionNorth,
		Terminal2: DirectionEast,
	},
	TileTypeBottomRightCorner: {
		Terminal1: DirectionNorth,
		Terminal2: DirectionWest,
	},
	TileTypeTopRightCorner: {
		Terminal1: DirectionSouth,
		Terminal2: DirectionWest,
	},
	TileTypeTopLeftCorner: {
		Terminal1: DirectionSouth,
		Terminal2: DirectionEast,
	},
}

func FindNextDirection(tile TileType, inputDirection Direction) (outputDirection Direction) {
	tileInfo := PipeInfoMap[tile]

	// Invert input direction
	// If next direction is south for example, the pipe needs to have a terminal pointing north.
	inputDirectionReversed := ReverseDirection(inputDirection)

	if tileInfo.Terminal1 == inputDirectionReversed {
		return tileInfo.Terminal2
	} else if tileInfo.Terminal2 == inputDirectionReversed {
		return tileInfo.Terminal1
	}

	// handle error as this should never happen
	return
}

func ReverseDirection(inputDirection Direction) (outputDirection Direction) {
	switch inputDirection {
	case DirectionNorth:
		return DirectionSouth
	case DirectionSouth:
		return DirectionNorth
	case DirectionWest:
		return DirectionEast
	case DirectionEast:
		return DirectionWest
	}

	// this is an error!
	return
}

// Given current position, and a direction, return the next position
func GetNextPosition(currentPos Position, direction Direction) Position {
	x := currentPos.x
	y := currentPos.y

	switch direction {
	case DirectionNorth:
		y -= 1
	case DirectionSouth:
		y += 1
	case DirectionWest:
		x -= 1
	case DirectionEast:
		x += 1
	}

	return Position{
		x: x,
		y: y,
	}
}

// Defines whether the tile contains a wall or a cell
type TileContainerType string

const (
	TileContainerType_Wall TileContainerType = "wall"
	TileContainerType_Cell TileContainerType = "cell"
)

type TileState struct {
	TileContainerType     TileContainerType
	WallType              TileType
	TileContainerLocation TileContainerLocation
}

type TileContainerLocation string

const (
	TileContainerLocation_Inside  TileContainerLocation = "inside"
	TileContainerLocation_Outside TileContainerLocation = "outside"
)
