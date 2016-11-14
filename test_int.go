package main

import "fmt"

func main() {
    const LENGTH int = 10
    const WIDTH int = 5
    const PI float32 = 3.14
    const COUNT float32 = 20
    var area int
    var pi float32

    area = LENGTH * WIDTH
    pi = PI * COUNT
    fmt.Printf("value of area : %d\n", area)
    fmt.Printf("pi of area : %f\n", pi)
}
