package LogoGen

import (
	"image"
	"os"

	"github.com/oliamb/cutter"

	"golang.org/x/image/font"

	"github.com/fogleman/gg"
)

var backgroundImage image.Image
var bebasKaiFont font.Face
var craftismFont font.Face

const TrimPadding = 100

func init() {
	var err error
	backgroundImage, err = loadImage("assets/bg.png")
	bebasKaiFont, err = gg.LoadFontFace("assets/BebasKai.ttf", 85)
	craftismFont, err = gg.LoadFontFace("assets/D3Craftism.ttf", 240)

	if err != nil {
		panic(err)
	}
}

func CreateLogo(topText, bottomText string) (image.Image, error) {
	ctx := gg.NewContextForImage(backgroundImage)

	maxWidth := float64(0)

	xCenter := float64(ctx.Width()) / 2
	ctx.SetHexColor("#FFF")
	ctx.SetFontFace(craftismFont)
	ctx.DrawStringAnchored(topText, xCenter, (float64(ctx.Height())/2)-(ctx.FontHeight()/3), 0.5, 0.5)
	maxWidth, _ = ctx.MeasureString(topText)

	ctx.SetFontFace(bebasKaiFont)
	newMaxWidth := drawStringWrapped(ctx, bottomText, xCenter, (float64(ctx.Height())/2)+ctx.FontHeight()+30, 0.5, 0.5, float64(ctx.Width())-30, 1, gg.AlignCenter)
	if newMaxWidth > maxWidth {
		maxWidth = newMaxWidth
	}

	if int(maxWidth)+TrimPadding <= ctx.Width() {
		if int(maxWidth)+TrimPadding < ctx.Height() {
			maxWidth = float64(ctx.Height() - TrimPadding)
		}
		return cutter.Crop(ctx.Image(), cutter.Config{
			Width:  int(maxWidth + TrimPadding),
			Height: ctx.Height(),
			Mode:   cutter.Centered,
		})
	}

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

func drawStringWrapped(dc *gg.Context, s string, x, y, ax, ay, width, lineSpacing float64, align gg.Align) float64 {
	lines := dc.WordWrap(s, width)
	h := float64(len(lines)) * dc.FontHeight() * lineSpacing
	h -= (lineSpacing - 1) * dc.FontHeight()
	x -= ax * width
	y -= ay * h
	switch align {
	case gg.AlignLeft:
		ax = 0
	case gg.AlignCenter:
		ax = 0.5
		x += width / 2
	case gg.AlignRight:
		ax = 1
		x += width
	}
	ay = 1
	retSize := float64(0)
	for _, line := range lines {
		if w, _ := dc.MeasureString(line); w > retSize {
			retSize = w
		}
		dc.DrawStringAnchored(line, x, y, ax, ay)
		y += dc.FontHeight() * lineSpacing
	}

	return retSize
}
