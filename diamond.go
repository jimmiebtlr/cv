package cv

// CircleOffset calculates the offsets for x and y.
func DiamondOffset(index int, radius int) (int, int) {
  v1 := radius - ((index + 4) / 8)
  v2 := (index + 4) / 8

  if (index/2)%2 == 1 {
    v1 *= -1
  }
  if (index/4)%2 == 1 {
    v2 *= -1
  }

  if index%2 == 1 {
    return v2, v1
  } else {
    return v1, v2
  }
}

// every 2 flip v1 sign
// every 4 flip v2 sign
// every 8 dec v2 and inc v1

/*
3,0
0,3
-3,0
0,-3

2,-1
-1,2
-2,-1
-1,-2

2,1
1,2
-2,1
1,-2



*/
