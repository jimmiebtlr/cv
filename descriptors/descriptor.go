package descriptor

import (
  "github.com/jimmiebtlr/cv"
  "image"
  "math"
)

const (
  r = 3
)

type Descriptor struct {
  base uint16
}

type PxVal func(x, y int, img image.Image) (val int)

func RGBPointValue(x, y int, img image.Image) (val int) {
  r, g, b, _ := img.At(x, y).RGBA()
  return int(r/3 + g/3 + b/3)
}

func CalcDescriptors(points chan image.Point, img image.Image, descr chan Descriptor, f PxVal) {
  go fillDescrChan(points, img, descr, f)
}

func fillDescrChan(points chan image.Point, img image.Image, descr chan Descriptor, f PxVal) {
  for p := range points {
    descr <- calcDescriptor(p, img, f)
  }
  close(descr)
}

func calcDescriptor(point image.Point, img image.Image, f PxVal) (desc Descriptor) {
  v := uint16(0)
  for i := 0; i < 16; i++ {
    v += uint16(math.Pow(float64(pointVal(i, point, img, f)), float64(i)))
  }
  v2, _ := cv.MinBits(uint(v))
  desc = Descriptor{uint16(v2)}
  return desc
}

// pointVal returns 1 if the index in question is greater than point and
// 0 if it is less
func pointVal(i int, point image.Point, img image.Image, f PxVal) (val int) {
  dx, dy := cv.DiamondOffset(i, r)
  if f(point.X, point.Y, img) < f(point.X+dx, point.Y+dy, img) {
    return 0
  } else {
    return 1
  }
}
