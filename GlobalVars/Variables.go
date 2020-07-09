package GlobalVars

const (
	N          int = 15
	S_Computer int = 1
)

var (
	CurrTurn  int  = 0
	FirstTurn int  = 0
	IsEnd     bool = false
	Board     [N+2][N+2]int
	Dx        [8]int = [8]int{1, 1, 0, -1, -1, -1,  0, 1 }
	Dy        [8]int = [8]int{0, 1, 1,  1,  0, -1, -1,-1 }
)
