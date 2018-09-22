package main

import(
	"github.com/ricallinson/tts"
)

func main() {
	tts := tts.CreateTts()
	tts.Speak("Make the computer speak.")
}
