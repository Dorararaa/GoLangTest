package main

import "fmt"

func main(){
  var a int = 20
  if a >= 15 {
    fmt.Println("a는 15 이상")
  }

  b := 30
  if b >= 25 {
    fmt.Println("b는 25 이상")
  }

  //if b := 30; b >= 25 {
  //  fmt.Println("b는 25 이상")
  //}
  //위와 같은 문법

  if a < 15 {
    fmt.Println("a는 15 미만")
  } else if a > 30 {
    fmt.Println("a는 30 초과")
  } else {
    fmt.Println("a는 15 이상 30 이하")
  }

}
