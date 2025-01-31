package variables

import "fmt"

func PrintVars() {
    var a = "a string"
    var b, c int = 1, 2
    var d = true
    var e int
    x := 1
    y, z := 2, 3

    fmt.Println(a, b, c, d, e)
    fmt.Println(e)
    fmt.Println(x, y, z)
}