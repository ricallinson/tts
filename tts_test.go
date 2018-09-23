package tts

import (
	. "github.com/ricallinson/simplebdd"
	"path"
	"reflect"
	"testing"
)

func TestTts(t *testing.T) {

	var tts *Tts

	BeforeEach(func() {
		tts = CreateTts()
	})

	AfterEach(func() {

	})

	Describe("CreateStory()", func() {

		It("should return an Story object", func() {
			AssertEqual(reflect.TypeOf(tts).String(), "*tts.Tts")
		})
	})

	Describe("Speak()", func() {

		It("should speak 'hello tests'", func() {
			defer func() {
				if r := recover(); r == nil {
					AssertEqual(true, true)
				} else {
					AssertEqual(false, true)
				}
			}()
			tts.Speak("hello tests")
		})

		It("should NOT speak 'hello tests'", func() {
			defer func() {
				if r := recover(); r == nil {
					AssertEqual(false, true)
				} else {
					AssertEqual(true, true)
				}
			}()
			tts.espeakExe = path.Join("dev", "null")
			tts.Speak("hello tests")
		})

	})

	Describe("listEspeakDir()", func() {
		It("should list all espeak files for OS", func() {
			count := 0
			for _, file := range tts.listEspeakDir() {
				if file == "dictsource" || file == "espeak-data" || file == "speak" {
					count++
				}
			}
			AssertEqual(count, 3)
		})
	})

	Report(t)
}
