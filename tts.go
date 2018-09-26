package tts

import (
	"io/ioutil"
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
	this := &Tts{}
	this.espeakDir = path.Join(os.Getenv("GOPATH"), "src", "github.com", "ricallinson", "tts", "espeak", runtime.GOOS)
	this.espeakExe = path.Join(this.espeakDir, "espeak")
	return this
}

// Output the given string as audio.
func (this *Tts) Speak(s string) {
	cmd := exec.Command(this.espeakExe, "--path="+this.espeakDir, s)
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func (this *Tts) validateEspeakDir() bool {
	files, err := ioutil.ReadDir(this.espeakDir)
	if err != nil {
		return false
	}
	count := 0
	for _, f := range files {
		if f.Name() == "espeak-data" || f.Name() == "espeak" {
			count++
		}
	}
	return count == 2
}

func (this *Tts) validateEspeakExe() bool {
	info, _ := os.Stat(this.espeakExe)
	mode := info.Mode()
	return info.Mode().String()[len(mode.String())-1:] == "x"
}
