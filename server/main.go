package main

import (
	"log"

	"github.com/rssh-jp/game-maze/server/internal/maze"
)

func main() {
	log.Println("START")
	defer log.Println("END")

	maze.New(20, 20)

	//m.Print()
}
