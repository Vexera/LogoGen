package vexeralogomeme_test

import (
	"image/png"
	"os"
	"testing"

	"github.com/SilverCory/vexeralogomeme"
)

func TestMeme(t *testing.T) {

	img, err := vexeralogomeme.CreateLogo("V", "vexera")
	if err != nil {
		t.Error(err)
	}

	f, err := os.Create("img.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)

}
