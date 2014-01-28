package brisk_detector

import  (
  "image"
  "github.com/nfnt/resize"
)

const(
  r = 3
  n = 9
)

type Detector struct{
  octaves []image.Image
  pointBuff chan<-image.Point
  threshold uint
}

func (d *Detector)GetPoints( )( locations []image.Point, err error ){
  if( len(d.octaves) == 0 ){ return nil,nil }

  width := float64(d.octaves[0].Bounds().Dx())/1.5
  downsampled := resize.Resize( uint(width), 0, d.octaves[0], resize.Bilinear)
  d.octaves = append( d.octaves, downsampled )

  d.pointBuff = make( chan image.Point )

  numPixels := int(d.octaves[0].Bounds().Dx()*d.octaves[0].Bounds().Dy())
  for i := int(0); i<numPixels; i++ {
    go d.examineArea( i/d.octaves[0].Bounds().Dx(), i%d.octaves[0].Bounds().Dy(), 0)
  }

  // get number of octaves further halfsampled
  // Make some halfsampled octaves

  // Start the detector

  // Read back from buffer into locations array

  return locations,nil
}

func (d *Detector)examineArea( x int, y int, layer uint){
  // if within 3 of one of the edges, don't do anything (yet)
  if d.octaves[layer].Bounds().Max.X > (x-d.octaves[layer].Bounds().Min.X) || d.octaves[layer].Bounds().Max.X > (x-d.octaves[layer].Bounds().Min.X) { return }

  above := uint(0)
  below := uint(0)
  img := d.octaves[layer]
  for i := 0; i<r*r; i++{
    // check if above or below threshold
    if img.At(x,y) > (img.At(x,y)-d.threshold) {
      above++
    }else if img.At(x,y) < (img.At(x,y) - d.threshold) {
      below++
    }
    if above > n{
      d.pointBuff <- image.Point{ x,y }
    }
  }
}
