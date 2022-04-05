package transcoder

import "strings"

type video struct {
	filename	string
	extension	string
}

func NewVideo(file string) video {
	init := strings.Split(file, ".")
	return video{filename: init[0], extension: init[1]}
}

func ReturnVideo(file video) string {
	return file.filename + "." + file.extension
}