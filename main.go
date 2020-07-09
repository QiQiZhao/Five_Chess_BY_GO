package main

import (
	"fmt"
	"time"
)

import . "FiveChess/GlobalVars"
import . "FiveChess/ChessTypeJudge"
import . "FiveChess/DrawFunc"
import . "FiveChess/MaxMin"
import . "FiveChess/BasicFuncs"






func Player() int{
	DrawBroad()
	fmt.Printf("It is your turn，please input your position ： ")
	var(
		c byte = '\n'
		row int = 0
		col int = 0
	)
	for{
		if !(c<'0'){
			break
		}
		fmt.Scanf("%c",&c)
		fmt.Scanf("%d",&row)
	}

	if c<'a'{
		col=int(c)-int('A')+1
	}else{
		col=int(c)-int('a')+1
	}

	if !IsOk(row,col)||int(c)-int('A')<0||int(c)-int('z')>0{
		fmt.Printf("It is not a vaild position. Please input a vaild one.")
		time.Sleep(1000)
		Player()
		return 0
	}
	UnderPawn(row,col)
	return 0
}


func main() {
	BoardInit()
	for{
		if IsEnd {
			break
		}
		if CurrTurn== S_Computer {
			Max1()
		}else{
			Player()
		}
		CurrTurn=3-CurrTurn
	}
}