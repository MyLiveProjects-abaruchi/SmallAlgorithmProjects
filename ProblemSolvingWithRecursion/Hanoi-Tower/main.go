package main

import "fmt"

const num_disks = 3

func push(post []int, disk int) []int {
    return append([]int{disk}, post...)
}

func pop(post []int) (int, []int) {
    return post[0], post[1:]
}

func move_disk(posts [][]int, from_post, to_post int) {
    var disk int
    disk, posts[from_post] = pop(posts[from_post])
    posts[to_post] = push(posts[to_post], disk)
}

func draw_posts(posts [][]int) {
    for p := 0; p < 3; p++ {
        for len(posts[p]) < num_disks {
            posts[p] = push(posts[p], 0)
        }
    }

    for row := 0; row < num_disks; row++ {
        for p := 0; p < 3; p++ {
            fmt.Printf("%d ", posts[p][row])
        }
        fmt.Println()
    }
    fmt.Println("-----")

    for p := 0; p < 3; p++ {
        for len(posts[p]) > 0 && posts[p][0] == 0 {
            _, posts[p] = pop(posts[p])
        }
    }
}

func move_disks(posts [][]int, num_to_move, from_post, to_post, temp_post int) {
    if num_to_move > 0 {
        move_disks(posts, num_to_move-1, from_post, temp_post, to_post)
        move_disk(posts, from_post, to_post)
        draw_posts(posts)
        move_disks(posts, num_to_move-1, temp_post, to_post, from_post)
    }
}

func main() {
    // Make three posts.
    posts := [][]int{}

    // Push the disks onto post 0 biggest first.
    posts = append(posts, []int{})
    for disk := num_disks; disk > 0; disk-- {
        posts[0] = push(posts[0], disk)
    }

    // Make the other posts empty.
    for p := 1; p < 3; p++ {
        posts = append(posts, []int{})
    }

    // Draw the initial setup.
    draw_posts(posts)

    // Move the disks.
    move_disks(posts, num_disks, 0, 1, 2)
}
