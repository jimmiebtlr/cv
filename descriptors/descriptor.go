package briskDescriptor

import (
  "github.com/jimmiebtlr/cv"
  "image"
)

type Descriptor struct {
  uint16
}

func CalcDescriptors(points chan image.Point, descr chan Descriptor) {
  for p := range points {
    calcDescriptor(p)
  }

  // Calculate multiple rotations
  rawDesc := uint(0)
  for i := 0; i < 16; i++ {

  }

  // compact and order them
}

func calcDescriptor(points image.Point) (desc Descriptor) {
  return desc
}
