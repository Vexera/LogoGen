package vexeralogomeme

import (
	"image"
	"os"

	"golang.org/x/image/font"

	"github.com/fogleman/gg"
)

var backgroundImage image.Image
var bebasKaiFont font.Face
var craftismFont font.Face

func init() {
	var err error
	backgroundImage, err = loadImage("assets/bg.png")
	bebasKaiFont, err = gg.LoadFontFace("assets/BebasKai.ttf", 70)
	craftismFont, err = gg.LoadFontFace("assets/D3Craftism.ttf", 200)

	if err != nil {
		panic(err)
	}
}

func CreateLogo(topText, bottomText string) (image.Image, error) {
	ctx := gg.NewContextForImage(backgroundImage)
	xCenter := float64(ctx.Width()) / 2
	ctx.SetFontFace(craftismFont)
	ctx.SetHexColor("#FFF")
	ctx.DrawStringAnchored(topText, xCenter, (float64(ctx.Height())/2)-(ctx.FontHeight()/2)-50, 0.5, 0.5)
	ctx.SetFontFace(bebasKaiFont)
	ctx.DrawStringWrapped(bottomText, xCenter, (float64(ctx.Height())/2)+ctx.FontHeight()+5, 0.5, 0.5, float64(ctx.Width())-30, 1, gg.AlignCenter)
	return ctx.Image(), nil
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	return img, err
}
