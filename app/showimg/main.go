// What it does:
//
// This example uses the Window class to open an image file, and then display
// the image in a Window class.
//
// How to run:
//
// 		go run ./cmd/showimage/main.go /home/ron/Pictures/mcp23017.jpg
//
// +build example

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"gocv.io/x/gocv"
	"image/jpeg"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to run:\n\tshowimage [imgfile]")
		return
	}

	filename := os.Args[1]
	window := gocv.NewWindow("Hello")
	img := gocv.IMRead(filename, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Error reading image from: %v\n", filename)
		return
	}
	for {
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

}

/*
output base64 string and could use with below url
https://codebeautify.org/base64-to-image-converter
*/
func MatToB64(Mat gocv.Mat) {
	ti, _ := Mat.ToImage()
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, ti, nil)
	send_s3 := buf.Bytes()

	pic := base64.StdEncoding.EncodeToString(send_s3)
	fmt.Println(pic)
}
