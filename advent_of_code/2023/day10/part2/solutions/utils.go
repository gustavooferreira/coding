package solutions

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

// FindNextDirection finds the next direction.
// Given a TileType and a direction, it finds the tile's other terminal and returns that direction.
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

// TileHasDirection checks whether the provided tile has one of the terminals pointing in the provided direction.
func TileHasDirection(tile TileType, direction Direction) bool {
	tileInfo := PipeInfoMap[tile]

	if tileInfo.Terminal1 == direction || tileInfo.Terminal2 == direction {
		return true
	}

	return false
}

// ReverseDirection reverses the direction.
// Example: If direction is North, it returns South.
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

// Given current position and a direction, return the next position.
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

// DetermineCandidate determines a possible cell candidate to the starting cell.
// Given a surrounding cell coordinates, determine if that cell is a candidate or not.
func DetermineCandidate(puzzle [][]rune, x int, y int, direction Direction) (candidate PositionDirection, ok bool) {
	if y < 0 || y >= len(puzzle) {
		return PositionDirection{}, false
	}

	if x < 0 || x >= len(puzzle[y]) {
		return PositionDirection{}, false
	}

	reverseDirection := ReverseDirection(direction)

	tileType := TileType(puzzle[y][x])
	tileInfo := PipeInfoMap[tileType]
	if tileInfo.Terminal1 == reverseDirection || tileInfo.Terminal2 == reverseDirection {
		outDirection := FindNextDirection(tileType, direction)
		return PositionDirection{
			Position: Position{
				x: x,
				y: y,
			},
			Out: outDirection,
		}, true
	}

	return PositionDirection{}, false
}

type SimplePositionsQueue struct {
	queue []Position
}

func (q *SimplePositionsQueue) Enqueue(pos Position) {
	q.queue = append(q.queue, pos)
}

func (q *SimplePositionsQueue) Dequeue() (pos Position) {
	pos = q.queue[0]
	q.queue = q.queue[1:]
	return pos
}

func (q *SimplePositionsQueue) Len() int {
	return len(q.queue)
}
