package xcode

import "fmt"


func InputInt() int {
    var tmp int
    fmt.Scanln(&tmp)
    return tmp
}

func InputString() string {
    var tmp string
    fmt.Scanln(&tmp)
    return tmp
}


func InputIntArray(num int) []int {
    var result []int
    var tmp int
    for i := 0; i < num; i++ {
        fmt.Scan(&tmp)
        result = append(result, tmp)
    }
    return result
}

