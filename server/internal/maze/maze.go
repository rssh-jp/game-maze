package maze

import (
	"fmt"

	"github.com/rssh-jp/game-maze/server/internal/maze/maze"
)

type Maze struct {
	w     int
	h     int
	block [][]Block
}

func New(w, h int) *Maze {
	block := make([][]Block, 0, h)

	for i := 0; i < h; i++ {
		wb := make([]Block, 0, w)
		for k := 0; k < w; k++ {
			wb = append(wb, NewBlock(BlockWall))
		}
		block = append(block, wb)
	}

	block[8][8] = NewBlock(BlockRoad)

	maze.CreateMaze(w, h)

	return &Maze{
		w:     w,
		h:     h,
		block: block,
	}
}

func (m *Maze) Print() {
	for _, itemh := range m.block {
		for _, itemw := range itemh {
			var str string
			switch itemw {
			case BlockWall:
				str = "■"
			case BlockRoad:
				str = "  "
			}
			fmt.Printf("%s", str)
		}
		fmt.Println()
	}
}
