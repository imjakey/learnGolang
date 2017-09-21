// Implement of perceptual hash algorithm
package phash

import (
	"math"
	"log"
	"github.com/nfnt/resize"

	"imghash/hash"
)

func GetHash(imgPath string) uint64 {
	img, err := hash.LoadImg(imgPath)
	if err != nil {
		log.Fatal(err)
	}

	thumbnail := resize.Resize(64, 64, img, resize.Bilinear)
	gray2DArr := hash.Rgb2Gray(thumbnail)

	dct2DArr := DCT2D(gray2DArr, 64, 64)

	// 获取dct变换后的低频部分，即左上角部分
	arr8x8 := hash.Cut2DArr(dct2DArr, 8, 8)
	dct1DArr, avg := hash.Arr2DTo1D(arr8x8)

	var h uint64
	for i, gray := range dct1DArr {
		if gray > avg {
			h |= 1 << uint(i)
		}
	}
	return h
}

// 1维dct变换
func DCT1D (in []float64) []float64 {
	l := len(in)
	out := make([]float64, l)
	for i := 0; i < l; i ++ {
		sum := 0.0
		for n := 0; n < l; n ++ {
			sum += in[n] * math.Cos(math.Pi * float64(i) * (float64(n) + 0.5) / float64(l))
		}

		sum *= math.Sqrt(2.0 / float64(l))
		if i == 0 {
			sum *= 1.0 / math.Sqrt(2.0)
		}
		out[i] = sum
	}
	return out
}

// 2维dct变换，以1维变换为基础，分别对行和列做变换
func DCT2D(in [][]float64, w int, h int) [][]float64 {
	out := make([][]float64, h)
	for i := 0; i < h; i++ {
		out[i] = DCT1D(in[i])
	}

	for j := 0; j < w; j++ {
		column := make([]float64, h)
		for k := 0; k < h; k ++ {
			column[k] = out[k][j]
		}
		column = DCT1D(column)

		for k := 0; k < h; k ++ {
			out[k][j] = column[k]
		}
	}
	return out
}