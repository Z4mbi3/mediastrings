package ocr

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v2.0/computervision"
	"github.com/Azure/go-autorest/autorest"
	"github.com/otiai10/gosseract/v2"

	"github.com/Z4mbi3/mediastrings/env"
	"io/ioutil"
)

const (
	Tesseract string = "tesseract"
	Azure	  string = "azure"
)

var computerVisionContext context.Context
var COMPUTER_VISION_KEY = "COMPUTER_VISION_KEY"
var ENDPOINT_URL = "ENDPOINT_URL"

func VisionSetup() computervision.BaseClient{
	computerVisionKey := env.GetEnv(COMPUTER_VISION_KEY)
	endpointURL := env.GetEnv(ENDPOINT_URL)

	computerVisionClient := computervision.New(endpointURL)
	computerVisionClient.Authorizer = autorest.NewCognitiveServicesAuthorizer(computerVisionKey)
	computerVisionContext = context.Background()

	return computerVisionClient
}

func ReadImageSequence(path string, engine string) {
	switch engine {
		case Tesseract:
			images, err := ioutil.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}
			for _, image := range images {
				ReadImage(fmt.Sprintf("%s/%s", path, image.Name()), engine)
			}
			return
		case Azure:
			// Not implemented yet
	}
}

func ReadImage(image string, engine string) {
	switch engine {
		case Tesseract:
			TesseractReadImage(image)
		case Azure:
			BatchReadFileRemoteImage(VisionSetup(), image)
	}
}

// --- Engines ---

// Tesseract
func TesseractReadImage(image string) {
	fmt.Printf("Reading: %s\n\n", image)

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(image)
	text, _ := client.Text()
	if text == "" {
		fmt.Println("Could not detect data")
		return
	}
	fmt.Println(text)
}

// Azure
func BatchReadFileRemoteImage(client computervision.BaseClient, remoteImageURL string) {
	fmt.Println("Reading image: " + remoteImageURL)
    fmt.Println()
	var remoteImage computervision.ImageURL
	remoteImage.URL = &remoteImageURL

	textHeaders, err := client.BatchReadFile(computerVisionContext, remoteImage)
	if err != nil { log.Fatal(err) }

	operationLocation := autorest.ExtractHeaderValue("Operation-Location", textHeaders.Response)

	numbersOfCharsInOperationId := 36
	operationId := string(operationLocation[len(operationLocation)-numbersOfCharsInOperationId : len(operationLocation)])

	readOperationResult, err := client.GetReadOperationResult(computerVisionContext, operationId)
	if err != nil { log.Fatal(err) }

	// Wait for the operation to complete.
	i := 0
	maxRetries := 10

	fmt.Println("Recognizing text in a remote image with the batch Read API ...")
	for readOperationResult.Status != computervision.Failed &&
			readOperationResult.Status != computervision.Succeeded {
		if i >= maxRetries {
			break
		}
		i++

		fmt.Printf("Server status: %v, waiting %v seconds...\n", readOperationResult.Status, i)
		time.Sleep(1 * time.Second)

		readOperationResult, err = client.GetReadOperationResult(computerVisionContext, operationId)
		if err != nil { log.Fatal(err) }
	}
	// Display the results.
	fmt.Println()
	for _, recResult := range *(readOperationResult.RecognitionResults) {
		for _, line := range *recResult.Lines {
			fmt.Println(*line.Text)
		}
	}
}
