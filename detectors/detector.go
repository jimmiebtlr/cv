// Detector package provides the very basic beginning of a image keypoint detector implementation.
package detector

import (
  "github.com/jimmiebtlr/mask"
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
  return int(r/3 + g/3 + b/3)
}

// Keypoints launches a go routine which fills the passed in channel of points.
func (d *Detector) Keypoints(ch chan image.Point) {
  go d.fillKeypointBuffer(ch)
}

// fillKeypointBuffer loops through all pixels in detector.img and adds to chan
// if it was a keypoint.
func (d *Detector) fillKeypointBuffer(ch chan image.Point) {
  for x := d.img.Bounds().Min.X; x < d.img.Bounds().Max.X; x++ {
    for y := d.img.Bounds().Min.Y; y < d.img.Bounds().Max.Y; y++ {
      if d.examineArea(x, y) {
        ch <- image.Point{x, y}
      }
    }
  }
  close(ch)
}

// examineArea looks in the area surrounding x,y.
// if n points are above or n points are below, return keypoint=true.
func (d *Detector) examineArea(x int, y int) (keypoint bool) {
  // if within 3 of one of the edges, don't do anything (yet)
  if d.img.Bounds().Max.X-r < x {
    return false
  }
  if d.img.Bounds().Min.Y+r > y {
    return false
  }
  if d.img.Bounds().Min.X+r > x {
    return false
  }
  if d.img.Bounds().Max.Y-r < y {
    return false
  }

  abv := uint(0)
  bel := uint(0)
  for i := 0; i < 16; i++ {
    dx, dy := mask.DiamondOffset(i, r)

    // check if above or below threshold
    if d.Value(x+dx, y+dy, d) > (d.Value(x, y, d) + d.threshold) {
      abv++
    } else if d.Value(x+dx, y+dy, d) < (d.Value(x, y, d) - d.threshold) {
      bel++
    }
  }

  if abv > n {
    return true
  } else if bel > n {
    return true
  }
  return false
}
