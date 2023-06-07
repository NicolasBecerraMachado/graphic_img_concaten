package main

import(
	"fmt"
	"os"
	"log"
	"image"
	"bufio"
	//"io"
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
func GIRO(im *image.NRGBA){
	/*for i := 0; i < im.Height; i++{
		for j := 0; j < im.Width; j++{

		}
	}*/
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

func getConf(filePath string) image.Config{
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	
	reader := bufio.NewReader(imgFile)

	Rconfig,_,err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("Cannot get config:", err)
	}

	return Rconfig
}

func main(){
	//read images
	R := load("R.png")
	
	S:= load("S.png")

	Rconfig := getConf("R.png")

	Sconfig := getConf("S.png")

	fmt.Println("--------------------------")
	fmt.Println(R.Pix[0])
	fmt.Println(S.Pix[0], S.Pix[1], S.Pix[255*255])
	fmt.Println("--------------------------")
	fmt.Println(Rconfig.Height)
	fmt.Println(Sconfig.Width)
	//call functions to make transformations

	//print image

	//store image
	
}