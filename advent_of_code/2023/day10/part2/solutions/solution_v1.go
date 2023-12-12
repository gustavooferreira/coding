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

	// setup map with the walls part of the path only (and substitute start with correct pipe)
	s.SetTilesForWallsOnly(tileCandidates[0], startTileType)

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

	// Mark all cells as being inside
	for i, tilesLine := range s.tiles {
		for j, tile := range tilesLine {
			if tile.TileContainerType == TileContainerType_Cell {
				s.tiles[i][j].TileContainerLocation = TileContainerLocation_Inside
			}
		}
	}

	// expand puzzle, rows, columns and borders
	s.ExpandPuzzle()

	// Run Flood fill algo and mark cells we find as being outside.

	// Count how many tiles are of type cell and are inside the loop
	result := 0
	for _, line := range s.tiles {
		for _, tile := range line {
			if tile.TileContainerType == TileContainerType_Cell && tile.TileContainerLocation == TileContainerLocation_Inside {
				result++
			}
		}
	}

	// if s.debug {
	// 	fmt.Println("Tiles State:")
	// 	for _, line := range s.tiles {
	// 		for _, tile := range line {
	// 			if tile.TileContainerType == TileContainerType_Wall {
	// 				fmt.Printf("%s", string(tile.WallType))
	// 			} else {
	// 				if tile.TileContainerLocation == TileContainerLocation_Inside {
	// 					fmt.Printf("I")
	// 				} else if tile.TileContainerLocation == TileContainerLocation_Outside {
	// 					fmt.Printf("O")
	// 				}
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}
	// 	fmt.Println()
	// }

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

func (s *Solver) ExpandPuzzle() {
	var tiles [][]TileState

	// TODO: remove this
	tiles = s.tiles

	// Expand the puzzle to include one extra column per column
	// Expand the puzzle to include one extra row per row
	// Make sure we have a border all around so the next algo can fill around

	s.tiles = tiles
}

// FloodFill flood fills the puzzle marking all cells it can find as being outside.
func (s *Solver) FloodFill() {

}

// setup map with the walls part of the path only (and substitute start position with correct pipe)
func (s *Solver) SetTilesForWallsOnly(initialTile PositionDirection, startTileType TileType) {
	currentPosition := initialTile.Position
	nextDirection := initialTile.Out
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
}

// FindTileCandidates finds possible tile candidates.
// This function should only ever return 2 candidates.
// It explores the 4 cells around the starting cell and returns the connections pipes.
func (s *Solver) FindTileCandidates(startPosition Position) (posDirections []PositionDirection, startTileType TileType) {
	var candidates []PositionDirection
	var connections []Direction

	// check cell above
	candidate, ok := DetermineCandidate(s.puzzle, startPosition.x, startPosition.y-1, DirectionNorth)
	if ok {
		candidates = append(candidates, candidate)
		connections = append(connections, DirectionNorth) // start tile must be pointing north
	}

	// check cell below
	candidate, ok = DetermineCandidate(s.puzzle, startPosition.x, startPosition.y+1, DirectionSouth)
	if ok {
		candidates = append(candidates, candidate)
		connections = append(connections, DirectionSouth) // start tile must be pointing south
	}

	// check cell to the left
	candidate, ok = DetermineCandidate(s.puzzle, startPosition.x-1, startPosition.y, DirectionWest)
	if ok {
		candidates = append(candidates, candidate)
		connections = append(connections, DirectionWest) // start tile must be pointing west
	}

	// check cell to the right
	candidate, ok = DetermineCandidate(s.puzzle, startPosition.x+1, startPosition.y, DirectionEast)
	if ok {
		candidates = append(candidates, candidate)
		connections = append(connections, DirectionEast) // start tile must be pointing east
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
