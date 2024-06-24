package getter

import (
	"fmt"
	"os"
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

func (receiver *FileWriter) closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
