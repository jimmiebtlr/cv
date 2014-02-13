package briskDescriptor

import(
  "image"
)

type Descriptor struct{
  int8
}

func CalcDescriptors( points chan image.Point, descr chan Descriptor ){
  // Calculate multiple rotations
  // compact and order them
}
