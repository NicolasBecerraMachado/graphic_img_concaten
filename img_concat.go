package main

import(
	"fmt"
	"os"
	"log"
	"image"
	"bufio"
	//"crypto/rand"
	//"io"
	"image/png"
	//"string"
)


//joins images the image in the second parameter to the bottom in the first
//images should have the sema bAse length
//returns joint image
func UneY(){

}

//joins images the image in the second parameter to the image in the first
//images should have the same height length
//returns joint image
func UneX(im1, im2 *image.NRGBA){

}

//rotates image 180°
//returns rotated image
func GIRO(im *image.NRGBA) (*image.NRGBA){
	//fmt.Println("size ni = ",conf.Height, conf.Height)
	
	pix := make([]uint8, im.Rect.Max.X*im.Rect.Max.Y*4)

	for i := 0; i < im.Rect.Max.Y; i++{
		for j := 0; j < im.Rect.Max.Y; j++{
			for k:= 0; k < 4; k++{
				pix[j*im.Rect.Max.X*4 + i*4 + k] = im.Pix[i*im.Rect.Max.X*4 + j*4 + k]
			}
		}
	}
	//rand.Read(pix)
	created := &image.NRGBA{
		Pix:    pix,
		Stride: im.Rect.Max.X*2 + im.Rect.Max.Y*2,
		Rect:   image.Rect(0,0,im.Rect.Max.Y, im.Rect.Max.X),
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
	fmt.Println("weird line")
	png.Encode(imgFile, img.SubImage(img.Rect))
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
	
	fmt.Println(len(S.Pix))

	/*f, err := os.Create("data.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

	_, err2 := f.WriteString(string(tst.Pix))

    if err2 != nil {
        log.Fatal(err2)
    }*/

	save("tst.png", tst)
	//print image

	//store image
	
}