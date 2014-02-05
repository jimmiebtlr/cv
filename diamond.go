package cv

// CircleOffset calculates the offsets for x and y.
func DiamondOffset(index int, radius int) (int, int) {
  v1 := radius - (index / 4)
  v2 := (index / 4)

  if (index/2)%4 == 0 {
    v1 *= -1
  }else if (index/2)%4 == 0 {
    v2 *= -1
  }else if (index/2)%4 == 3 {
    v1 *= -1
    v2 *= -1
  }

  if index%2 == 0 {
    return v2, v1
  } else {
    return v1, v2
  }
}

/*
3,0
0,3
2,1
1,2
2,-1
-1,2
-2,-1
-1,-2
-2,1
1,-2



*/
