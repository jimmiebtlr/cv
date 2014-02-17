package cv

import (
  "math"
  "strconv"
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
  in  string
  out uint
}{
  {"0", 0},
  {strconv.FormatUint(math.MaxUint32, 2), math.MaxUint32},
  {"100000", 1}, // binary "100000"
  {"101000", 5}, // binary "101000"
}

func TestMinBitsZero(t *testing.T) {
  for _, v := range minBitsZeroTests {
    in, _ := strconv.ParseUint(v.in, 2, 32)
    if min, _ := MinBits(uint(in)); min != v.out {
      t.Logf("Expected %b actual %b", v.out, min)
      t.Fail()
    }
  }
}

var minBitsZeroRevTests = []struct {
  in       string
  out      uint
  reversed bool
}{
  {"1101000", 11, true}, // binary "1101000"
  {"1101000", 11, true}, // binary "1101000"
}

func TestMinBitsZeroRev(t *testing.T) {
  for _, v := range minBitsZeroRevTests {
    in, _ := strconv.ParseUint(v.in, 2, 32)
    if min, rev := MinBits(uint(in)); min != v.out || rev != v.reversed {
      t.Logf("Expected min = %b and reversed = %v actual min = %b and reversed = %v", v.out, v.reversed, min, rev)
      t.Fail()
    }
  }
}
