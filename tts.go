package tts

import (
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

type Tts struct {
	espeakDir string
	espeakExe string
}

func CreateTts() *Tts {
	// Get the absolute path to this file.
	// _, filename, _, _ := runtime.Caller(0)
	this := &Tts{}
	this.espeakDir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "ricallinson", "tts", "espeak", runtime.GOOS)
	this.espeakExe = path.Join(this.espeakDir, "speak")
	return this
}

// Output the given string as audio.
func (this *Tts) Speak(s string) {
	cmd := exec.Command(this.espeakExe, "--path="+this.espeakDir, s)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
