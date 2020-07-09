package MaxMin

import . "FiveChess/DrawFunc"
import . "FiveChess/GlobalVars"
import . "FiveChess/BasicFuncs"
import . "FiveChess/ChessTypeJudge"

func Max1() int{
	DrawBroad()
	if Board[8][8]==0{
		UnderPawn(8,8)
		return 0
	}
	var(
		i int = 0
		j int = 0
		alpha int = -100000
		keyi int
		keyj int
	)
	for i=1;i<=N;i++{
		for j=1;j<=N;j++{
			if !IsOk(i,j){
				continue
			}
			Board[i][j]=FirstTurn
			var pScore int = ComputeScore(i, j)
			if pScore==0{
				Board[i][j]=0
				continue
			}
			if pScore==10000{
				UnderPawn(i,j)
				return 0
			}
			pScore=Min(alpha)
			Board[i][j]=0
			if pScore>alpha{
				alpha=pScore
				keyi=i
				keyj=j
			}
		}
	}
	UnderPawn(keyi,keyj)
	return 0
}

func Min(alpha int) int {
	var(
		i int
		j int
		beta int = 100000
	)
	for i=1;i<=N;i++{
		for j=1;j<=N;j++{
			if !IsOk(i,j){
				continue
			}
			Board[i][j]=3-FirstTurn
			var pScore=ComputeScore(i,j)
			if pScore==0{
				Board[i][j]=0
				continue
			}

			if pScore==10000{
				Board[i][j]=0
				return -10000
			}

			pScore=Max2(beta)
			Board[i][j]=0
			if pScore<beta{
				beta=pScore
			}
			if alpha>beta{
				return beta
			}
		}
	}
	return beta
}

func Max2(beta int) int {
	var(
		i int
		j int
		alpha int = -100000
	)
	for i=1;i<=N;i++{
		for j=1;j<=N;j++{
			if !IsOk(i,j){
				continue
			}
			Board[i][j]=FirstTurn
			var pScore int = ComputeScore(i,j)
			if pScore==0{
				Board[i][j]=0
				continue
			}
			if pScore==10000{
				Board[i][j]=0
				return 10000
			}
			Board[i][j]=0
			if pScore>alpha{
				alpha=pScore
			}
			if alpha>beta{
				return alpha
			}
		}
	}
	return alpha
}
