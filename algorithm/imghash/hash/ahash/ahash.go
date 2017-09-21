// Implement of average hash algorithm
package ahash

import (
	"log"
	"github.com/nfnt/resize"
	_ "image/jpeg"

	"imghash/hash"
)

func GetHash(imgPath string) uint64 {
	img, err := hash.LoadImg(imgPath)
	if err != nil {
		log.Fatal(err)
	}

	thumbnail := resize.Resize(8, 8, img, resize.Bilinear)
	gray2dArray := hash.Rgb2Gray(thumbnail)
	gray1dArray, avg := hash.Arr2DTo1D(gray2dArray)

	var h uint64
	for i, gray := range gray1dArray {
		if gray > avg {
			h |= 1 << uint(i)
		}
	}
	return h
}