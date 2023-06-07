package main

import(
	"fmt"
	"os"
	"log"
	"image"

	_ "image/png"
)


//joins images the image in the second parameter to the bottom in the first
//images should have the sema bAse length
//returns joint image
func UneY(){

}

//joins images the image in the second parameter to the image in the first
//images should have the same height length
//returns joint image
func UneX(){

}

//rotates image 90°
//returns rotated image
func GIRO(){

}

//load image
func load(filePath string) *image.Gray {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		log.Println("Cannot read file:", err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
	}
	return img.(*image.Gray)
}

func main(){
	//read images
	R := load("R.png")
	S := load("S.png")
	fmt.Println(R.Pix[0])
	fmt.Println(S.Pix[0])
	//call functions to make transformations

	//print image

	//store image
	
}