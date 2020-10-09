package lines

import (
	"bufio"
	"errors"
	"os"

	"github.com/gogs/chardet"
)

const (
	MEGABYTE int64 = 1024 * 1024
)

type lines struct {
	filePath     string
	lines        []string
	data         []byte
	file         *os.File
	charDetector *chardet.Detector
	charResult   *chardet.Result
}

func New(filePath string) *lines {
	return &lines{
		filePath:     filePath,
		charDetector: chardet.NewTextDetector(),
	}
}

// ReadFile reads the file passed in the constructor
// and runs a few checks on it as well as splits the file data into lines,
// so we can run checks against it easily
//
// this function needs to be run before anything else
func (l *lines) ReadFile() error {
	file, err := os.Open(l.filePath)

	if err != nil {
		return err
	}

	l.file = file
	defer l.file.Close()

	isLarge, err := l.isLarge()

	if err != nil {
		return err
	}

	if isLarge {
		return errors.New("file is too big to read (> 1mb)")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l.lines = append(l.lines, scanner.Text())
		l.data = append(l.data, scanner.Bytes()...)
	}

	l.detectCharset()

	if !l.isText() {
		return errors.New("file is not readable text")
	}

	return nil
}

// isViewable checks if the file is readable
func (l *lines) isViewable() (bool, error) {
	return l.isLarge()
}

// isLarge checks if the file is too big to load
func (l *lines) isLarge() (bool, error) {
	stats, err := l.file.Stat()

	if err != nil {
		return false, err
	}

	return stats.Size() > MEGABYTE, nil
}

func (l *lines) isText() bool {
	// couldn't detect a charset, assume binary
	if l.charResult == nil {
		return false
	}

	// treat empty files as text
	if len(l.data) == 0 {
		return true
	}

	return true
}

func (l *lines) detectCharset() {
	res, err := l.charDetector.DetectBest(l.data)

	// Could not detect charset so ignore it
	if err != nil {
		return
	}

	l.charResult = res
}
