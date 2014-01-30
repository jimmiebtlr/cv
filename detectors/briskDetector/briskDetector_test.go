package briskDetector

import (
  "image"
  "testing"
)

func TestExamineAreaBelow(t *testing.T) {
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

  d := Detector{}
  d.threshold = 40
  d.img = img
  d.Value = RGBPointValue
  abv, bel := d.examineArea(16, 16)
  if abv {
    t.Error("Above should have returned false")
  }
  if !bel {
    t.Error("Below should have returned true")
  }
}

func TestExamineAreaAbove(t *testing.T) {
  img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{32, 32}})

  for i := 0; i <= 32; i++ {
    for j := 0; j <= 32; j++ {
      if j >= 16 && i >= 16 {
        img.Set(i, j, image.Black)
      } else {
        img.Set(i, j, image.White)
      }
    }
  }

  d := Detector{}
  d.threshold = 40
  d.img = img
  d.Value = RGBPointValue
  abv, bel := d.examineArea(16, 16)
  if !abv {
    t.Error("Above should have returned true")
  }
  if bel {
    t.Error("Below should have returned false")
  }
}
