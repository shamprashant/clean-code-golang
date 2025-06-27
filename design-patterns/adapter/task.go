/*
🎧 Problem Statement
You have a legacy media player VLCPlayer that doesn't match the modern MusicPlayer interface. You must adapt it using the Adapter Pattern — without modifying existing code.

🏗️ Interface You Need to Implement:

type MusicPlayer interface {
	PlayAudio(file string) string
}
📼 Legacy Class (Can’t Change This):

type VLCPlayer struct{}

func (v VLCPlayer) PlayVLC(file string) string {
	return fmt.Sprintf("Playing VLC file: %s", file)
}
🧪 Your Task:
Create a struct VLCAdapter

Make it implement the MusicPlayer interface

Inside PlayAudio(), call PlayVLC() of VLCPlayer

In main(), demonstrate that a VLCPlayer can be used as a MusicPlayer

✅ Expected Output Example:

Playing VLC file: song.mp3
*/

package main

import "fmt"

type MusicPlayer interface {
	PlayAudio(file string) string
}

type VLCPlayer struct{}

type VLCPlayerAdapter struct {
	VLCPlayer VLCPlayer
}

func (v VLCPlayer) PlayVLC(file string) string {
	return fmt.Sprintf("Playing VLC file: %s", file)
}

func (v VLCPlayerAdapter) PlayAudio(file string) string {
	return v.VLCPlayer.PlayVLC(file)
}

// func main() {
// 	vlc := VLCPlayerAdapter{}
// 	fmt.Println(vlc.PlayAudio("oonchi.mp3"))
// }
