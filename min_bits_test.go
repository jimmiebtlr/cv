package cv

import (
  "math"
  "testing"
)

var maxZerosRunTests = []struct {
  in     string
  start  int
  length int
}{
  {"0", 0, 1},
  {"010", 2, 2},
  {"101000", 3, 3},
  {"1000101", 1, 3},
  {"1000101", 1, 3},
}

func TestMaxZerosRun(t *testing.T) {
  for _, v := range maxZerosRunTests {
    if start, length := maxZerosRun(v.in); start != v.start || length != v.length {
      t.Logf("Expected start = 0 and length = 1 actual start = %v and length = %v", start, length)
      t.Fail()
    }
  }
}

var minBitsZeroTests = []struct {
  in  uint
  out uint
}{
  {0, 0},
  {math.MaxInt32, math.MaxInt32},
  {32, 1},   // binary "100000"
  {40, 5},   // binary "101000"
  {104, 11}, // binary "1101000"
}

func TestMinBitsZero(t *testing.T) {
  for _, v := range minBitsZeroTests {
    if min, _ := MinBits(v.in); min != v.out {
      t.Logf("Expected %b actual %b", v.out, min)
      t.Fail()
    }
  }
}
