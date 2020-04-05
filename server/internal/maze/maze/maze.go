package maze

import (
	"fmt"
	"math/rand"
)

func CreateMaze(w, h int) [][]byte {
	m := make([][]byte, 0, h)
	for i := 0; i < h; i++ {
		work := make([]byte, 0, w)
		for k := 0; k < w; k++ {
			work = append(work, 0)
		}
		m = append(m, work)
	}

	p := getStarPos(w, h)
	set1(m, p)

	Print(m)

	count := 1000
	for loop := true; loop; {
		p = func() *Pos {
			valids := findValidPos(m, p)

			list := make([]*Pos, 0, len(valids))
			for _, item := range valids {
				if item == nil {
					continue
				}

				list = append(list, item)
			}

			if len(list) == 0 {
				return nil
			}

			return list[rand.Intn(len(list))]
		}()

		if p == nil {
			if count == 0 {
				loop = false
				break
			}

			p = func() *Pos {
				list := listValid(m)

				if len(list) == 0 {
					return nil
				}

				return list[rand.Intn(len(list))]
			}()

			if p == nil {
				loop = false
				break
			}

			count--
		}

		set1(m, p)
		Print(m)

	}

	list := listValid(m)

	set2(m, list[rand.Intn(len(list))])
	set3(m, list[rand.Intn(len(list))])

	Print(m)

	return m
}

func Print(m [][]byte) {
	fmt.Println("---------------------------------------------------------------------------------")
	for _, itemh := range m {
		for _, itemw := range itemh {
			var str string
			switch itemw {
			case 0:
				str = "■"
			case 1:
				str = "  "
			case 2:
				str = "ス"
			case 3:
				str = "ゴ"
			}
			fmt.Printf("%s", str)
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------------------------------------------------")
}

type Pos struct {
	x int
	y int
}

func listValid(m [][]byte) []*Pos {
	list := make([]*Pos, 0, 8)
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] == 1 {
				list = append(list, &Pos{x: x, y: y})
			}
		}
	}

	return list
}

func set1(m [][]byte, p *Pos) {
	m[p.y][p.x] = 1
}
func set2(m [][]byte, p *Pos) {
	m[p.y][p.x] = 2
}
func set3(m [][]byte, p *Pos) {
	m[p.y][p.x] = 3
}

func getStarPos(w, h int) *Pos {
	return &Pos{
		x: rand.Intn(w),
		y: rand.Intn(h),
	}
}

// 0: left, 1: right, 2: top, 3: bottom
func findValidPos(m [][]byte, p *Pos) (ret [4]*Pos) {
	w := len(m[0])
	h := len(m)

	// left
	if p.x-1 >= 0 && m[p.y][p.x-1] == 0 {
		ret[0] = &Pos{x: p.x - 1, y: p.y}
	}

	// right
	if p.x+1 < w && m[p.y][p.x+1] == 0 {
		ret[1] = &Pos{x: p.x + 1, y: p.y}
	}

	// top
	if p.y-1 >= 0 && m[p.y-1][p.x] == 0 {
		ret[2] = &Pos{x: p.x, y: p.y - 1}
	}

	// bottom
	if p.y+1 < h && m[p.y+1][p.x] == 0 {
		ret[3] = &Pos{x: p.x, y: p.y + 1}
	}

	for index, item := range ret {
		if item == nil {
			continue
		}
		var count int

		if item.x-1 >= 0 && m[item.y][item.x-1] == 1 {
			count++
		}

		if item.x+1 < w && m[item.y][item.x+1] == 1 {
			count++
		}

		if item.y-1 >= 0 && m[item.y-1][item.x] == 1 {
			count++
		}

		if item.y+1 < h && m[item.y+1][item.x] == 1 {
			count++
		}

		if count < 2 {
			continue
		}

		ret[index] = nil
	}

	return
}
