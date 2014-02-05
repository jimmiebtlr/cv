package cv

import (
  "fmt"
  "testing"
)

func TestDiamondOffset(t *testing.T) {
  //perimeter size should always be

  for i := 0; i < 8; i++ {
    x, y := DiamondOffset(i, 2)
    fmt.Print(x)
    fmt.Print(" ")
    fmt.Println(y)
  }
}
