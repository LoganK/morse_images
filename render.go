package morse_images

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/martinlindhe/morse"
	"image"
	_ "image/png"
	"os"
)

// Accepts a morse message, returns a list of file names of the form:
// {Dit,Dah,Shh}_{Dit,Dah,Shh}.png
func asImageNames(morse string) ([]string, error) {
	var els []string
	for _, r := range morse {
		if r == ' ' {
			els = append(els, "Shh")
		} else if r == '.' {
			els = append(els, "Dit")
		} else if r == '-' {
			els = append(els, "Dah")
		} else if r == '/' {
			els = append(els, "Shh", "Shh")
		} else {
			return nil, fmt.Errorf("unknown morse symbol: %c (%U)", r, r)
		}
	}
	if len(els)%2 == 1 {
		// Pad the last character so we don't have to special case.
		els = append(els, "Shh")
	}

	var res []string
	for i := 0; i < len(els); i += 2 {
		res = append(res, fmt.Sprintf("%s_%s.png", els[i], els[i+1]))
	}

	return res, nil
}

func RenderMessage(s string) (*gg.Context, error) {
	files, err := asImageNames(morse.EncodeITU(s))
	if err != nil {
		return nil, fmt.Errorf("unable to translate morse: %s", err)
	}

	// Preload all the images for efficiency and to ensure we have everything.
	morseImages := make(map[string]image.Image)
	maxWidth := 0
	maxHeight := 0
	for _, file := range files {
		if _, ok := morseImages[file]; ok {
			continue
		}

		f, err := os.Open(file)
		if err != nil {
			return nil, fmt.Errorf("unable to open file %s: %s", file, err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			return nil, fmt.Errorf("unable to decode file %s: %s", file, err)
		}

		width := img.Bounds().Max.X - img.Bounds().Min.X
		if width > maxWidth {
			maxWidth = width
		}

		height := img.Bounds().Max.Y - img.Bounds().Min.Y
		if height > maxHeight {
			maxHeight = height
		}

		morseImages[file] = img
	}

	padding := 10
	needWidth := len(files)*maxWidth + ((len(files) + 1) * padding)
	needHeight := maxHeight * 2 * padding
	g := gg.NewContext(needWidth, needHeight)

	x := padding + (maxWidth / 2)
	y := padding + (maxHeight / 2)
	for _, imgFile := range files {
		g.DrawImageAnchored(morseImages[imgFile], x, y, 0.5, 0.5)
		x += maxWidth
	}

	return g, nil
}
