package main

import "fmt"

func main() {
  var j string = "Hi GoLaNi"
  var l = "Hi GoGuMa"
  var i float32 = 11.4
  fmt.Println("j:", j)
  fmt.Println("l:", l)
  fmt.Println("i:", i)

  var(
    name string = "machine"
    height int32
    weight float32
    isRunning bool
  )

  height = 250
  weight = 140
  isRunning = true

  fmt.Println("name:", name, " height:", height, " weight:", weight, " isRunning:", isRunning)
}
