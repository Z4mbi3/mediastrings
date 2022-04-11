package main

import (
	"github.com/Z4mbi3/mediastrings/ocr"
)

func main() {
	ocr.ReadImage("https://www.exampleimage.com/example.png", ocr.Azure)
}