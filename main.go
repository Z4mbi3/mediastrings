package main

import (
	"github.com/Z4mbi3/vidstrings/ocr"
	// W "vidstrings/web"
)

func main() {
	// https://res.cloudinary.com/mike-student/image/upload/v1649150407/test_ks6ffw.png
	ocr.ReadImage("https://i.insider.com/61ddf46c1025b20018bb3cdc?width=1000&format=jpeg&auto=webp", ocr.Azure)
	// W.RunWeb()
}