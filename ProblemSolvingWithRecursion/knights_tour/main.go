package main

import (
    "fmt"
    "strconv"
    "time"
)

// The board dimensions.
const num_rows = 8
const num_cols = num_rows

// Whether we want an open or closed tour.
const require_closed_tour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
    dr, dc int
}

var move_offsets []Offset

var num_calls int64

func initialize_offsets() {
    move_offsets = []Offset{
        Offset{-2, -1},
        Offset{-1, -2},
        Offset{+2, -1},
        Offset{+1, -2},
        Offset{-2, +1},
        Offset{-1, +2},
        Offset{+2, +1},
        Offset{+1, +2},
    }
}

func make_board(num_rows, num_cols int) [][]int {
    board := make([][]int, num_rows)
    for i := range board {
        board[i] = make([]int, num_cols)
        for j := range board[i] {
            board[i][j] = unvisited
        }
    }

    return board
}

func dump_board(board [][]int) {
    var print_val string

    for i := range board {
        for j := range board[i] {
            if board[i][j] >= 0 && board[i][j] <= 9 {
                print_val = "0" + strconv.Itoa(board[i][j])
            } else {
                print_val = strconv.Itoa(board[i][j])
            }
            fmt.Printf("%s ", print_val)
        }
        fmt.Printf("\n")
    }
}

// Try to extend a knight's tour starting at (start_row, start_col).
// Return true or false to indicate whether we have found a solution.
func find_tour(board [][]int, num_rows, num_cols, cur_row, cur_col, num_visited int) bool {

    num_calls += 1

    if num_visited == num_rows*num_cols {
        if !require_closed_tour {
            return true
        } else {
            for _, offset := range move_offsets {
                r := cur_row + offset.dr
                c := cur_col + offset.dc
                if r < 0 || r >= num_rows {
                    continue
                }
                if c < 0 || c >= num_cols {
                    continue
                }
                if board[r][c] != unvisited {
                    continue
                }
                if board[r][c] == 0 {
                    return true
                }
            }
            return false
        }

    } else {
        for _, offset := range move_offsets {
            r := cur_row + offset.dr
            c := cur_col + offset.dc
            if r < 0 || r >= num_rows {
                continue
            }
            if c < 0 || c >= num_cols {
                continue
            }
            if board[r][c] != unvisited {
                continue
            }

            board[r][c] = num_visited
            if find_tour(board, num_rows, num_cols, r, c, num_visited+1) {
                return true
            }
            board[r][c] = unvisited
        }
    }
    return false
}

func main() {
    num_calls = 0

    // Initialize the move offsets.
    initialize_offsets()

    // Create the blank board.
    board := make_board(num_rows, num_cols)

    // Try to find a tour.
    start := time.Now()
    board[0][0] = 0
    if find_tour(board, num_rows, num_cols, 0, 0, 1) {
        fmt.Println("Success!")
    } else {
        fmt.Println("Could not find a tour.")
    }
    elapsed := time.Since(start)
    dump_board(board)
    fmt.Printf("%f seconds\n", elapsed.Seconds())
    fmt.Printf("%d calls\n", num_calls)
}
