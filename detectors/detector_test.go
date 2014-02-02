package detector

import (
  "image"
  "strconv"
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
  keypoint := d.examineArea(16, 16)
  if !keypoint {
    t.Fatal("Area should return as keypoint")
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
  keypoint := d.examineArea(16, 16)
  if !keypoint {
    t.Fatal("Area should return as keypoint")
  }
}

func TestKeypoints(t *testing.T) {
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

  ch := make(chan image.Point)
  d.Keypoints(ch)
  i := 0
  for _ = range ch {
    i++
  }
  if i != 1 {
    t.Fatal("Expected one keypoint from image with single corner. Instead " + strconv.Itoa(i) + " were found. ")
  }
}
