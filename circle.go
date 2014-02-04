package cv

// CircleOffset calculates the offsets for x and y.
func CircleOffset(index int, radius int) (int, int) {
  // normal
  // swap x and y
  // remove 1

  v1 := radius - (index / 4)
  v2 := (index / 4)
  if index %4 == 1 {
    v1 *= -1
  }
  if index %4 == 2 {
    v2 *= -1
  }
  if index%4 == 3 {
    v1 *= -1
    v2 *= -1
  }

  //if index%2 == 0 {
    return v2, v1
  //} else {
  //  return v1, v2
  //}
}
