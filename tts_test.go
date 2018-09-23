package tts

import (
	. "github.com/ricallinson/simplebdd"
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
			tts.Speak("hello tests")
			AssertEqual(true, true)
		})

	})

	Report(t)
}
