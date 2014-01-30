// briskDetector package provides the beginning of a brisk detector implementation.
package briskDetector

import (
  "image"
)

const (
  r = 4
  n = 9
)

type Detector struct {
  threshold int
  img       image.Image
  Value     PxVal
}

type PxVal func(x, y int, d *Detector) (val int)

func RGBPointValue(x, y int, d *Detector) (val int) {
  r, g, b, _ := d.img.At(x, y).RGBA()
  return int(r/3 + g/3 + b/3)
}

// examineArea 
func (d *Detector) examineArea(x int, y int) (above, below bool) {
  // if within 3 of one of the edges, don't do anything (yet)
  if d.img.Bounds().Max.X+r < x {
    return false, false
  }
  if d.img.Bounds().Min.Y+r > y {
    return false, false
  }
  if d.img.Bounds().Min.X-r > x {
    return false, false
  }
  if d.img.Bounds().Max.Y-r < y {
    return false, false
  }

  abv := uint(0)
  bel := uint(0)
  // TODO
  // 23 is 4! - 1 
  for i := 0; i < 23; i++ {
    dx := offsets[i][0]
    dy := offsets[i][1]

    // check if above or below threshold
    if d.Value(x+dx, y+dy, d) > (d.Value(x, y, d) + d.threshold) {
      abv++
    } else if d.Value(x+dx, y+dy, d) < (d.Value(x, y, d) - d.threshold) {
      bel++
    }
  }
  if abv > n {
    return true, false
  } else if bel > n {
    return false, true
  }
  return false, false
}
