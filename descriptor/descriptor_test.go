package descriptor

import (
  "image"
  "image/png"
  "os"
  "testing"
)

func TestDescriptor(t *testing.T) {
  square, err := os.Open("../test/images/square.png")
  if err != nil {
    t.Fatal(err)
  }
  img, err := png.Decode(square)

  points := make(chan image.Point, 2)

  go func() {
    points <- image.Point{16, 16}
    close(points)
  }()

  desc := make(chan Descriptor)
  CalcDescriptors(points, img, desc, RGBPointValue)

  i := 0
  for d := range desc {
    if d.base != 16191 {
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

func TestDescriptorZero(t *testing.T) {
  square, err := os.Open("../test/images/square.png")
  if err != nil {
    t.Fatal(err)
  }
  img, err := png.Decode(square)

  points := make(chan image.Point, 2)

  go func() {
    points <- image.Point{4, 4}
    close(points)
  }()

  desc := make(chan Descriptor)
  CalcDescriptors(points, img, desc, RGBPointValue)

  i := 0
  for d := range desc {
    if d.base != 0 {
      t.Logf("Base value of descriptor should be 0, instead it was %v", d.base)
      t.Fail()
    }
    i++
  }
  if i != 1 {
    t.Logf("CalcDescriptor should have returned exactly 1 descriptor, given 1 point")
    t.Fail()
  }
}
