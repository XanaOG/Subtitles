package main

import (
	Extract "github.com/XanaOG/Subtitles/Core/Extract"
	Transcribe "github.com/XanaOG/Subtitles/Core/Transcribe"
)

func main() {
	Extract.Audio()
	Transcribe.Upload()
	Transcribe.Transcribe()
	Transcribe.Save()
}
