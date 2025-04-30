package main

import "fmt"

func main(){
    var a, t float64

    fmt.Print("Masukkan alas segitiga : ")
    fmt.Scanln(&a)
    fmt.Print("Masukkan tinggi segitiga : ")
    fmt.Scanln(&t)

    luas := 0.5 * a * t
    kel := a * 3

    fmt.Println("Luas segitiga adalah ", luas)
    fmt.Println("Keliling segitiga adalah ", kel)
}
