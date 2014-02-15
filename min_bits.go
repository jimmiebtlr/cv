package cv

import (
  "strconv"
  "fmt"
)

// MinBits returns the smallest possible value while not
// changing the order of the bits.  It does however allow
// reversal of the bits.
func MinBits(value uint) (min uint, reversed bool) {
  bits := fullLenUintStr( value )

  zerosStart, zerosEnd := maxZerosRun( bits )

  minStr := bits[zerosStart:] + bits[:zerosEnd]


  mintmp,err := strconv.ParseUint( minStr, 2, 64)

  min = uint( mintmp )
  if err != nil {
    // rabble
  }

  minRevStr := ""
  for _,v := range minStr {
    minRevStr = string(v) + minRevStr
  }
  fmt.Println( minRevStr )
  mintmpRev, err := strconv.ParseUint( minRevStr, 2, 64 )

  if mintmpRev < mintmp {
    min = uint(mintmpRev)
  }else{
    min = uint(mintmp)
  }

  // find longest string of zero's
  // determine if reverse is smaller than normal ordering
  return min, false
}

// fullLenUintStr returns the value as binary string 
// with precending 0's correctly filled in
func fullLenUintStr( v uint )( vstr string ){
  bits := strconv.FormatUint(uint64(v), 2)
  for i := len(bits); i < 32; i++ {
    bits = "0" + bits
  }
  return bits
}

// maxZerosRun returns the start and end index of the larges consecutive
// run of zeros in the string
func maxZerosRun( bits string )( start, end int ){
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
    fmt.Printf( "%v %v \n", curStart, curLen )
    fmt.Printf( "%v \n", bits[i%l] )
  }
  return maxStart, (maxStart+maxLen) % len( bits )
}
