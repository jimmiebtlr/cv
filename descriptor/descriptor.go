package descriptor

import (
  "fmt"
  "github.com/jimmiebtlr/cv"
  "github.com/jimmiebtlr/cv/mask"
  "image"
)

const (
  r = 3
)

type Descriptor struct {
  base uint16
}

func (d *Descriptor) Value() (val uint16) {
  return d.base
}

type PxVal func(x, y int, img image.Image) (val int)

func RGBPointValue(x, y int, img image.Image) (val int) {
  r, g, b, _ := img.At(x, y).RGBA()
  return int(r/3.0 + g/3.0 + b/3.0)
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
  v := 0
  for i := 0; i < 16; i++ {
    fmt.Println(v)
    v += pointVal(i, point, img, f)
  }
  v2, _ := cv.MinBits(uint(v))
  desc = Descriptor{uint16(v2)}
  return desc
}

// pointVal returns 1 if the index in question is greater than point and
// 0 if it is less or equal to
func pointVal(i int, point image.Point, img image.Image, f PxVal) (val int) {
  dx, dy := mask.DiamondOffset(i, r)
  if f(point.X, point.Y, img) <= f(point.X+dx, point.Y+dy, img) {
    return 0
  } else {
    return 1 * pow(2, i)
  }
}

func pow(v, exp int) (ret int) {
  ret = 1
  for i := 0; i < exp; i++ {
    ret *= v
  }
  return ret
}
