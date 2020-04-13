package main

import "fmt"

func main(){
  const a string = "test1"
  const b = "test2"
  const c int32 = 10 * 10

  fmt.Println("a:", a, " b:", b, " c:", c)

  const (
    x, y int16 = 50, 90
    i, k = "Data", 7776
  )
  fmt.Println("x:", x, " y:", y, " i:", i, " k:", k)
}
