package main

import "fmt"

func main() {
    var test int = 27
    var res bool

    res = powerOfThree(test)

    fmt.Printf("result is : %t\n", res)
}

func powerOfThree(n int) bool {
    var ret bool

    n = n / 3

    fmt.Printf("Value is : %d\n", n)

    if (n == 1 || n == 0) {
        ret = true
    } else if (n < 0) {
        ret = false
    } else {
        powerOfThree(n)
    }

    return ret
}
