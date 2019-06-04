package main

import (
	"log"
	"reflect"
)

//Ref GIFParse function for valid GIF byte format
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

type JPEG struct {
	SOI        []byte //Start of Image
	APP0       []byte //Application use marker
	Length     []byte //Length of APP0 Field
	Identifier []byte //JFIF + zerp terminated
	Version    []byte //JFIF Format Version
	Units      byte   //Units used for Resolution
	Xres       []byte //Horizontal Resolution
	Yres       []byte //Vertical Resolution
	XThumbnail byte   //Horizontal pixel count
	YThumbnail byte   //Vertical pixel count
	Image      []byte
}

//Convert GIF Byte array to GIF object
func GIFParse(gifarray []byte) GIF {

	//Invalid header size
	if len(gifarray) < 13 {
		log.Fatal("Invalid GIF Format: Invalid header length")
	}
	gif := GIF{
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

func JPEGParse(jpegarray []byte) JPEG {

	//Invalid header size
	if len(jpegarray) < 20 {
		log.Fatal("Invalid GIF Format: Invalid header length")
	}

	jpeg := JPEG{
		SOI:        jpegarray[0:2],
		APP0:       jpegarray[2:4],
		Length:     jpegarray[4:6],
		Identifier: jpegarray[6:11],
		Version:    jpegarray[11:13],
		Units:      jpegarray[13],
		Xres:       jpegarray[14:16],
		Yres:       jpegarray[16:18],
		XThumbnail: jpegarray[18],
		YThumbnail: jpegarray[19],
		Image:      jpegarray[19:],
	}

	//Check signature
	if string(jpeg.Identifier[0:4]) != "JFIF" {
		log.Fatal("Invalid JPEG Format: Invalid Identifier " + string(jpeg.Identifier))
	}
	return jpeg
}

//Concat byte arrays in image interface
func Concat(x interface{}) []byte {
	Concat := []byte{}
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.Slice:
			Concat = append(Concat, v.Field(i).Interface().([]byte)...)
		case reflect.Uint8:
			Concat = append(Concat, v.Field(i).Interface().(byte))
		}
	}

	return Concat
}
