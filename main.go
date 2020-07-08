package main

import (
	"fmt"
	"os"
	"time"
)
import  "os/exec"

// Global values

const (
	N int = 15
	s_Computer int= 1
)

var (
	CurrTurn  int       = 0
	FirstTurn int       = 0
	isEnd     bool      = false
	Board     [N+2][N+2]int
	dx [8]int = [8]int{ 1, 1, 0, -1, -1, -1,  0, 1 }
	dy [8]int = [8]int{ 0, 1, 1,  1,  0, -1, -1,-1 }
)

func Num2Bool(num int) bool{
	if num==0{
		return false
	}else{
		return true
	}
}

func BoardInit() int{
	//define the background and word color of command
	//system("color 0f")
	fmt.Println("input 1 or 2 to choose\n 1.Computer is black and plays firstly\n 2.Player is black and plays firstly\n")
	fmt.Scanf("%d",&FirstTurn)
	if FirstTurn!=1 && FirstTurn!=2 {
		fmt.Println("Please input a vaild number!")
		return BoardInit()
	}
	//ensure who plays firstly
	CurrTurn = FirstTurn;
	//use space surrounds the board
	for i:=0;i<=N+1;i++{
		for j:=0;j<=N+1;j++{
			Board[i][j] = 0
		}
	}
	DrawBroad()
	return 0
}

func DrawBroad() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	//system("cls")
	var(
		row int=0
		col int=0
		keyr int=0
		keyc int =0
		alpha byte='A'
	)
	fmt.Printf("\n\n\n     ")
	for col=1;col<=N;col++{
		fmt.Printf("%c",alpha)
		alpha= 1 + alpha
		fmt.Printf(" ")
	}
	for row=N;row>=1;row--{
		fmt.Printf("\n   %2d", row)
		for col=1;col<=N;col++{
			PrintSymbol(row, col)
			if Board[row][col] < 0 {
				keyr = row;
				keyc = col;
			}
		}
		println(col)
	}
	alpha='A'
	fmt.Printf("\n     ")
	for col=1;col<=N;col++{
		fmt.Printf("%c",alpha)
		alpha+=1
		fmt.Printf(" ")
	}
	fmt.Println("\n\n")
	if s_Computer==FirstTurn{
		fmt.Printf("The computer is black, the player is white\n")
	}else{
		fmt.Printf("The computer is white, the player is black\n")
	}
	alpha = 'A'
	if keyr!=0{
		fmt.Printf("The final position is：")
		fmt.Printf("%c",byte(int(alpha) + keyc - 1))
		fmt.Printf("%d",keyr)
		fmt.Printf("\n")
	}

}

func PrintSymbol(i, j int) (int, error) {
	if Board[i][j] == 1{
		if j==N{
			return fmt.Printf("○")
		}else{
			return fmt.Printf("○-")
		}
	}
	if Board[i][j] == 2{
		if j==N{
			return fmt.Printf("●-")
		}else{
			return fmt.Printf("●")
		}
	}
	if Board[i][j] == -1{
		if j==N{
			return fmt.Printf("△")
		}else{
			return fmt.Printf("△-")
		}

	}
	if Board[i][j] == -2{
		if j==N{
			return fmt.Printf("▲")
		}else{
			return fmt.Printf("▲")
		}
	}
	if N==i{
		if 1==j{
			return fmt.Printf("┏─")
		}
		if N==j{
			return fmt.Printf("┓")
		}
		return fmt.Printf("┯─")
	}
	if 1==i{
		if 1==j{
			return fmt.Printf("┗—")
		}
		if N==j{
			return fmt.Printf("┛")
		}
		return fmt.Printf("┷—")
	}
	if 1==j{
		return fmt.Printf("┠─")
	}
	if N==j{
		return fmt.Printf("┨")
	}
	return fmt.Printf("┼─")

}

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
		i int = row+dx[u]
		j int = col+dy[u]
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
		i+=dx[u]
		j+=dy[u]
	}
	return sum
}

func openFour(row,col int) int{
	var(
		key int = Board[row][col]
		sum int = 0
		i   int = 0
		u   int = 0
		sumNum int =0
	)
	for u=0;u<4;u++{
		sumNum=1
		for i=1;IsSame(row + dx[u] * i, col + dy[u] * i, key);i++{
			sumNum+=1
		}
		if !InBoard(row + dx[u] * i, col + dy[u] * i) || Board[row + dx[u] * i][col + dy[u] * i] != 0{
			continue
		}
		for i=-1;IsSame(row + dx[u] * i, col + dy[u] * i, key);i--{
			sumNum+=1
		}
		if !InBoard(row + dx[u] * i, col + dy[u] * i) || Board[row + dx[u] * i][col + dy[u] * i] != 0{
			continue
		}
		if 4==sumNum{
			sum+=1
		}
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
	var flag bool = OverLine(row, col) || openFour(row, col)>1  //Need to Change
	return flag
}

func GameEnd(row,col int) bool{
	var u int = 0
	for u=0;u<4;u++{
		if CountNum(row, col, u) + CountNum(row, col, u + 4) >= 4{
			isEnd=true
		}
	}
	if isEnd==true{
		return true
	}
	isEnd=DisAllowedMove(row,col)
	return isEnd
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
		if FirstTurn==s_Computer{
			fmt.Printf("Winner is You.")
		}else{
			fmt.Printf("Winner is Computer.")
		}
		time.Sleep(10000)
	}
}

func ComputeScore(row,col int) int{
	if DisAllowedMove(row,col){
		return 0
	}
	if GameEnd(row,col){
		isEnd=false
		return 10000
	}
	var score int = openFour(row, col) * 1000
	for i:=0;i<8;i++{
		if Num2Bool(Board[row + dx[i]][col + dy[i]]){
			score+=1
		}
	}
	return score
}

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
		if isEnd{
			break
		}
		if CurrTurn==s_Computer{
			Max1()
		}else{
			Player()
		}
		CurrTurn=3-CurrTurn
	}
}