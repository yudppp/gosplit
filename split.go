package gosplit

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Options struct {
	Infile          string
	OutPrefix       string
	LineCount       int
	HeaderLineCount int
}

func Split(opts Options) error {
	filepaths := strings.Split(opts.Infile, "/")
	fileNames := strings.Split(filepaths[len(filepaths)-1], ".")
	extention := ""
	if len(fileNames) > 1 {
		extention = fileNames[len(fileNames)-1]
	}

	file, err := os.Open(opts.Infile)
	if err != nil {
		return fmt.Errorf("infile is invalid")
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	headerLines := make([]string, opts.HeaderLineCount)
	for i := 0; i < opts.HeaderLineCount; i++ {
		b, _, err := reader.ReadLine()
		if err != nil {
			return fmt.Errorf("header line count is invalid")
		}
		headerLines[i] = string(b)
	}

	err = os.MkdirAll(filepath.Dir(opts.OutPrefix), 0755)
	if err != nil {
		return err
	}

	isLast := false
	index := 0
	for !isLast {
		lines := make([]string, 0, opts.LineCount)
		for i := 0; i < opts.LineCount; i++ {
			b, _, err := reader.ReadLine()
			if err == io.EOF {
				isLast = true
				break
			} else if err != nil {
				return err
			}
			lines = append(lines, string(b))
		}
		if len(lines) == 0 {
			break
		}
		data := append(headerLines, lines...)
		fileName := opts.OutPrefix + GenerateFileNameSuffix(index)
		if extention != "" {
			fileName = fmt.Sprintf("%s.%s", fileName, extention)
		}
		err = ioutil.WriteFile(fileName, []byte(strings.Join(data, "\n")), 0644)
		if err != nil {
			return err
		}
		index++
	}
	return nil
}

func GenerateFileNameSuffix(input int) string {
	// aa, ab, ac ...
	abc := "abcdefghijklmnopqrstuvwxyz"
	n := len(abc)
	result := ""

	data := input
	for data > 0 {
		result = string(abc[data%n]) + result
		data = data / n
	}
	if result == "" {
		return "aa"
	}
	if len(result) < 2 {
		return "a" + result
	}
	return result
}
