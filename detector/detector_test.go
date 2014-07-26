package detector

import (
  "image"
  "image/png"
  "os"
  "strconv"
  "testing"
)

func TestExamineAreaBelow(t *testing.T) {
  square, err := os.Open("square.png")

  if err != nil {
    t.Fatal(err)
  }

  img, err := png.Decode(square)

  if err != nil {
    t.Fatal(err)
  }

  d := Detector{}
  d.threshold = 40
  d.img = img
  d.Value = RGBPointValue
  keypoint := d.isKeypoint(16, 16)
  if !keypoint {
    t.Fatal("Area should return as keypoint")
  }
}

func TestExamineAreaAbove(t *testing.T) {
  square, err := os.Open("square.png")

  if err != nil {
    t.Fatal(err)
  }

  img, err := png.Decode(square)

  d := Detector{}
  d.threshold = 40
  d.img = img
  d.Value = RGBPointValue
  keypoint := d.isKeypoint(16, 16)
  if !keypoint {
    t.Fatal("Area should return as keypoint")
  }
}

func TestKeypoints(t *testing.T) {
  square, err := os.Open("square.png")

  if err != nil {
    t.Fatal(err)
  }

  img, err := png.Decode(square)

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
  if i != 4 {
    // this is correct for the moment, since we don't do any suppression of points near each other
    // a corner will have four points that meet the criteria of a keypoint
    t.Fatal("Expected four keypoints from image with single corner. Instead " + strconv.Itoa(i) + " were found. ")
  }
}
