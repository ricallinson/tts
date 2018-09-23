# tts

[![Build Status](https://travis-ci.org/ricallinson/tts.svg?branch=master)](https://travis-ci.org/ricallinson/tts) [![Build status](https://ci.appveyor.com/api/projects/status/02xh6a9o5fqtlvtq/branch/master?svg=true)](https://ci.appveyor.com/project/ricallinson/tts/branch/master)

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
