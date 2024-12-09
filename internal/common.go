package internal

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var ErrStdEmpty = errors.New("std empty")

func GetStdin() (string, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to stat stdin: %w", err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return "", ErrStdEmpty
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to read stdin: %w", err)
	}
	return string(data), nil
}

func FileExists(path string) (exists bool, isDir bool, err error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, false, fmt.Errorf("failed to stat file: %w", err)
	}
	if fi.IsDir() {
		return true, true, nil
	}
	return true, false, nil
}

func ReadFile(path string) (string, error) {
	exists, isDir, err := FileExists(path)
	if err != nil {
		return "", fmt.Errorf("failed to check if file exists: %w", err)
	}
	if !exists {
		return "", fmt.Errorf("file does not exist: %s", path)
	}
	if isDir {
		return "", fmt.Errorf("file is a directory: %s", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return string(data), nil
}

func WriteFile(path string, data string) error {
    exists, isDir,  err := FileExists(path)
    if err != nil {
        return fmt.Errorf("failed to check if output file exists: %w", err)
    }
    if exists {
        if isDir {
            return WriteFile(paths, data string)

        }
    }
}
