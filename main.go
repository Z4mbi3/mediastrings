package main

import (
	"github.com/Z4mbi3/mediastrings/ocr"
)

func main() {
	ocr.ReadImage("https://res.cloudinary.com/mike-student/image/upload/v1649150407/test_ks6ffw.png", ocr.Azure)
}