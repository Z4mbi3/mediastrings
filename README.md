# Mediastrings

This is a OCR library for videos and images written in the `Go` language. It currently supports 2 engines (Tesseract and Azure Vision).

# Install

1. [tesseract-ocr](https://github.com/tesseract-ocr/tessdoc)

# Setup

To use the Azure engine you will have to make a `Computer Vision Resource` on your Azure portal. Once you created this, add a .env file in your project with the following keys and their values from the `Keys and Endpoint` tab under resource management.

```
COMPUTER_VISION_KEY=<your_computer_vision_key>
ENDPOINT_URL=https://<your_endpoint>.cognitiveservices.azure.com/
```

# Examples

## Read Image

The `ReadImage` method allows you to extract text from a single image.

```go
package main

import (
	"github.com/Z4mbi3/mediastrings/ocr"
)

func main() {
	ocr.ReadImage("https://www.exampleimage.com/example.png", ocr.Azure) // The Azure engine currently only works with remote images.
    ocr.ReadImage("./example.png", ocr.Tesseract) // The Tesseract engine only works for local images.
}
```

## Read Video

The `ReadVideo` method takes a video file and splits it into frames/images, then extracts the text from the sequence of all those frames/images.

```go
package main

import (
	"github.com/Z4mbi3/mediastrings/ocr"
)

func main() {
	
}
```