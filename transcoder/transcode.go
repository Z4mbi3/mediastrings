package transcoder

import (
	"fmt"
	"os"
	"os/exec"
	b64 "encoding/base64"
)

func Transcode(file string) {
	os.Mkdir("./assets/", 0700)
	TranscodeVideo(NewVideo(file))
}

func TranscodeVideo(file video) {
	os.Mkdir("./assets/" + file.filename, 0700)

	tc := exec.Command(
		"ffmpeg", "-i", ReturnVideo(file),
		"-vf", "fps=1", "./assets/" + file.filename + "/" + b64.StdEncoding.EncodeToString([]byte(file.filename)) + "%d.png", 
	)
	tc.Stderr = os.Stderr
	tc.Run()
	
	fmt.Println("Splitting video into frames")
}