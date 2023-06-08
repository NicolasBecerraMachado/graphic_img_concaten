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
	
	newIm := image.NewNRGBA(image.Rect(0,0,im.Rect.Max.X, im.Rect.Max.Y))

	for y := 0; y < im.Rect.Max.Y; y++{
		for x := 0; x < im.Rect.Max.X; x++{
			newIm.Set(x,y,im.At(y,(im.Rect.Max.X - 1 - x)))
		}
	}

	return newIm
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

	//call functions to make transformations
	/*tst := GIRO(S)
	
	tst2 := UneX(S,GIRO(R))
	fmt.Println(tst2.Rect)
	//store image
	save("tst.png", tst)
	tst3:= UneY(tst2,UneX(R,GIRO(S)))
	save("tst2.png", tst2)
	fmt.Println(tst3.Rect)
	save("tst3.png", tst3)
	tst4:=UneY(R,S)
	tst4 = UneX(tst4,UneY(S,R))
	fmt.Println(tst4.Rect)
	fmt.Println("Y done")
	save("tst4.png", tst4)*/

	row1 := UneX(UneX(S,GIRO(GIRO(R))),UneX(S,GIRO(GIRO(R))))
	
	row2 :=  UneX(UneX(R,GIRO(GIRO(S))),UneX(R,GIRO(GIRO(S))))

	row3 :=  UneX(UneX(S,GIRO(GIRO(R))),UneX(S,GIRO(GIRO(R))))

	row4 :=  UneX(UneX(R,GIRO(GIRO(S))),UneX(R,GIRO(GIRO(S))))

	image := UneY(UneY(row1,row2),UneY(row3,row4))
	save("image.png", image)
}