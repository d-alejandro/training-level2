package getter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
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

	if _, errorWrite := fmt.Fprint(file, content); errorWrite != nil {
		fmt.Println(errorWrite)
		os.Exit(1)
	}

	receiver.closeFile(file)
}

func (receiver *FileWriter) WriteResourceFile(url, path string) {
	if url == "" || path == "" {
		return
	}

	directory, resourceFile := filepath.Split(path)
	if resourceFile == "" {
		return
	}

	directory = receiver.currentDirectory + directory

	errorMakeDir := os.MkdirAll(directory, os.ModePerm)
	if errorMakeDir != nil {
		fmt.Println(errorMakeDir.Error())
		os.Exit(1)
	}

	file, errorCreate := os.Create(directory + resourceFile)
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

	bodyBytes, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		fmt.Println(readErr.Error())
		os.Exit(1)
	}

	body := string(bodyBytes)

	if _, errorWrite := fmt.Fprint(file, body); errorWrite != nil {
		fmt.Println(errorWrite)
		os.Exit(1)
	}

	if strings.HasSuffix(resourceFile, ".css") {
		receiver.processCSSFile(body, url)
	}
}

func (receiver *FileWriter) processCSSFile(body, url string) {
	regExpr := regexp.MustCompile(`url\(["'](....[^:].+?)['"]\)`)

	array := regExpr.FindAllStringSubmatch(body, -1)
	if array == nil {
		return
	}

	array = slices.CompactFunc(array, func(a, b []string) bool {
		return a[0] == b[0]
	})

	directory, resourceFile := filepath.Split(url)
	if resourceFile == "" {
		return
	}

	for _, values := range array {
		cssUrl := strings.TrimPrefix(values[1], "./")
		cssUrl = strings.TrimPrefix(cssUrl, "/")

		resourceUrl := directory + cssUrl
		resourcePath := filepath.Clean(receiver.replaceUrlToPath(resourceUrl))
		receiver.WriteResourceFile(resourceUrl, resourcePath)
	}
}

func (receiver *FileWriter) replaceUrlToPath(url string) string {
	if strings.HasPrefix(url, "https://") {
		return strings.TrimPrefix(url, "https://")
	}
	return strings.TrimPrefix(url, "http://")
}

func (receiver *FileWriter) closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
