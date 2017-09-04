package main

import (
    "fmt"

)
func swap(x, y string) (string, string){
 return y, x
}
func main() {
 a, b := swap("hello", "gopher")
 fmt.Printf("swap of a and b are %s and %s", a, b)
}
