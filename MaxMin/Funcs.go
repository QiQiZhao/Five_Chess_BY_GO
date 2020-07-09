package MaxMin

import . "FiveChess/ChessTypeJudge"
import . "FiveChess/GlobalVars"
import . "FiveChess/BasicFuncs"

func ComputeScore(row,col int) int{
	if DisAllowedMove(row,col){
		return 0
	}
	if GameEnd(row,col){
		IsEnd =false
		return 10000
	}
	var score int = OpenFour(row, col) * 1000
	for i:=0;i<8;i++{
		if Num2Bool(Board[row + Dx[i]][col + Dy[i]]){
			score+=1
		}
	}
	return score
}
