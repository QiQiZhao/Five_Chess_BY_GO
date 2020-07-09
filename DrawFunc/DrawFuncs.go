package DrawFunc

import (
	"fmt"
	"os"
	"os/exec"
)

import . "FiveChess/GlobalVars"

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
	if S_Computer ==FirstTurn{
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
