package solutions

// PositionDirection holds a position and the out direction for the next tile.
// The out direction represents the direction of the next move.
type PositionDirection struct {
	Position Position
	Out      Direction
}

// Position represents a tile's position.
type Position struct {
	x int
	y int
}

// TileType represents the tile type (walls, empty space and the starting point)
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

// Direction represents the direction of movement (cardinal direction).
type Direction string

const (
	DirectionNorth Direction = "north"
	DirectionSouth Direction = "south"
	DirectionEast  Direction = "east"
	DirectionWest  Direction = "west"
)

// Pipe represents a pipe and contains information about its terminals direction.
type Pipe struct {
	Terminal1 Direction
	Terminal2 Direction
}

// TileContainerType defines whether the tile contains a wall or a cell.
type TileContainerType string

const (
	TileContainerType_Wall TileContainerType = "wall"
	TileContainerType_Cell TileContainerType = "cell"
)

// TileState represents the tile state.
// Whether it's a cell or a wall.
// If it's a wall, it also contains information about what type of wall.
// If it's a cell, it also contains information about whether the cell is inside or outside the loop.
type TileState struct {
	TileContainerType     TileContainerType
	WallType              TileType
	TileContainerLocation TileContainerLocation
}

// TileContainerLocation represents whether a given cell is inside or outside the loop.
type TileContainerLocation string

const (
	TileContainerLocation_Inside  TileContainerLocation = "inside"
	TileContainerLocation_Outside TileContainerLocation = "outside"
)
