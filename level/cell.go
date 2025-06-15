package level

type CellType int

const (
	CellTypeEmpty CellType = iota
	CellTypeFloor
	CellTypeWall
)

type Cell struct {
	Type CellType
}
