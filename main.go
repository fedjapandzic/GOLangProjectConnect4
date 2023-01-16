package main

import (
	"math/rand"
	"time"

	"4-in-a-row-go/cpu"
	"4-in-a-row-go/game"
	"4-in-a-row-go/player"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	b := &game.Board{}
	b.DrawTitle()
	for {
		// init
		b.Init()
		b.DrawBoard()
		for {

			// player's turn
			player.ExecPlayerTurn(b)
			// draw bord
			b.DrawBoard()
			// check
			b.Judge()
			if b.GameStatus != game.Playing {
				break
			}
			// CPU's turn
			cpu.ExecCPUTurn(b)
			// draw bord
			b.DrawBoard()
			// check
			b.Judge()
			if b.GameStatus != game.Playing {
				break
			}
		}
		// game end
		b.EndGame()
	}
}
