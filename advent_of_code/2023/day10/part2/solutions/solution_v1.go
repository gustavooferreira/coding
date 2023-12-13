package solutions

import (
	"fmt"
	"strings"
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
	line = strings.TrimSpace(line)
	if line == "" {
		return

	}

	s.inputContent = append(s.inputContent, line)
}

// ComputeResult computes the result and stores it in the result field.
func (s *Solver) ComputeResult() {
	if s.debug {
		fmt.Println("%%%%%%%%%%")
		fmt.Println("Input:")
		for _, line := range s.inputContent {
			fmt.Println(line)
		}
		fmt.Println("%%%%%%%%%%")
	}

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
		fmt.Println("%%%%%%%%%%")
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
		fmt.Println("%%%%%%%%%%")
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
	s.FloodFill()

	// Count how many tiles are of type cell and are inside the loop
	result := 0
	for _, line := range s.tiles {
		for _, tile := range line {
			if tile.TileContainerType == TileContainerType_Cell && tile.TileContainerLocation == TileContainerLocation_Inside {
				result++
			}
		}
	}

	if s.debug {
		fmt.Println("%%%%%%%%%%")
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
		fmt.Println("%%%%%%%%%%")
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

func (s *Solver) ExpandPuzzle() {
	var tiles [][]TileState

	// create first row with the right size
	row := make([]TileState, len(s.tiles[0])*2+1)
	for i, c := range row {
		c.TileContainerType = TileContainerType_Cell
		c.TileContainerLocation = TileContainerLocation_Outside
		row[i] = c
	}
	tiles = append(tiles, row)

	for i, tilesRow := range s.tiles {
		// put normal row -------------------------------
		row := make([]TileState, len(s.tiles[i])*2+1)

		// create first empty cell
		row[0] = TileState{
			TileContainerType:     TileContainerType_Cell,
			TileContainerLocation: TileContainerLocation_Outside,
		}

		for j, tileState := range tilesRow {
			// put normal cell
			row[1+j*2] = tileState

			// put spacing cell
			if tileState.TileContainerType == TileContainerType_Wall && TileHasDirection(tileState.WallType, DirectionEast) {
				row[1+j*2+1] = TileState{
					TileContainerType: TileContainerType_Wall,
					WallType:          TileTypeHorizontal,
				}
			} else {
				row[1+j*2+1] = TileState{
					TileContainerType:     TileContainerType_Cell,
					TileContainerLocation: TileContainerLocation_Outside,
				}
			}
		}

		tiles = append(tiles, row)
		// ----------------------------------------------

		// put spacing row ------------------------------
		spacingRow := make([]TileState, len(s.tiles[i])*2+1)

		// create first empty cell
		spacingRow[0] = TileState{
			TileContainerType:     TileContainerType_Cell,
			TileContainerLocation: TileContainerLocation_Outside,
		}

		// put spacing row
		for j, tileState := range tilesRow {
			// put vertical spacing cell
			if tileState.TileContainerType == TileContainerType_Wall && TileHasDirection(tileState.WallType, DirectionSouth) {
				spacingRow[1+j*2] = TileState{
					TileContainerType: TileContainerType_Wall,
					WallType:          TileTypeVertical,
				}
			} else {
				spacingRow[1+j*2] = TileState{
					TileContainerType:     TileContainerType_Cell,
					TileContainerLocation: TileContainerLocation_Outside,
				}
			}

			// put spacing cell
			spacingRow[1+j*2+1] = TileState{
				TileContainerType:     TileContainerType_Cell,
				TileContainerLocation: TileContainerLocation_Outside,
			}
		}

		tiles = append(tiles, spacingRow)
		// ----------------------------------------------
	}

	s.tiles = tiles
}

// FloodFill flood fills the puzzle marking all cells it can find as being outside.
func (s *Solver) FloodFill() {
	queue := SimplePositionsQueue{}
	visited := make(map[Position]struct{})

	// start at 0,0 coordinates
	queue.Enqueue(Position{x: 0, y: 0})

	for queue.Len() > 0 {
		pos := queue.Dequeue()
		visited[pos] = struct{}{}

		// mark cell as being outside
		s.tiles[pos.y][pos.x] = TileState{
			TileContainerType:     TileContainerType_Cell,
			TileContainerLocation: TileContainerLocation_Outside,
		}

		// loop over neighbors
		for _, pos := range s.FindValidNeighbors(pos) {
			if _, ok := visited[pos]; !ok {
				queue.Enqueue(pos)
			}
		}
	}
}

func (s *Solver) FindValidNeighbors(tilePos Position) []Position {
	abovePos := Position{x: tilePos.x, y: tilePos.y - 1}
	belowPos := Position{x: tilePos.x, y: tilePos.y + 1}
	leftPos := Position{x: tilePos.x - 1, y: tilePos.y}
	rightPos := Position{x: tilePos.x + 1, y: tilePos.y}

	positions := []Position{abovePos, belowPos, leftPos, rightPos}
	var validNeighbors []Position

	for _, pos := range positions {
		if !IndexesAreWithinBounds(s.tiles, pos) {
			continue
		}
		if s.tiles[pos.y][pos.x].TileContainerType == TileContainerType_Cell {
			validNeighbors = append(validNeighbors, pos)
		}
	}

	return validNeighbors
}

func IndexesAreWithinBounds(tiles [][]TileState, pos Position) bool {
	return pos.y >= 0 && pos.y < len(tiles) && pos.x >= 0 && pos.x < len(tiles[0])
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
