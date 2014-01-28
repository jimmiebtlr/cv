package briskDetector

import (
  "image"
)

const (
  r = 3
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
  return int(r + g + b)
}

func (d *Detector) examineArea(x int, y int) (above, below bool) {
  // if within 3 of one of the edges, don't do anything (yet)
  if d.img.Bounds().Max.X+r > x {
    return false, false
  }
  if d.img.Bounds().Min.Y+r > y {
    return false, false
  }
  if d.img.Bounds().Max.X-r < x {
    return false, false
  }
  if d.img.Bounds().Max.Y-r < y {
    return false, false
  }

  abv := uint(0)
  bel := uint(0)
  for i := 0; i < r*r; i++ {
    // check if above or below threshold
    if d.Value(x, y, d) > (d.Value(x, y, d) - d.threshold) {
      abv++
    } else if d.Value(x, y, d) < (d.Value(x, y, d) - d.threshold) {
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
