package main

import (
    "fmt"
    "time"
)

func make_board(num_rows int) [][]string {
    board := make([][]string, num_rows)
    for i := range board {
        board[i] = make([]string, num_rows)
        for j := range board[i] {
            board[i][j] = "."
        }
    }

    return board
}

func dump_board(board [][]string) {
    for i := range board {
        for j := range board[i] {
            fmt.Printf("%s ", board[i][j])
        }
        fmt.Printf("\n")
    }
}

func series_is_legal(board [][]string, num_rows, r0, c0, dr, dc int) bool {
    nr_queens := 0

    for ; r0 < num_rows && c0 < num_rows && r0 >= 0 && c0 >= 0; {
        if board[r0][c0] != "." {
            nr_queens += 1
            if nr_queens >= 2 {
                return false
            }
        }
        r0 += dr
        c0 += dc
    }

    return true
}

func board_is_legal(board [][]string, num_rows int) bool {
    // See if each row is legal.
    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, 0, 0, 1) {
            return false
        }
    }

    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, 0) {
            return false
        }
    }

    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, 0, 1, 1) {
            return false
        }
    }
    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, 1) {
            return false
        }
    }

    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, num_rows-1, 1, -1) {
            return false
        }
    }
    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, -1) {
            return false
        }
    }

    return true
}

// Return true if the board is legal and a solution.
func board_is_a_solution(board [][]string, num_rows int) bool {

    if board_is_legal(board, num_rows) {
        nr_queens := 0

        for i := 0; i < num_rows; i++ {
            for j := 0; j < num_rows; j++ {
                if board[i][j] != "." {
                    nr_queens++
                }
            }
        }
        return nr_queens == num_rows
    }

    return false
}

func place_queens_1(board [][]string, num_rows, r, c int) bool {

    var next_col, next_row int

    if r >= num_rows {
        return board_is_a_solution(board, num_rows)
    }

    next_col = c + 1
    next_row = r
    if next_col >= num_rows {
        next_row += 1
        next_col = 0
    }

    // Eval if we do not set the Queen here
    if place_queens_1(board, num_rows, next_row, next_col) {
        return true
    }

    // Eval if we set the Queen here
    board[r][c] = "Q"
    if place_queens_1(board, num_rows, next_row, next_col) {
        return true
    }
    board[r][c] = "."
    return false
}

func main() {
    const num_rows = 6
    board := make_board(num_rows)

    start := time.Now()
    success := place_queens_1(board, num_rows, 0, 0)
    //success := place_queens_2(board, num_rows, 0, 0, 0)
    //success := place_queens_3(board, num_rows, 0, 0, 0)

    elapsed := time.Since(start)
    if success {
        fmt.Println("Success!")
        dump_board(board)
    } else {
        fmt.Println("No solution")
    }
    fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
