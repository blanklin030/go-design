package main

import "testing"

func TestNewLogManager(t *testing.T) {
	fileLogging := &FileLogging{}
	logManager := NewLogManager(fileLogging)
	logManager.Info()
	logManager.Error()

	databaseLogging := &DatabaseLogging{}
	logManager.Logging = databaseLogging
	logManager.Info()
	logManager.Error()
}
