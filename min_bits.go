package cv

import (
  "log"
  "strconv"
)

// MinBits returns the smallest possible value while not
// changing the order of the bits.  It does however allow
// reversal of the bits.
func MinBits(value uint) (min uint, reversed bool) {
  bits := strconv.FormatUint(uint64(value), 2)

  zerosStart, zerosLen := maxZerosRun(bits)

  minStr := ""
  if zerosStart+zerosLen > len(bits) {
    minStr = bits[:zerosStart+zerosLen] + bits[zerosStart:]
  } else {
    minStr = bits[zerosStart+zerosLen:] + bits[:zerosStart]
  }

  // unreversed case
  mintmp := uint64(0)
  var err error
  if minStr != "" {
    mintmp, err = strconv.ParseUint(minStr, 2, 64)
  }
  if err != nil {
    log.Printf("Error: %v", err)
    log.Fatal("MinBits parse uint failed")
  }

  // Reversed case
  mintmpRev := uint64(0)
  if minStr != "" {
    mintmpRev, err = strconv.ParseUint(revStr(minStr), 2, 64)
  }
  if err != nil {
    log.Printf("Error: %v", err)
    log.Fatal("MinBits parse uint failed")
  }

  // if the reversed string is smaller use it
  if mintmpRev < mintmp {
    min = uint(mintmpRev)
    reversed = true
  } else {
    min = uint(mintmp)
  }

  return min, reversed
}

// revStr finds the exact reverse of the input string
func revStr(str string) (rev string) {
  for _, v := range str {
    rev = string(v) + rev
  }
  return rev
}

// maxZerosRun returns the start and end index of the larges consecutive
// run of zeros in the string
func maxZerosRun(bits string) (start, length int) {
  maxStart := 0
  maxLen := 0

  curStart := 0
  curLen := 0

  l := len(bits)

  for i := 0; i < l*2; i++ {
    if bits[i%l] == '0' {
      if curLen == 0 {
        curStart = i
      }
      curLen++
    } else {
      if maxLen < curLen {
        maxLen = curLen
        maxStart = curStart
      }
      curLen = 0
      curStart = 0
    }
  }

  if maxLen < curLen {
    maxLen = curLen
    maxStart = curStart
  }

  if maxLen > l {
    return maxStart, l
  } else {
    return maxStart, maxLen
  }
}
