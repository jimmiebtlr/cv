package briskDetector

import (
  "image"
  "testing"
)

func TestExamineArea(t *testing.T) {
  img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{32, 32}})

  for i := 0; i <= 32; i++ {
    for j := 0; j <= 32; j++ {
      
    }
  }

  d := Detector{}
  d.threshold = 40
  d.img = img
  d.Value = RGBPointValue
}
