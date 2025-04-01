package files

import (
	"path/filepath"
	"strings"
)

type State int

const (
	Directory = State(0)
	File      = State(1)
	NotPath   = State(2)
)

func isDirectory(path string) bool {
	return strings.HasSuffix(path, "/") || strings.HasSuffix(path, "\\")
}

func isFile(path string) bool {
	return filepath.Ext(path) != ""
}

func GetPathState(path string) State {
	directory := isDirectory(path)
	file := isFile(path)

	switch {
	case directory:
		return Directory
	case file:
		return File
	default:
		return NotPath
	}
}
