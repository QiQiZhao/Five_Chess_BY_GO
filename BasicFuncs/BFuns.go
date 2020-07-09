package BasicFuncs

import (
	"fmt"
	"time"
)

import . "FiveChess/GlobalVars"
import . "FiveChess/ChessTypeJudge"
import . "FiveChess/DrawFunc"

func Num2Bool(num int) bool{
	if num==0{
		return false
	}else{
		return true
	}
}

func GameEnd(row,col int) bool{
	var u int = 0
	for u=0;u<4;u++{
		if CountNum(row, col, u) + CountNum(row, col, u + 4) >= 4{
			IsEnd =true
		}
	}
	if IsEnd ==true{
		return true
	}
	IsEnd =DisAllowedMove(row,col)
	return IsEnd
}

func UnderPawn(row,col int){
	if CurrTurn==FirstTurn{
		Board[row][col] = -1
	}else{
		Board[row][col] = -2
	}
	for i:=0;i<=N;i++{
		for j:=0;j<=N;j++{
			if i==row && j==col{
				continue
			}
			if Board[i][j] < 0{
				Board[i][j] *= -1
			}
		}
	}
	DrawBroad()
	if DisAllowedMove(row,col){
		if FirstTurn== S_Computer {
			fmt.Printf("Winner is You.")
		}else{
			fmt.Printf("Winner is Computer.")
		}
		time.Sleep(10000)
	}
}
