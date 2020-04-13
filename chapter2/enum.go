package main

import "fmt"

func main(){
  const (
    A = iota + 3  //3부터 시작
    B
    _ //C         //_를 쓰면 생략
    D
    E
  )

  fmt.Println("A:", A)
  fmt.Println("B:", B)
  //fmt.Println("C:", C)
  fmt.Println("D:", D)
  fmt.Println("E:", E)

  const (
    _ = iota + 0.75*2   //시작값이 1.5이지만 _이므로 생략
    DEFAULT             //2.5부터 시작
    SILVER
    GOLD
    PLATINUM
  )

  fmt.Println("DEFAULT:", DEFAULT)
  fmt.Println("SILVER:", SILVER)
  fmt.Println("GOLD:", GOLD)
  fmt.Println("PLATINUM:", PLATINUM)
}
