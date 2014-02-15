package cv

import (
  "math"
  "strconv"
  "testing"
)

func TestMinBitsZero(t *testing.T) {
  if min, _ := MinBits(0); min != 0 {
    t.Logf("Expected 0 got %b", min)
    t.FailNow()
  }

  if min, _ := MinBits(math.MaxInt32); min != math.MaxInt32 {
    t.Logf("Expected %b got %b", math.MaxInt32, min)
    t.FailNow()
  }

  v, _ := strconv.ParseInt("100000", 2, 64)
  if min, _ := MinBits(int(v)); min != 1 {
    t.Logf("Expected %b got %b", 1, min)
    t.FailNow()
  }

  v, _ = strconv.ParseInt("101000", 2, 64)
  if min, _ := MinBits(int(v)); min != 5 {
    t.Logf("Expected %b got %b", 5, min)
    t.FailNow()
  }

  v, _ = strconv.ParseInt("1101000", 2, 64)
  if min, _ := MinBits(int(v)); min != 11 {
    t.Logf("Expected %b got %b", 11, min)
    t.FailNow()
  }
}
