package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type GIF struct {
	Signature       []byte
	Version         []byte
	Width           []byte
	Height          []byte
	Packed          byte
	BackgroundColor byte
	AspectRatio     byte
	Image           []byte
}

func GIFParse(gifarray []byte) GIF {
	gif := GIF{}

	//Invalid header size
	if len(gifarray) < 13 {
		log.Fatal("Invalid GIF Format: Invalid header length")
	}
	gif = GIF{
		Signature:       gifarray[0:3],
		Version:         gifarray[3:6],
		Width:           gifarray[6:8],
		Height:          gifarray[8:10],
		Packed:          gifarray[10],
		BackgroundColor: gifarray[11],
		AspectRatio:     gifarray[12],
		Image:           gifarray[13:],
	}

	//Check signature
	if string(gif.Signature) != "GIF" {
		log.Fatal("Invalid GIF Format: Invalid Signature")
	}

	return gif
}

func GifToByteArray(gif GIF) []byte {
	Concat := []byte{}
	Concat = append(Concat, gif.Signature...)
	Concat = append(Concat, gif.Version...)
	Concat = append(Concat, gif.Width...)
	Concat = append(Concat, gif.Height...)
	Concat = append(Concat, gif.Packed)
	Concat = append(Concat, gif.BackgroundColor)
	Concat = append(Concat, gif.AspectRatio)
	Concat = append(Concat, gif.Image...)
	return Concat
}

func GifJsPolyglot(gif GIF, js []byte) {
	// GIFa8=/* GIF meta data + image block */=1;Javascript
	gif.Width = []byte("/*")
	gifbytes := GifToByteArray(gif)

	//Remove '*/' to ensure that the image block gets commented out [42 47] (2a2f)
	for i := 0; i < len(gifbytes)-2; i++ {
		if gifbytes[i] == byte('*') && gifbytes[i+1] == byte('/') {
			gifbytes = append(gifbytes[:i], 42)
			gifbytes = append(gifbytes[:i], gifbytes[i+1:]...)
		}
	}

	gifbytes = append(gifbytes, []byte("*/")...)
	gifbytes = append(gifbytes, []byte("=1;")...)
	gifbytes = append(gifbytes, js...)

	err := ioutil.WriteFile("giphy.js.gif", gifbytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func JpegJsPolyglot(jpeg []byte, js []byte) {

}

func main() {
	imagefile := os.Args[1]
	jsfile := os.Args[2]

	log.Println("Opening Image File: " + imagefile)
	bytearray, err := ioutil.ReadFile(imagefile)

	log.Println("Opening JS File: " + jsfile)
	jsarray, err := ioutil.ReadFile(jsfile)
	gif := GIFParse(bytearray)
	if err != nil {
		fmt.Println("Could not load")
	}
	gg := []byte("*/")

	fmt.Println(gg)
	GifJsPolyglot(gif, jsarray)
}
