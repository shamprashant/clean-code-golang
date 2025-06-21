/*
🧪 Task – Interface Segregation Challenge
You are building a media player system.

🎯 Your scenario:
Some players only support audio.

Some players support both audio and video.

Some devices are record-only, they don’t play anything.

🚧 Your Objective:
Design interfaces that do not force any type to implement unwanted methods.

Implement at least 3 types of devices:

An audio player

A video player

A voice recorder

Keep interfaces small and focused, and use composition where needed.
*/

package main

import "fmt"

type AudioPlayer interface {
	PlayAudio()
}

type VideoPlayer interface {
	PlayVideo()
}

type AudioRecorder interface {
	RecordAudio()
}

type Radio struct {
	AudioPlayer
}
type LEDTV struct {
	AudioPlayer
	VideoPlayer
}
type TapeRecorder struct {
	AudioPlayer
	AudioRecorder
}

func (r Radio) PlayAudio() {
	fmt.Println("Playing audio in Radio")
}

func (ledTv LEDTV) PlayAudio() {
	fmt.Println("Playing audio in LED TV")
}

func (ledTv LEDTV) PlayVideo() {
	fmt.Println("Playing video in LED TV")
}

func (tr TapeRecorder) PlayAudio() {
	fmt.Println("Playing audio in Tape Recorder")
}

func (tr TapeRecorder) RecordAudio() {
	fmt.Println("Playing audio in Tape Recorder")
}

func main() {
	var a AudioPlayer = Radio{}
	a.PlayAudio()

	var av VideoPlayer = LEDTV{}
	av.PlayVideo()

	var r AudioRecorder = TapeRecorder{}
	r.RecordAudio()
}
