package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GifJsPolyglot(gif GIF, js []byte) {
	// GIFa8=/* GIF meta data + image block */=1;Javascript
	gif.Width = []byte("/*")
	gifbytes := Concat(gif)

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

func JpegHTMLPolyglot(jpeg JPEG, html []byte) {
	// ....FF D8 FF E0
	htmlLength := len(html)
	log.Println(htmlLength)

	//Extend APP0 length to inject html code in the header + <!-- comment
	jpeg.Length[1] = byte(htmlLength) + jpeg.Length[1] + 4

	log.Println(jpeg.Length[1])
	html = append(html, []byte("<!--")...)
	jpeg.Image = append(html, jpeg.Image...)
	jpegbyte := Concat(jpeg)
	err := ioutil.WriteFile("jpeg.html.html", jpegbyte, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	imagefile := os.Args[1]
	jsfile := os.Args[2]

	//Currently checks the file type by parsing the name
	//Make better later
	imageSplit := strings.Split(imagefile, ".")
	imageFileType := imageSplit[len(imageSplit)-1]

	log.Println("Opening Image File: " + imagefile)
	bytearray, err := ioutil.ReadFile(imagefile)
	log.Println("Image File Type is: " + imageFileType)

	//fail to load image file
	if err != nil {
		log.Fatal("Could not load Image File")
	}

	log.Println("Opening JS File: " + jsfile)
	jsarray, err := ioutil.ReadFile(jsfile)

	//fail to load image file
	if err != nil {
		log.Fatal("Could not load JS File")
	}

	//Invote different polyglot function for different filetype
	switch imageFileType {
	case "gif":
		gif := GIFParse(bytearray)
		GifJsPolyglot(gif, jsarray)
	case "jpg", "jpeg":
		jpeg := JPEGParse(bytearray)
		JpegHTMLPolyglot(jpeg, jsarray)
	}

}
