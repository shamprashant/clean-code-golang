/*
🧠 Open/Closed Challenge Task
You are designing a logging system for a Go application.

🎯 Requirements:
The application currently supports logging to console.

Now there's a need to add logging to file and remote server.

You should be able to add new logging destinations in the future (like DB, S3, etc.) without modifying the core logic that triggers logs.

📝 Your Objective:
Design a logging system that follows the Open/Closed Principle.

Use Go interfaces and structs as you see fit.

You are free to choose how the log message flows to different destinations.

The main code that sends logs should not change if a new destination is added.
*/

package main

import "fmt"

type Logger interface {
	Log(string)
}

type ConsoleLogger struct{}
type FileLogger struct{}
type RemoteServerLogger struct{}
type DBLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("Logging to Console: ", message)
}

func (fl FileLogger) Log(message string) {
	fmt.Println("Logging to File: ", message)
}

func (rsl RemoteServerLogger) Log(message string) {
	fmt.Println("Logging to Remote Server: ", message)
}

// log to DB

func (dbl DBLogger) Log(message string) {
	fmt.Println("Logging to DB: ", message)
}

func main() {
	message := "Welcome to OCP!!"
	cl := ConsoleLogger{}
	fl := FileLogger{}
	rsl := RemoteServerLogger{}
	dbl := DBLogger{}

	cl.Log(message)
	fl.Log(message)
	rsl.Log(message)
	dbl.Log(message)
}
