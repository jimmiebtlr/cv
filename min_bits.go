package cv

import (
  "strconv"
)

// MinBits returns the smallest possible value while not
// changing the order of the bits.  It does however allow
// reversal of the bits.
func MinBits(value int) (min int, reversed bool) {
  bits := strconv.FormatInt(int64(value), 2)

  maxStart := 0
  maxLen := 0

  curStart := 0
  curLen := 0

  for i := 0; i < len(bits)*2; i++ {
    if bits[i] == '0' {
      if curLen == 0 {
        curStart = i
      }
      curLen++
    } else {
      if maxLen < curLen {
        maxLen = curLen
        maxStart = curStart
        curLen = 0
        curStart = 0
      }
    }
  }

  // find longest string of zero's
  // determine if reverse is smaller than normal ordering
  return 0, false
}
