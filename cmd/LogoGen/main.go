package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/Vexera/LogoGen"
)

var (
	output     string
	topText    string
	bottomText string
)

func init() {
	flag.StringVar(&output, "output", "vexera.png", "-output=path.png")
	flag.StringVar(&topText, "top", "", `-top="dab"`)
	flag.StringVar(&bottomText, "bottom", "", `-bottom="bot"`)
}

func main() {
	flag.Parse()
	image, err := LogoGen.CreateLogo(topText, bottomText)
	if err != nil {
		fmt.Printf("error generating logo: %+v\n", err)
		os.Exit(1)
	}
	file, err := os.Create(output)
	if err != nil {
		fmt.Printf("error opening output file: %+v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	err = png.Encode(file, image)
	if err != nil {
		fmt.Printf("error encoding image: %+v\n", err)
		os.Exit(1)
	}
}
