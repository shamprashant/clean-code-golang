package main

import (
	"errors"
	"fmt"
)

// --- Interfaces ---

type BackupSource interface {
	Name() string
	Backup() ([]byte, error)
}

type StorageTarget interface {
	Name() string
	Store(data []byte) error
}

type Logger interface {
	Log(message string)
}

type Notifier interface {
	Notify(message string)
}

// --- Concrete Loggers ---

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("[LOG]:", message)
}

// --- Concrete Notifiers ---

type EmailNotifier struct{}

func (e EmailNotifier) Notify(msg string) {
	fmt.Println("[EMAIL ALERT]:", msg)
}

// --- Concrete Backup Sources ---

type FileSystemSource struct{}

func (fs FileSystemSource) Name() string { return "FileSystem" }
func (fs FileSystemSource) Backup() ([]byte, error) {
	return []byte("file data"), nil
}

type DatabaseSource struct{}

func (db DatabaseSource) Name() string { return "Database" }
func (db DatabaseSource) Backup() ([]byte, error) {
	return nil, errors.New("could not connect to DB")
}

// --- Concrete Storage Targets ---

type S3Target struct{}

func (s3 S3Target) Name() string            { return "S3" }
func (s3 S3Target) Store(data []byte) error { return nil }

type LocalDiskTarget struct{}

func (ld LocalDiskTarget) Name() string            { return "LocalDisk" }
func (ld LocalDiskTarget) Store(data []byte) error { return errors.New("disk full") }

// --- Coordinator ---

type BackupCoordinator struct {
	source   BackupSource
	target   StorageTarget
	logger   Logger
	notifier Notifier
}

func (bc BackupCoordinator) Run() {
	bc.logger.Log("Starting backup from " + bc.source.Name())

	data, err := bc.source.Backup()
	if err != nil {
		msg := fmt.Sprintf("Backup failed from %s: %v", bc.source.Name(), err)
		bc.logger.Log(msg)
		bc.notifier.Notify(msg)
		return
	}
	bc.logger.Log("Backup succeeded. Data size: " + fmt.Sprint(len(data)))

	err = bc.target.Store(data)
	if err != nil {
		msg := fmt.Sprintf("Storing to %s failed: %v", bc.target.Name(), err)
		bc.logger.Log(msg)
		bc.notifier.Notify(msg)
		return
	}
	bc.logger.Log("Stored successfully to " + bc.target.Name())
}

// --- Main ---

func main() {
	// Dependencies via interfaces (DIP)
	source := FileSystemSource{}
	target := S3Target{}
	logger := ConsoleLogger{}
	notifier := EmailNotifier{}

	coordinator := BackupCoordinator{
		source:   source,
		target:   target,
		logger:   logger,
		notifier: notifier,
	}

	coordinator.Run()

	// Try failing case
	fmt.Println("\n---- Trying failed backup from DB ----")
	coordinator2 := BackupCoordinator{
		source:   DatabaseSource{},
		target:   LocalDiskTarget{},
		logger:   logger,
		notifier: notifier,
	}
	coordinator2.Run()
}
