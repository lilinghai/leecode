package backtrack

import "fmt"

var result []int // 全局或成员变量, 下标表示行, 值表示 queen 存储在哪一列
func init(){
	result=make([]int,8)
}
func Call8Queens(row int) { // 调用方式：cal8queens(0);
	if (row == 8) { // 8 个棋子都放置好了，打印结果
		printQueens(result);
		return; // 8 行棋子都放好了，已经没法再往下递归了，所以就 return //
	}
	for column:=0;column<8;column++{ // 每一行都有 8 中放法
		if isOk(row, column) { // 有些放法不满足要求
			result[row] = column // 第 row 行的棋子放到了 column 列
			Call8Queens(row+1) // 考察下一行
		}
	}
}

func isOk(row int, column int)bool  {// 判断 row 行 column 列放置是否合适
	leftup:= column - 1
	rightup := column + 1
	for i:= row-1; i >= 0; i-- { // 逐行往上考察每一行
		if result[i] == column {
			return false
		}// 第 i 行的 column 列有棋子吗？
		if leftup >= 0&&result[i]==leftup { // 考察左上对角线：第 i 行 leftup 列有棋子吗？
			return false
		}
	if rightup < 8 && result[i]==rightup{ // 考察右上对角线：第 i 行 rightup 列有棋子吗？
		return false
	}
		leftup--
		rightup++
	}
	return true;
}

func printQueens(result []int) { // 打印出一个二维矩阵
	for row := 0; row < 8; row++ {
		for  column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Print("Q ");
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

