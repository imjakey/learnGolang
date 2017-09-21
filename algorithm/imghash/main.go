package main

import (
	"fmt"
	"strconv"

	"imghash/hash"
	"imghash/hash/ahash"
	"imghash/hash/phash"
)

const IMGPATH = "D:\\workspace\\golang\\src\\imghash\\resource\\img\\similarimg\\"

func main() {
	img1 := IMGPATH + "girl1.jpg"
	h1 := ahash.GetHash(img1)

	fmt.Println("img1\timg2\tdistance\ttype")
	for n := 1; n <= 20; n++ {
		strN := strconv.Itoa(n)
		imgN := IMGPATH + "girl" + strN + ".jpg"
		hN := ahash.GetHash(imgN)

		fmt.Println("girl1\tgirl" + strN + "\t" + strconv.Itoa(int(hash.HammingDistance(h1, hN))) + "\tahash")
	}

	fmt.Println("---------------- im a dividing line --------------")

	fmt.Println("img1\timg2\tdistance\ttype")
	ph1 := phash.GetHash(img1)
	for n := 1; n <= 20; n++ {
		strN := strconv.Itoa(n)
		imgN := IMGPATH + "girl" + strN + ".jpg"
		phN := phash.GetHash(imgN)

		fmt.Println("girl1\tgirl" + strN + "\t" + strconv.Itoa(int(hash.HammingDistance(ph1, phN))) + "\tphash")
	}
}
