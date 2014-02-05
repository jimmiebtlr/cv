package cv

import (
  "testing"
)

func TestDiamondOffset(t *testing.T) {
  //perimeter size should always be
  expected := [][]int{
    []int{3, 0},
    []int{0, 3},
    []int{-3, 0},
    []int{0, -3},
    []int{2, -1},
    []int{-1, 2},
    []int{-2, -1},
    []int{-1, -2},
    []int{2, 1},
    []int{1, 2},
    []int{-2, 1},
    []int{1, -2},
  }

  index := 0
  for i, v := range expected {
    actual_x, actual_y := DiamondOffset(i, 3)
    if v[0] != actual_x || v[1] != actual_y {
      t.Fatalf("expected: %i,%i + actual: %i,%i", v[0], v[1], actual_x, actual_y)
    }
    index++
  }
  if index != 12 {
    t.Fatal("there should be exactly 12 offsets generated")
  }
}
