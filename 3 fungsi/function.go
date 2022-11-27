package main

import "fmt"


func main() {
    var a = 20
    var b = 4
    jumlah(a, b)
    kurang(a, b)
    kali(a, b)
    bagi(a, b)
}

func jumlah(a int, b int){
    fmt.Println(a, "+", b, "=", a+b)
}

func kurang(a int, b int){
    fmt.Println(a, "-", b, "=", a-b)
}

func kali(a int, b int){
    fmt.Println(a, "*", b, "=", a*b)
}

func bagi(a int, b int){
    fmt.Println(a, "/", b, "=", a/b)
}