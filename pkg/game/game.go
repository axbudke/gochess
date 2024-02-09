package game

import (
	"fmt"
	"gochess/pkg/generation"
	"gochess/pkg/notation"

	"github.com/manifoldco/promptui"
)

func GameLoop() {
	boardPosition, err := notation.NewPosition(notation.StartingFEN)
	if err != nil {
		fmt.Println("invalid starting position")
		return
	}

	for {
		// Display position
		fmt.Println(boardPosition.AsciiString())

		// Display Moves
		moves := generation.GenerateMoves(boardPosition)
		fmt.Println(fmt.Sprint("Legal Moves: ", moves))

		// Select Move
		sel := promptui.Select{
			Label: "Move?",
			Items: moves,
		}
		i, _, err := sel.Run()
		if err == promptui.ErrInterrupt {
			break
		} else if err != nil {
			fmt.Println("select failed: ", err)
			continue
		}

		// Make Move
		boardPosition = generation.MakeMove(boardPosition, *moves[i])
	}
}
