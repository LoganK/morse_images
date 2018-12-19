package main

import (
	"flag"
	"github.com/logank/morse_images"
)

func main() {
	flag.Parse()
	msg := "Hidden in plain site"
	if flag.NArg() > 0 {
		msg = flag.Arg(0)
	}
	outFile := "out.png"
	if flag.NArg() > 1 {
		outFile = flag.Arg(1)
	}

	img, err := morse_images.RenderMessage(msg)
	if err != nil {
		panic(err)
	}
	img.SavePNG(outFile)
}
