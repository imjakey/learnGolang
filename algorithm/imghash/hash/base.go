package hash

import (
	"image"
	"os"
	"fmt"
)

// 使用位运算计算汉明距离
func HammingDistance(h1, h2 uint64) uint8 {
	hamming := h1 ^ h2
	var dis uint8 = 0
	for hamming != 0 {
		dis += uint8(hamming & 1)
		hamming >>= 1
	}
	return dis
}

// 加载并解析图片
func LoadImg(imgPath string) (img image.Image, err error) {
	file, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err, 111)
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

// Rgb转换成灰度图
func Rgb2Gray(colorImg image.Image) [][]float64 {
	bounds := colorImg.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	pixels := make([][]float64, h)

	for i := range pixels {
		pixels[i] = make([]float64, w)
		for j := range pixels[i] {
			color := colorImg.At(j, i)
			r, g, b, _ := color.RGBA()
			lum := 0.299*float64(r/257) + 0.587*float64(g/257) + 0.114*float64(b/256)
			pixels[i][j] = lum
		}
	}
	return pixels
}

// 切割2维数组
func Cut2DArr(in [][]float64, w int, h int) [][]float64 {
	out := make([][]float64, h)
	for i := 0; i < w; i ++ {
		out[i] = make([]float64, w)
		for j := 0; j < h; j++ {
			out[i][j] = in[i][j]
		}
	}
	return out
}

// 2维数组降为1维并返回平均值
func Arr2DTo1D(gray2dArray [][]float64) ([]float64, float64) {
	var LinearArray []float64
	var sum float64
	var counts float64
	for _, row := range gray2dArray {
		for _, gray := range row {
			LinearArray = append(LinearArray, gray)
			sum += gray
			counts ++
		}
	}
	return LinearArray, sum / counts
}