package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  a := -7
  switch {
  case a < 0:
    fmt.Println("a는 음수")
  case a > 0:
    fmt.Println("a는 양수")
  default:
    fmt.Println("a는 0")
  }

  switch b := 27; {
  case b < 0:
    fmt.Println("b는 음수")
  case b > 0:
    fmt.Println("b는 양수")
  default:
    fmt.Println("b는 0")
  }

  switch c := "go"; c {
  case "go":
    fmt.Println("go")
  case "gorani":
    fmt.Println("gorani")
  default:
    fmt.Println("unknown")
  }

  switch d := "go"; d + "lang" {
  case "go":
    fmt.Println("go")
  case "golang":
    fmt.Println("golang")
  default:
    fmt.Println("unknown")
  }

  switch i, j := 20, 30; {
  case i < j:
    fmt.Println("i < j")
  case i > j:
    fmt.Println("i > j")
  default:
    fmt.Println("i == j")
  }

  rand.Seed(time.Now().UnixNano())
  switch k := rand.Intn(100); {
  case k >= 25 && k < 50:
    fmt.Println("k는 25 이상 50 미만")
  case k >= 50:
    fmt.Println("k는 50 이상")
  default:
    fmt.Println("k는 25 이하")
  }

  z := 30/15
  switch z {
  case 2, 4, 6:
    fmt.Println("짝수")
  case 1, 3, 5:
    fmt.Println("홀수")
  }

  switch e := "go"; e {
  case "go":
    fmt.Println("go")
    fallthrough
  case "java":
    fmt.Println("java")
    fallthrough
  case "python":
    fmt.Println("python")
    //fallthrough는 마지막 case에는 못 쓴다
  }
}
