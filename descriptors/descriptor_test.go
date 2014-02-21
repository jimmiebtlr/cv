package descriptor

import (
  "image"
  "testing"
)

func TestDescriptor(t *testing.T) {
  img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{32, 32}})

  for i := 0; i <= 32; i++ {
    for j := 0; j <= 32; j++ {
      if j >= 16 && i >= 16 {
        img.Set(i, j, image.White)
      } else {
        img.Set(i, j, image.Black)
      }
    }
  }

  points := make(chan image.Point, 2)

  go func() {
    points <- image.Point{16, 16}
    close(points)
  }()

  desc := make(chan Descriptor)
  CalcDescriptors(points, img, desc, RGBPointValue)

  i := 0
  for d := range desc {
    if d.base != 15 {
      t.Logf("Base value of descriptor should be 15, instead it was %v", d.base)
      t.Fail()
    }
    i++
  }
  if i != 1 {
    t.Logf("CalcDescriptor should have returned exactly 1 descriptor, given 1 point")
    t.Fail()
  }
}
