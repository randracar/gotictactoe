package main

import (
	"fmt"
)

var running bool = true

/* [0, 0, 0]
   [0, 0, 0]
   [0, 0, 0]
*/


func check_win_p1(array [3][3]int, nick string) bool {
	if array[0][0] == 1 && array[1][1] == 1 && array[2][2] == 1 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[2][0] == 1 && array[1][1] == 1 && array[0][2] == 1 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[0][0] == 1 && array[0][1] == 1 && array[0][2] == 1 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[1][0] == 1 && array[1][1] == 1 && array[1][2] == 1 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[2][0] == 1 && array[2][1] == 1 && array[2][2] == 1 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[0][0] == 1 && array[1][0] == 1 && array[2][0] == 1 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[0][1] == 1 && array[1][1] == 1 && array[2][1] == 1 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[0][2] == 1 && array[1][2] == 1 && array[2][2] == 1 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else {
		return false
	}
}

func check_win_p2(array [3][3]int, nick string) bool {
	if array[0][0] == 2 && array[1][1] == 2 && array[2][2] == 2 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[2][0] == 2 && array[1][1] == 2 && array[0][2] == 2 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[0][0] == 2 && array[0][1] == 2 && array[0][2] == 2 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[1][0] == 2 && array[1][1] == 2 && array[1][2] == 2 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[2][0] == 2 && array[2][1] == 2 && array[2][2] == 2 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[0][0] == 2 && array[1][0] == 2 && array[2][0] == 2 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[0][1] == 2 && array[1][1] == 2 && array[2][1] == 2 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else if array[0][2] == 2 && array[1][2] == 2 && array[2][2] == 2 {
		fmt.Printf("%s won! Good game!\n", nick)
		return true
	} else {
		return false
	}
}

func update_array_p1(array [3][3]int, row int, column int) [3][3]int {
	array[row][column] = 1

	return array
}

func update_array_p2(array [3][3]int, row int, column int) [3][3]int {
	array[row][column] = 2

	return array
}

func main(){
	var game_array[3][3] int

	var player1_nick string
	var player2_nick string
	var played_row int
	var played_column int

	fmt.Println("Hello! Welcome to randracar's tic tac toe game in go!")

	fmt.Print("Player1, please enter your name/nickname: ")
	fmt.Scanln(&player1_nick)
	fmt.Println("")

	fmt.Print("Player2, please enter your name/nickname: ")
	fmt.Scanln(&player2_nick)
	fmt.Println("")

	fmt.Printf("Welcome %s and %s! The game is simple: for %s, we use the number 1 to signal where you played, and for %s we use the number 2 to signal it instead. Good game!\n", player1_nick, player2_nick, player1_nick, player2_nick)

	for running == true {

		// player 1 is playing here

	player1turn:
		for i := 0; i < len(game_array); i++ {
			fmt.Println(i+1," ", game_array[i])
		}

		fmt.Printf("%s, please input the row you want to play in: ", player1_nick)
		fmt.Scanln(&played_row)
		played_row--
		fmt.Println("")

		fmt.Printf("%s, please input the column you want to play in: ", player1_nick)
		fmt.Scanln(&played_column)
		played_column--
		fmt.Println("")

		if game_array[played_row][played_column] != 0 || played_row > 2 || played_row < 0 || played_column > 2 || played_column < 0 {
			fmt.Println("Invalid choice! Please select a valid choice.")
			goto player1turn
		}

		game_array = update_array_p1(game_array, played_row, played_column)

		var checkwin bool = check_win_p1(game_array, player1_nick)
		
		if checkwin == true {
			break;
		} 

		// player2 is playing here
	
	player2turn:
		for i := 0; i < len(game_array); i++ {
			fmt.Println(i+1," ", game_array[i])
		}

		fmt.Printf("%s, please input the row you want to play in: ", player2_nick)
		fmt.Scanln(&played_row)
		played_row--
		fmt.Println("")

		fmt.Printf("%s, please input the column you want to play in: ", player2_nick)
		fmt.Scanln(&played_column)
		played_column--
		fmt.Println("")
		
		if game_array[played_row][played_column] != 0 || played_row > 2 || played_row < 0 || played_column > 2 || played_column < 0 {
			fmt.Println("Invalid choice! Please select a valid choice.")
			goto player2turn
		}

		game_array = update_array_p2(game_array, played_row, played_column)

		checkwin = check_win_p2(game_array, player2_nick)

		if checkwin == true {
			break;
		}
		
	}
}

