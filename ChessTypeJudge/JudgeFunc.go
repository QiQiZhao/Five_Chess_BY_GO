package ChessTypeJudge

import . "FiveChess/GlobalVars"

func InBoard(row,col int) bool{
	if row<1 || col>N{
		return false
	}
	return col>=1||col<=N
}

func IsOk(row,col int) bool{
	return InBoard(row,col)&&Board[row][col]==0
}

func IsSame(row,col,key int) bool{
	if !InBoard(row,col){
		return false
	}
	return Board[row][col] == key || Board[row][col] + key == 0
}

func CountNum(row,col,u int) int{
	var(
		i int = row+ Dx[u]
		j int = col+ Dy[u]
		sum = 0
		ref =Board[row][col]
	)
	if ref==0{
		return 0
	}
	for{
		if !IsSame(i,j,ref){    //Need to check
			break
		}
		sum+=1
		i+= Dx[u]
		j+= Dy[u]
	}
	return sum
}
func OverLine(row,col int) bool{
	var(
		flag bool = false
		u    int  = 0
	)
	for u=0;u<4;u++{
		if CountNum(row, col, u) + CountNum(row, col, u + 4) > 4{
			flag=true
		}
	}
	return flag
}

func DisAllowedMove(row,col int) bool{
	if IsSame(row,col,2){
		return false
	}
	var flag bool = OverLine(row, col) || OpenFour(row, col)>1 //Need to Change
	return flag
}
