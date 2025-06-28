/*
🎯 Decorator Challenge (Interview Style)
You’re building an analytics system.

There’s a base DataWriter interface:

type DataWriter interface {
	Write(data string)
}
Your current implementation writes to a file.
Now, you want to add logging, encryption, and compression — without changing FileWriter.

🧪 Your Task:

Implement FileWriter

Then decorate it with:

LoggerWriter (logs every write)

EncryptWriter (adds dummy "encrypted:" prefix)

CompressWriter (adds dummy "compressed:" prefix)

All decorators should call the underlying writer’s Write().

You choose the layering order in main().
*/

package main

import "fmt"

type DataWriter interface {
	Write(data string)
}

type FileWriter struct{}

func (fw FileWriter) Write(data string) {
	fmt.Println("Writing data to file: ", data)
}

type LoggerWriter struct {
	writer DataWriter
}

func (lw LoggerWriter) Write(data string) {
	fmt.Println("Started writing... ")
	lw.writer.Write("hello")
	fmt.Println("Ended writing... ")
}

type EncryptWriter struct {
	writer DataWriter
}

func (ew EncryptWriter) Write(data string) {
	ew.writer.Write("encrypted: " + data)
}

type CompressWriter struct {
	writer DataWriter
}

func (cw CompressWriter) Write(data string) {
	cw.writer.Write("compressed: " + data)
}

func main() {
	fileWriter := FileWriter{}
	encryptWriter := EncryptWriter{writer: fileWriter}
	encryptWriter.Write("Prashant Sharma")
	compressWriter := CompressWriter{writer: fileWriter}
	compressWriter.Write("Prashant Sharma")
}
