package maze

const (
	BlockWall = iota + 1
	BlockRoad
)

type Block byte

func NewBlock(typ byte) Block {
	switch typ {
	case BlockWall:
		return BlockWall
	case BlockRoad:
		return BlockRoad
	default:
		return BlockWall
	}
}
