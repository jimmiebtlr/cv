package cv

import (
  "fmt"
  "testing"
)

func TestCircleOffset(t *testing.T) {
  for i := 0; i < 9; i++ {
    x, y := CircleOffset(i, 3)
    fmt.Print(x)
    fmt.Print(" ")
    fmt.Println(y)
  }
}
