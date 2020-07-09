package ChessTypeJudge

import . "FiveChess/GlobalVars"

func OpenFour(row,col int) int{
	var(
		key int = Board[row][col]
		sum int = 0
		i   int = 0
		u   int = 0
		sumNum int =0
	)
	for u=0;u<4;u++{
		sumNum=1
		for i=1;IsSame(row + Dx[u] * i, col + Dy[u] * i, key);i++{
			sumNum+=1
		}
		if !InBoard(row + Dx[u] * i, col + Dy[u] * i) || Board[row + Dx[u] * i][col + Dy[u] * i] != 0{
			continue
		}
		for i=-1;IsSame(row + Dx[u] * i, col + Dy[u] * i, key);i--{
			sumNum+=1
		}
		if !InBoard(row + Dx[u] * i, col + Dy[u] * i) || Board[row + Dx[u] * i][col + Dy[u] * i] != 0{
			continue
		}
		if 4==sumNum{
			sum+=1
		}
	}
	return sum
}
