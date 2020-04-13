package main

import "fmt"

func main(){
  sum, i := 0, 0

  for i <= 100 {
    fmt.Println("i:", i)
    sum += i
    i++
    //sum += i++는 안됨, 증감연산(i++)의 반환값이 없기 때문
  }
  fmt.Println("sum:", sum)

  for i := 0; i <= 10; i++ {
    fmt.Print("ex: ")
    fmt.Println(i)
  }


}
