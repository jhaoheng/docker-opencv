package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"

	"gocv.io/x/gocv"
)

type PIC struct {
	PicPath string
	Image   image.Image
	Byte    []byte
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var PICs = []string{
		"imgs/b.jpg",
		"imgs/g.jpg",
		"imgs/r.jpg",
	}

	for _, jpgPath := range PICs {
		var pic = PIC{}
		pic = ProcessPic(jpgPath)
		ShowSummary(pic)
	}
}

func ProcessPic(jpgPath string) PIC {
	b := ReadFile(jpgPath)
	image := GetImage(b)

	var pic = PIC{}
	pic.PicPath = jpgPath
	pic.Byte = b
	pic.Image = image
	return pic
}

func ShowSummary(pic PIC) {
	Matrix, err := gocv.ImageToMatRGBA(pic.Image)
	defer Matrix.Close()
	if err != nil {
		panic(err)
	}

	B := Matrix.Mean().Val1
	G := Matrix.Mean().Val2
	R := Matrix.Mean().Val3
	A := Matrix.Mean().Val4
	fmt.Printf("The Picture : %v\n", pic.PicPath)
	fmt.Printf("B: %v, G: %v, R: %v, A: %v\n\n", B, G, R, A)
}

func ReadFile(jpgPath string) (b []byte) {
	b, err := ioutil.ReadFile(jpgPath)
	if err != nil {
		panic(err)
	}
	return
}

func GetImage(b []byte) (img image.Image) {
	img, extension, err := image.Decode(bytes.NewReader(b))
	if extension == "jpg" || extension == "jpeg" {
	} else if err != nil {
		panic(err)
	} else {
		panic("Pic Extension should be `jpg` or `jpeg`")
	}
	return
}

/*
output base64 string and could use with below url
https://codebeautify.org/base64-to-image-converter
*/
func MatToB64(Mat gocv.Mat) (B64 string) {
	toImg, _ := Mat.ToImage()
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, toImg, nil)
	b := buf.Bytes()

	B64 = base64.StdEncoding.EncodeToString(b)
	return
}
