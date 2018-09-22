# tts

## Install

	go get github.com/ricallinson/tts

## Usage

	package main

	import(
		"github.com/ricallinson/tts"
	)

	func main() {
		tts := tts.CreateTts()
		tts.Speak("Make the computer speak.")
	}

## Notes

	./speak --path="./" "Make the computer speak."
