package main

import(
	"fmt"
	"os"
	"log"
	"image"
	//"bufio"
	//"crypto/rand"
	//"io"
	"image/png"
	"image/draw"
	//"image/jpeg"
	//"string"
)


//joins images the image in the second parameter to the bottom in the first
//images should have the sema bAse length
//returns joint image
func UneY(im1, im2 *image.NRGBA)(*image.NRGBA){
	newIm := image.NewNRGBA(image.Rect(0,0,im1.Rect.Max.X, im1.Rect.Max.Y + im2.Rect.Max.Y))
	draw.Draw(newIm, im1.Bounds(), im1, image.Pt(0,0),draw.Src)
	draw.Draw(newIm, image.Rect(0, im1.Rect.Max.Y,im1.Rect.Max.X, im1.Rect.Max.Y + im2.Rect.Max.Y), im2, image.Pt(0,0),draw.Src)
	fmt.Println(newIm.Bounds())
	return newIm
}

//joins images the image in the second parameter to the image in the first
//images should have the same height length
//returns joint image
func UneX(im1, im2 *image.NRGBA)(*image.NRGBA){
	newIm := image.NewNRGBA(image.Rect(0,0,im1.Rect.Max.X + im2.Rect.Max.X, im1.Rect.Max.Y))
	draw.Draw(newIm, im1.Bounds(), im1, image.Pt(0,0),draw.Src)
	draw.Draw(newIm, image.Rect(im1.Rect.Max.X, 0, im1.Rect.Max.X + im2.Rect.Max.X, im1.Rect.Max.Y), im2, image.Pt(0,0),draw.Src)
	fmt.Println(newIm.Bounds())
	return newIm
}

//rotates image 180°
//returns rotated image
func GIRO(im *image.NRGBA) (*image.NRGBA){
	//fmt.Println("size ni = ",conf.Height, conf.Height)
	
	pix := make([]uint8, im.Rect.Max.X*im.Rect.Max.Y*4)

	for y := 0; y < im.Rect.Max.Y; y++{
		for x := 0; x < im.Rect.Max.X; x++{
			for k:= 0; k < 4; k++{
				pix[y*im.Rect.Max.X*4 + x*4 + k] = im.Pix[(im.Rect.Max.Y-1-y)*im.Rect.Max.X*4 + (im.Rect.Max.X-1-x)*4 + k]
			}
		}
	}
	created := &image.NRGBA{
		Pix:    pix,
		Stride: im.Rect.Max.X*2 + im.Rect.Max.Y*2,
		Rect:   image.Rect(0,0,im.Rect.Max.X, im.Rect.Max.Y),
	}
	return created
}

//load image
func load(filePath string) (*image.NRGBA) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}

	return img.(*image.NRGBA)
}

func save(filePath string, img *image.NRGBA) {
	imgFile, err := os.Create(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	fmt.Println("set aux")
	//aux := img.SubImage(img.Rect)
	fmt.Println("Problem in Encode")
	//png.Encode(imgFile, img.SubImage(image.Rect{0,0,256,256}))

	encoder := png.Encoder{CompressionLevel: png.BestCompression}
	fmt.Println(img.Rect)
	err = encoder.Encode(imgFile, img.SubImage(img.Rect))
	if err != nil {
		log.Fatal(err)
	}


}

func main(){
	//read images
	R := load("R.png")
	
	S:= load("S.png")

	fmt.Println("--------------------------")
	fmt.Println(R.Pix[0])
	fmt.Println(S.Pix[0], S.Pix[1], S.Pix[255*255])
	fmt.Println("--------------------------")
	//call functions to make transformations
	tst := GIRO(S)
	
	tst2 := UneX(S,GIRO(R))
	fmt.Println(tst2.Rect)
	//store image
	save("tst.png", tst)
	tst3:= UneY(tst2,UneX(R,GIRO(S)))
	save("tst2.png", tst2)
	fmt.Println(tst3.Rect)
	save("tst3.png", tst3)
	/*tst4:=UneY(R,S)
	tst4 = UneX(tst4,UneY(S,R))
	fmt.Println(tst4.Rect)
	fmt.Println("Y done")
	save("tst4.png", tst4)*/
}