package getter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type FileWriter struct {
	currentDirectory string
}

func NewFileWriter() *FileWriter {
	currentDirectory, directoryError := os.Getwd()
	if directoryError != nil {
		fmt.Println("Error:", directoryError.Error())
		os.Exit(1)
	}
	return &FileWriter{currentDirectory + "/"}
}

func (receiver *FileWriter) WriteContent(path, content string) {
	directory := receiver.currentDirectory + path

	errorMakeDir := os.MkdirAll(directory, os.ModePerm)
	if errorMakeDir != nil {
		fmt.Println(errorMakeDir.Error())
		os.Exit(1)
	}

	const OutputFileName = "index.html"

	file, errorCreate := os.Create(directory + OutputFileName)
	if errorCreate != nil {
		fmt.Println(errorCreate.Error())
		os.Exit(1)
	}

	if _, errorWrite := fmt.Fprintln(file, content); errorWrite != nil {
		fmt.Println(errorWrite)
		os.Exit(1)
	}

	receiver.closeFile(file)
}

func (receiver *FileWriter) WriteImage(url, path string) {
	directory, imageFile := filepath.Split(path)
	directory = receiver.currentDirectory + directory

	errorMakeDir := os.MkdirAll(directory, os.ModePerm)
	if errorMakeDir != nil {
		fmt.Println(errorMakeDir.Error())
		os.Exit(1)
	}

	file, errorCreate := os.Create(directory + imageFile)
	if errorCreate != nil {
		fmt.Println(errorCreate.Error())
		os.Exit(1)
	}
	defer receiver.closeFile(file)

	response, errGet := http.Get(url)
	if errGet != nil {
		fmt.Println(errGet.Error())
		os.Exit(1)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	if _, err := io.Copy(file, response.Body); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func (receiver *FileWriter) closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
