package main

import (
    "fmt"
    "strconv"
)

var fibonacci_values []int64

func fibonacci_on_the_fly(n int64) int64 {

    if n <= 1 {
        return fibonacci_values[n]
    }

    if int64(len(fibonacci_values)) > n {
        return fibonacci_values[n]
    }

    fibVal := fibonacci_on_the_fly(n-2) + fibonacci_on_the_fly(n-1)

    fibonacci_values = append(fibonacci_values, fibVal)
    return fibonacci_values[n]
}

func main() {
    // Fill-on-the-fly.
    fibonacci_values = make([]int64, 2)
    fibonacci_values[0] = 0
    fibonacci_values[1] = 1

    // Prefilled.
    //initialize_slice()

    for {
        // Get n as a string.
        var n_string string
        fmt.Printf("N: ")
        fmt.Scanln(&n_string)

        // If the n string is blank, break out of the loop.
        if len(n_string) == 0 {
            break
        }

        // Convert to int and calculate the Fibonacci number.
        n, _ := strconv.ParseInt(n_string, 10, 64)

        // Uncomment one of the following.
        fmt.Printf("fibonacci_on_the_fly(%d) = %d\n", n, fibonacci_on_the_fly(n))
        //fmt.Printf("fibonacci_prefilled(%d) = %d\n", n, fibonacci_prefilled(n))
        //fmt.Printf("fibonacci_bottom_up(%d) = %d\n", n, fibonacci_bottom_up(n))
    }

    // Print out all memoized values just so we can see them.
    for i := 0; i < len(fibonacci_values); i++ {
        fmt.Printf("%d: %d\n", i, fibonacci_values[i])
    }
}
