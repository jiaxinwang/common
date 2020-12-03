package fs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

// Copy ...
func Copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}

// FileExist ...
func FileExist(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

// IsDir ...
func IsDir(name string) bool {
	info, err := os.Stat(name)
	return err == nil && info.IsDir()
}

// Readlines ...
func Readlines(file string) (lines []string, err error) {
	fd, err := os.Open(file)
	if err != nil {
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	var line string
	for {
		line, err = reader.ReadString('\n')
		lines = append(lines, line)
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
	}
	return
}

// Save ...
func Save(data []byte, fullpath string) (err error) {
	dir := filepath.Dir(fullpath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		logrus.WithError(err).WithField("fullpath", fullpath).Error()
		return
	}
	var f *os.File
	if f, err = os.Create(fullpath); err != nil {
		logrus.WithError(err).WithField("fullpath", fullpath).Error()
		return
	}
	defer f.Close()
	if _, err = f.Write(data); err != nil {
		logrus.WithError(err).WithField("fullpath", fullpath).Error()
		return
	}
	return nil
}
