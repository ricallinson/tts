package tts

import (
	"fmt"
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
	this.espeakExe = path.Join(this.espeakDir, "speak")
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
		if f.Name() == "dictsource" || f.Name() == "espeak-data" || f.Name() == "speak" {
			count++
		}
	}
	return count == 3
}

func (this *Tts) validateEspeakExe() bool {
	info, _ := os.Stat(this.espeakExe)
	mode := info.Mode()
	fmt.Println(mode)
	return false
}
