package lib

type Cell struct {
	Value string
}

// Checks whether the cell is occupied by the provided string
func (c Cell) IsOccupiedBy(occ string) bool {
	return c.Value == occ
}

// Checks whether the cell is occupied at all
func (c Cell) IsOccupied() bool {
	return len(c.Value) != 0
}

// Returns a string representing the current value of the cell
func (c Cell) String() string {
	return c.Value
}
