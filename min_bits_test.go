package cv

import (
  //"math"
  //"strconv"
  "testing"
)

func TestMaxZerosRun(t *testing.T) {
  if start, l := maxZerosRun("0"); start != 0 || l != 1 {
    t.Logf("Expected start = 0 and length = 1 actual start = %v and length = %v", start, l)
    t.Fail()
  }
  if start, l := maxZerosRun("010"); start != 2 || l != 2 {
    t.Logf("Expected start = 2 and length = 1 actual start = %v and length = %v", start, l)
    t.Fail()
  }
  if start, end := maxZerosRun("1000101"); start != 1 || end != 3 {
    t.Logf("Expected start = 1 and end = 3 actual start = %v and end = %v", start, end)
    t.Fail()
  }
}

/*
func TestMinBitsZero(t *testing.T) {
  if min, _ := MinBits(uint(0)); min != 0 {
    t.Logf("Expected 0 got %b", min)
    t.FailNow()
  }

  if min, _ := MinBits(math.MaxInt32); min != math.MaxInt32 {
    t.Logf("Expected %b got %b", math.MaxInt32, min)
    t.FailNow()
  }

  v, _ := strconv.ParseInt("100000", 2, 64)
  if min, _ := MinBits(uint(v)); min != 1 {
    t.Logf("Expected %b got %b", 1, min)
    t.FailNow()
  }

  v, _ = strconv.ParseInt("101000", 2, 64)
  if min, _ := MinBits(uint(v)); min != 5 {
    t.Logf("Expected %b got %b", 5, min)
    t.FailNow()
  }

  v, _ = strconv.ParseInt("1101000", 2, 64)
  if min, _ := MinBits(uint(v)); min != 11 {
    t.Logf("Expected %b got %b", 11, min)
    t.Fail()
  }
}*/
