package cv

import (
  "fmt"
  "testing"
)

func TestDiamondOffset(t *testing.T) {
  //perimeter size should always be

  for i := 0; i < 12; i++ {
    x, y := DiamondOffset(i, 3)
    fmt.Print(x)
    fmt.Print(" ")
    fmt.Println(y)
  }
}
