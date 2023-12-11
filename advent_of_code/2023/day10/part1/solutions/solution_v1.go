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
		var puzzleLine []rune
		for _, char := range line {
			puzzleLine = append(puzzleLine, char)
		}
		s.puzzle = append(s.puzzle, puzzleLine)
	}

	// find starting position
	s.startPosition = s.FindStartingPosition()

	if s.debug {
		fmt.Printf("Starting position found at x: %d, y: %d\n", s.startPosition.x, s.startPosition.y)
	}

	// Find candidates around
	tileCandidates := s.FindTileCandidates(s.startPosition)

	if s.debug {
		fmt.Printf("Candidates: %+v\n", tileCandidates)
	}

	tilesCount := 0
	currentPosition := tileCandidates[0].Position
	nextDirection := tileCandidates[0].Out

	for {
		tilesCount += 1 // it will count the start position as the last step
		currentTile := TileType(s.puzzle[currentPosition.y][currentPosition.x])

		if currentTile == TileTypeStart {
			break
		}

		currentPosition = GetNextPosition(currentPosition, nextDirection)

		newTile := TileType(s.puzzle[currentPosition.y][currentPosition.x])
		nextDirection = FindNextDirection(TileType(newTile), nextDirection)
	}

	if s.debug {
		fmt.Printf("Total tiles count: %d\n", tilesCount)
	}

	s.result = tilesCount / 2
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

func (s *Solver) FindTileCandidates(startPosition Position) []PositionDirection {
	var candidates []PositionDirection

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
		}
	}

	return candidates
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
